package main

import (
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iron-io/iron_go3/config"
	"github.com/iron-io/iron_go3/mq"
)

const DataSourcePushQueueErrorQueueKey = "error_queue"
const DataSourcePushQueueHeadersKey = "headers"
const DataSourcePushQueueMessageCountKey = "message_count"
const DataSourcePushQueueMessageCountTotalKey = "message_count_total"
const DataSourcePushQueueMulticastKey = "multicast"
const DataSourcePushQueueNameKey = "name"
const DataSourcePushQueueProjectIDKey = "project_id"
const DataSourcePushQueueRetriesDelayKey = "retries_delay"
const DataSourcePushQueueRetriesKey = "retries"
const DataSourcePushQueueSubscriberKey = "subscriber"
const DataSourcePushQueueURLKey = "url"

// dataSourcePushQueue() reads information about IronMQ push queues.
func dataSourcePushQueue() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			DataSourcePushQueueErrorQueueKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of an error queue",
			},
			DataSourcePushQueueMessageCountKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of messages currently in the queue",
			},
			DataSourcePushQueueMessageCountTotalKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of messages which have been processed by the queue",
			},
			DataSourcePushQueueMulticastKey: &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether to create a multicast push queue",
			},
			DataSourcePushQueueNameKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the queue",
				ForceNew:    true,
			},
			DataSourcePushQueueProjectIDKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The project id",
				ForceNew:    true,
			},
			DataSourcePushQueueRetriesKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of retries before moving on to the next message",
			},
			DataSourcePushQueueRetriesDelayKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of seconds to wait before re-sending a failed message",
			},
			DataSourcePushQueueSubscriberKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						DataSourcePushQueueHeadersKey: {
							Type:     schema.TypeMap,
							Optional: true,
						},
						DataSourcePushQueueNameKey: {
							Type:     schema.TypeString,
							Optional: true,
						},
						DataSourcePushQueueURLKey: {
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

	projectID := d.Get(DataSourcePushQueueProjectIDKey).(string)
	queueName := d.Get(DataSourcePushQueueNameKey).(string)

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
