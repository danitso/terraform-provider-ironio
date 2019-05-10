package main

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iron-io/iron_go3/config"
	"github.com/iron-io/iron_go3/mq"
)

// resourcePushQueue() manages IronMQ push queues.
func resourcePushQueue() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_queue": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of an error queue",
				Default:     "",
			},
			"multicast": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to create a multicast push queue",
				ForceNew:    true,
				Default:     true,
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the queue",
				ForceNew:    true,
			},
			"project_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The project id",
				ForceNew:    true,
			},
			"retries": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of retries before moving on to the next message",
				Default:     3,
			},
			"retries_delay": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of seconds to wait before re-sending a failed message",
				Default:     60,
			},
			"subscriber": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"headers": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"url": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},

		Create: resourcePushQueueCreate,
		Read:   resourcePushQueueRead,
		Update: resourcePushQueueUpdate,
		Delete: resourcePushQueueDelete,
	}
}

// resourcePushQueueBuildInfo() builds a queue information object.
func resourcePushQueueBuildInfo(d *schema.ResourceData) (mq.QueueInfo, error) {
	queueName := d.Get("name").(string)
	queueType := "unicast"

	multicast := d.Get("multicast").(bool)
	subscribers := d.Get("subscriber").([]interface{})

	if len(subscribers) == 0 {
		return mq.QueueInfo{}, fmt.Errorf("A push queue requires at least one subscriber")
	}

	if multicast {
		queueType = "multicast"
	}

	pushInfo := mq.PushInfo{
		RetriesDelay: d.Get("retries_delay").(int),
		Retries:      d.Get("retries").(int),
		ErrorQueue:   d.Get("error_queue").(string),
	}

	for _, v := range subscribers {
		resource_data := v.(map[string]interface{})
		url := resource_data["url"].(string)

		if url == "" {
			continue
		}

		push_headers := map[string]string{}

		for hn, hv := range resource_data["headers"].(map[string]interface{}) {
			push_headers[hn] = hv.(string)
		}

		subscriber := mq.QueueSubscriber{
			Name:    resource_data["name"].(string),
			URL:     url,
			Headers: push_headers,
		}

		pushInfo.Subscribers = append(pushInfo.Subscribers, subscriber)
	}

	queueInfo := mq.QueueInfo{
		Name: queueName,
		Type: queueType,
		Push: &pushInfo,
	}

	return queueInfo, nil
}

// resourcePushQueueCreate() creates a new push queue.
func resourcePushQueueCreate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsMQ := config.Settings{}
	clientSettingsMQ.UseSettings(&clientSettings.MQ)

	projectID := d.Get("project_id").(string)
	queueName := d.Get("name").(string)

	clientSettingsMQ.ProjectId = projectID
	queueInfo, err := resourcePushQueueBuildInfo(d)

	if err != nil {
		return err
	}

	_, err = mq.ConfigCreateQueue(queueInfo, &clientSettingsMQ)

	if err != nil {
		return err
	}

	d.SetId(queueNameToID(clientSettingsMQ.ProjectId, queueName))

	return resourcePushQueueRead(d, m)
}

// resourcePushQueueRead() reads information about an existing push queue.
func resourcePushQueueRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsMQ := config.Settings{}
	clientSettingsMQ.UseSettings(&clientSettings.MQ)

	projectID := d.State().Attributes["project_id"]
	queueName := d.State().Attributes["name"]

	if projectID != "" {
		clientSettingsMQ.ProjectId = projectID
	}

	queue := mq.ConfigNew(queueName, &clientSettingsMQ)
	queueInfo, err := queue.Info()

	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")

			return nil
		} else {
			return err
		}
	}

	if queueInfo.Type == "pull" {
		d.SetId("")

		return nil
	}

	if queueInfo.Type == "multicast" {
		d.Set("multicast", true)
	} else {
		d.Set("multicast", false)
	}

	if queueInfo.Push != nil {
		d.Set("error_queue", queueInfo.Push.ErrorQueue)
		d.Set("retries", queueInfo.Push.Retries)
		d.Set("retries_delay", queueInfo.Push.RetriesDelay)

		if queueInfo.Push.Subscribers != nil {
			subscribers := make([]interface{}, len(queueInfo.Push.Subscribers))

			for k, v := range queueInfo.Push.Subscribers {
				subscriberMap := make(map[string]interface{})
				headersMap := make(map[string]interface{})

				for hk, hv := range v.Headers {
					headersMap[hk] = hv
				}

				subscriberMap["name"] = v.Name
				subscriberMap["url"] = v.URL
				subscriberMap["headers"] = headersMap

				subscribers[k] = subscriberMap
			}

			d.Set("subscriber", subscribers)
		}
	}

	return nil
}

// resourcePushQueueUpdate() updates an existing push queue.
func resourcePushQueueUpdate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsMQ := config.Settings{}
	clientSettingsMQ.UseSettings(&clientSettings.MQ)

	projectID := d.Get("project_id").(string)
	queueName := d.Get("name").(string)

	clientSettingsMQ.ProjectId = projectID
	queueInfo, err := resourcePushQueueBuildInfo(d)

	if err != nil {
		return err
	}

	queue := mq.ConfigNew(queueName, &clientSettingsMQ)
	_, err = queue.Update(queueInfo)

	if err != nil {
		return err
	}

	return nil
}

// resourcePushQueueDelete() deletes an existing push queue.
func resourcePushQueueDelete(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsMQ := config.Settings{}
	clientSettingsMQ.UseSettings(&clientSettings.MQ)

	projectID := d.State().Attributes["project_id"]
	queueName := d.State().Attributes["name"]

	if projectID != "" {
		clientSettingsMQ.ProjectId = projectID
	}

	queue := mq.ConfigNew(queueName, &clientSettingsMQ)

	err := queue.Delete()

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
