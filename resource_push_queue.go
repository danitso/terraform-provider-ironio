package main

import (
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iron-io/iron_go3/config"
	"github.com/iron-io/iron_go3/mq"
)

const ResourcePushQueueErrorQueueKey = "error_queue"
const ResourcePushQueueHeadersKey = "headers"
const ResourcePushQueueMessageCountKey = "message_count"
const ResourcePushQueueMessageCountTotalKey = "message_count_total"
const ResourcePushQueueMulticastKey = "multicast"
const ResourcePushQueueNameKey = "name"
const ResourcePushQueueProjectIDKey = "project_id"
const ResourcePushQueueRetriesDelayKey = "retries_delay"
const ResourcePushQueueRetriesKey = "retries"
const ResourcePushQueueSubscriberKey = "subscriber"
const ResourcePushQueueURLKey = "url"

// resourcePushQueue() manages IronMQ push queues.
func resourcePushQueue() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			ResourcePushQueueErrorQueueKey: &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of an error queue",
				Default:     "",
			},
			ResourcePushQueueMessageCountKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of messages currently in the queue",
			},
			ResourcePushQueueMessageCountTotalKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of messages which have been processed by the queue",
			},
			ResourcePushQueueMulticastKey: &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to create a multicast push queue",
				ForceNew:    true,
				Default:     true,
			},
			ResourcePushQueueNameKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the queue",
				ForceNew:    true,
			},
			ResourcePushQueueProjectIDKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The project id",
				ForceNew:    true,
			},
			ResourcePushQueueRetriesKey: &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of retries before moving on to the next message",
				Default:     3,
			},
			ResourcePushQueueRetriesDelayKey: &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of seconds to wait before re-sending a failed message",
				Default:     60,
			},
			ResourcePushQueueSubscriberKey: &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						ResourcePushQueueHeadersKey: {
							Type:     schema.TypeMap,
							Optional: true,
						},
						ResourcePushQueueNameKey: {
							Type:     schema.TypeString,
							Optional: true,
						},
						ResourcePushQueueURLKey: {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				MinItems: 1,
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
	queueName := d.Get(ResourcePushQueueNameKey).(string)
	queueType := "unicast"

	multicast := d.Get(ResourcePushQueueMulticastKey).(bool)
	subscribers := d.Get(ResourcePushQueueSubscriberKey).([]interface{})

	if multicast {
		queueType = "multicast"
	}

	pushInfo := mq.PushInfo{
		RetriesDelay: d.Get(ResourcePushQueueRetriesDelayKey).(int),
		Retries:      d.Get(ResourcePushQueueRetriesKey).(int),
		ErrorQueue:   d.Get(ResourcePushQueueErrorQueueKey).(string),
	}

	for _, v := range subscribers {
		resource_data := v.(map[string]interface{})
		url := resource_data[ResourcePushQueueURLKey].(string)

		if url == "" {
			continue
		}

		push_headers := map[string]string{}

		for hn, hv := range resource_data[ResourcePushQueueHeadersKey].(map[string]interface{}) {
			push_headers[hn] = hv.(string)
		}

		subscriber := mq.QueueSubscriber{
			Name:    resource_data[ResourcePushQueueNameKey].(string),
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

	projectID := d.Get(ResourcePushQueueProjectIDKey).(string)
	queueName := d.Get(ResourcePushQueueNameKey).(string)

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

// resourcePushQueueParseInfo() parses information about an existing push queue.
func resourcePushQueueParseInfo(d *schema.ResourceData, queueInfo *mq.QueueInfo) error {
	if queueInfo.Type == "multicast" {
		d.Set(ResourcePushQueueMulticastKey, true)
	} else {
		d.Set(ResourcePushQueueMulticastKey, false)
	}

	if queueInfo.Push != nil {
		d.Set(ResourcePushQueueErrorQueueKey, queueInfo.Push.ErrorQueue)
		d.Set(ResourcePushQueueRetriesKey, queueInfo.Push.Retries)
		d.Set(ResourcePushQueueRetriesDelayKey, queueInfo.Push.RetriesDelay)

		if queueInfo.Push.Subscribers != nil {
			subscribers := make([]interface{}, len(queueInfo.Push.Subscribers))

			for k, v := range queueInfo.Push.Subscribers {
				subscriberMap := make(map[string]interface{})
				headersMap := make(map[string]interface{})

				for hk, hv := range v.Headers {
					headersMap[hk] = hv
				}

				subscriberMap[ResourcePushQueueNameKey] = v.Name
				subscriberMap[ResourcePushQueueURLKey] = v.URL
				subscriberMap[ResourcePushQueueHeadersKey] = headersMap

				subscribers[k] = subscriberMap
			}

			d.Set(ResourcePushQueueSubscriberKey, subscribers)
		}
	}

	d.Set(ResourcePushQueueMessageCountKey, queueInfo.Size)
	d.Set(ResourcePushQueueMessageCountTotalKey, queueInfo.TotalMessages)

	return nil
}

// resourcePushQueueRead() reads information about an existing push queue.
func resourcePushQueueRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsMQ := config.Settings{}
	clientSettingsMQ.UseSettings(&clientSettings.MQ)

	projectID := d.State().Attributes[ResourcePushQueueProjectIDKey]
	queueName := d.State().Attributes[ResourcePushQueueNameKey]

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

	return resourcePushQueueParseInfo(d, &queueInfo)
}

// resourcePushQueueUpdate() updates an existing push queue.
func resourcePushQueueUpdate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsMQ := config.Settings{}
	clientSettingsMQ.UseSettings(&clientSettings.MQ)

	projectID := d.Get(ResourcePushQueueProjectIDKey).(string)
	queueName := d.Get(ResourcePushQueueNameKey).(string)

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

	projectID := d.State().Attributes[ResourcePushQueueProjectIDKey]
	queueName := d.State().Attributes[ResourcePushQueueNameKey]

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
