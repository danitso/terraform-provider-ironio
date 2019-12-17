/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package ironiotf

import (
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iron-io/iron_go3/config"
	"github.com/iron-io/iron_go3/mq"
)

const dataSourcePushQueueErrorQueueKey = "error_queue"
const dataSourcePushQueueHeadersKey = "headers"
const dataSourcePushQueueMessageCountKey = "message_count"
const dataSourcePushQueueMessageCountTotalKey = "message_count_total"
const dataSourcePushQueueMulticastKey = "multicast"
const dataSourcePushQueueNameKey = "name"
const dataSourcePushQueueProjectIDKey = "project_id"
const dataSourcePushQueueRetriesDelayKey = "retries_delay"
const dataSourcePushQueueRetriesKey = "retries"
const dataSourcePushQueueSubscriberKey = "subscriber"
const dataSourcePushQueueURLKey = "url"

// dataSourcePushQueue reads information about IronMQ push queues.
func dataSourcePushQueue() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			dataSourcePushQueueErrorQueueKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of an error queue",
			},
			dataSourcePushQueueMessageCountKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of messages currently in the queue",
			},
			dataSourcePushQueueMessageCountTotalKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of messages which have been processed by the queue",
			},
			dataSourcePushQueueMulticastKey: &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether to create a multicast push queue",
			},
			dataSourcePushQueueNameKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the queue",
				ForceNew:    true,
			},
			dataSourcePushQueueProjectIDKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The project id",
				ForceNew:    true,
			},
			dataSourcePushQueueRetriesKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of retries before moving on to the next message",
			},
			dataSourcePushQueueRetriesDelayKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of seconds to wait before re-sending a failed message",
			},
			dataSourcePushQueueSubscriberKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						dataSourcePushQueueHeadersKey: {
							Type:     schema.TypeMap,
							Optional: true,
						},
						dataSourcePushQueueNameKey: {
							Type:     schema.TypeString,
							Optional: true,
						},
						dataSourcePushQueueURLKey: {
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

// dataSourcePushQueueRead reads information about an existing push queue.
func dataSourcePushQueueRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsMQ := config.Settings{}
	clientSettingsMQ.UseSettings(&clientSettings.MQ)

	projectID := d.Get(dataSourcePushQueueProjectIDKey).(string)
	queueName := d.Get(dataSourcePushQueueNameKey).(string)

	clientSettingsMQ.ProjectId = projectID

	queue := mq.ConfigNew(queueName, &clientSettingsMQ)
	queueInfo, err := queue.Info()

	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")

			return nil
		}

		return err
	}

	if queueInfo.Type == "pull" {
		d.SetId("")

		return nil
	}

	d.SetId(queueNameToID(clientSettingsMQ.ProjectId, queueName))

	return resourcePushQueueParseInfo(d, &queueInfo)
}
