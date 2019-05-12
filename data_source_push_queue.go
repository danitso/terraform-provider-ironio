package main

import (
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iron-io/iron_go3/config"
	"github.com/iron-io/iron_go3/mq"
)

// dataSourcePushQueue() reads information about IronMQ push queues.
func dataSourcePushQueue() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_queue": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of an error queue",
			},
			"message_count": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of messages currently in the queue",
			},
			"message_count_total": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of messages which have been processed by the queue",
			},
			"multicast": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether to create a multicast push queue",
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
				Computed:    true,
				Description: "The number of retries before moving on to the next message",
			},
			"retries_delay": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of seconds to wait before re-sending a failed message",
			},
			"subscriber": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
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

		Read: dataSourcePushQueueRead,
	}
}

// dataSourcePushQueueRead() reads information about an existing push queue.
func dataSourcePushQueueRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsMQ := config.Settings{}
	clientSettingsMQ.UseSettings(&clientSettings.MQ)

	projectID := d.Get("project_id").(string)
	queueName := d.Get("name").(string)

	clientSettingsMQ.ProjectId = projectID

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

	d.SetId(queueNameToID(clientSettingsMQ.ProjectId, queueName))

	return resourcePushQueueParseInfo(d, &queueInfo)
}
