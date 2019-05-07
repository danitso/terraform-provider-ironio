package ironio

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
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the queue",
				ForceNew:    true,
			},
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
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"url": {
							Type:     schema.TypeString,
							Required: true,
						},
						"headers": {
							Type:     schema.TypeMap,
							Optional: true,
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
	queue_name := d.Get("name").(string)
	queue_type := "unicast"

	multicast := d.Get("multicast").(bool)
	subscribers := d.Get("subscriber").(*schema.Set).List()

	if len(subscribers) == 0 {
		return mq.QueueInfo{}, fmt.Errorf("A push queue requires at least one subscriber")
	}

	if multicast {
		queue_type = "multicast"
	}

	push_info := mq.PushInfo{
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

		push_info.Subscribers = append(push_info.Subscribers, subscriber)
	}

	queue_info := mq.QueueInfo{
		Name: queue_name,
		Type: queue_type,
		Push: &push_info,
	}

	return queue_info, nil
}

// resourcePushQueueCreate() creates a new push queue.
func resourcePushQueueCreate(d *schema.ResourceData, m interface{}) error {
	client_settings := m.(config.Settings)

	queue_name := d.Get("name").(string)
	queue_info, err := resourcePushQueueBuildInfo(d)

	if err != nil {
		return err
	}

	_, err = mq.ConfigCreateQueue(queue_info, &client_settings)

	if err != nil {
		return err
	}

	d.SetId(queueNameToId(client_settings.ProjectId, queue_name))

	return resourcePushQueueRead(d, m)
}

// resourcePushQueueRead() reads information about an existing push queue.
func resourcePushQueueRead(d *schema.ResourceData, m interface{}) error {
	client_settings := m.(config.Settings)

	queue_name := d.Get("name").(string)
	queue := mq.ConfigNew(queue_name, &client_settings)

	queue_info, err := queue.Info()

	if err != nil {
		if strings.Contains(err.Error(), "Queue not found") {
			d.SetId("")

			return nil
		} else {
			return err
		}
	}

	if queue_info.Type == "pull" {
		d.SetId("")

		return nil
	}

	if queue_info.Type == "multicast" {
		d.Set("multicast", true)
	} else {
		d.Set("multicast", false)
	}

	if queue_info.Push != nil {
		d.Set("error_queue", queue_info.Push.ErrorQueue)
		d.Set("retries", queue_info.Push.Retries)
		d.Set("retries_delay", queue_info.Push.RetriesDelay)

		if queue_info.Push.Subscribers != nil {
			subscriber_set := d.Get("subscriber").(*schema.Set)
			subscriber_list := make([]interface{}, len(queue_info.Push.Subscribers))

			for k, v := range queue_info.Push.Subscribers {
				subscriber_map := make(map[string]interface{})
				headers_map := make(map[string]interface{})

				for hk, hv := range v.Headers {
					headers_map[hk] = hv
				}

				subscriber_map["name"] = v.Name
				subscriber_map["url"] = v.URL
				subscriber_map["headers"] = headers_map

				subscriber_list[k] = subscriber_map
			}

			subscriber_set = schema.NewSet(subscriber_set.F, subscriber_list)

			d.Set("subscriber", subscriber_set)
		}
	}

	return nil
}

// resourcePushQueueUpdate() updates an existing push queue.
func resourcePushQueueUpdate(d *schema.ResourceData, m interface{}) error {
	client_settings := m.(config.Settings)

	queue_name := d.Get("name").(string)
	queue_info, err := resourcePushQueueBuildInfo(d)

	if err != nil {
		return err
	}

	queue := mq.ConfigNew(queue_name, &client_settings)
	_, err = queue.Update(queue_info)

	if err != nil {
		return err
	}

	return nil
}

// resourcePushQueueDelete() deletes an existing push queue.
func resourcePushQueueDelete(d *schema.ResourceData, m interface{}) error {
	client_settings := m.(config.Settings)

	queue_name := d.Get("name").(string)
	queue := mq.ConfigNew(queue_name, &client_settings)

	err := queue.Delete()

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
