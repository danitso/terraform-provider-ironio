/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package ironiotf

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/iron-io/iron_go3/config"
	"github.com/iron-io/iron_go3/mq"
)

const resourcePullQueueMessageCountKey = "message_count"
const resourcePullQueueMessageCountTotalKey = "message_count_total"
const resourcePullQueueNameKey = "name"
const resourcePullQueueProjectIDKey = "project_id"

// resourcePullQueue manages IronMQ pull queues.
func resourcePullQueue() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			resourcePullQueueMessageCountKey: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of messages currently in the queue",
			},
			resourcePullQueueMessageCountTotalKey: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of messages which have been processed by the queue",
			},
			resourcePullQueueNameKey: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the queue",
				ForceNew:    true,
			},
			resourcePullQueueProjectIDKey: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The project id",
				ForceNew:    true,
			},
		},

		Create: resourcePullQueueCreate,
		Read:   resourcePullQueueRead,
		Delete: resourcePullQueueDelete,
	}
}

// resourcePullQueueCreate creates a new pull queue.
func resourcePullQueueCreate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsMQ := config.Settings{}
	clientSettingsMQ.UseSettings(&clientSettings.MQ)

	projectID := d.Get(resourcePullQueueProjectIDKey).(string)
	queueName := d.Get(resourcePullQueueNameKey).(string)

	clientSettingsMQ.ProjectId = projectID
	queueInfo := mq.QueueInfo{
		Name: queueName,
		Type: "pull",
	}
	_, err := mq.ConfigCreateQueue(queueInfo, &clientSettingsMQ)

	if err != nil {
		return err
	}

	d.SetId(queueNameToID(clientSettingsMQ.ProjectId, queueName))

	return resourcePullQueueRead(d, m)
}

// resourcePullQueueRead reads information about an existing pull queue.
func resourcePullQueueRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsMQ := config.Settings{}
	clientSettingsMQ.UseSettings(&clientSettings.MQ)

	projectID := d.State().Attributes[resourcePullQueueProjectIDKey]
	queueName := d.State().Attributes[resourcePullQueueNameKey]

	if projectID != "" {
		clientSettingsMQ.ProjectId = projectID
	}

	queue := mq.ConfigNew(queueName, &clientSettingsMQ)
	queueInfo, err := queue.Info()

	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")

			return nil
		}

		return err
	}

	if queueInfo.Type != "pull" {
		d.SetId("")

		return nil
	}

	d.Set(resourcePullQueueMessageCountKey, queueInfo.Size)
	d.Set(resourcePullQueueMessageCountTotalKey, queueInfo.TotalMessages)

	return nil
}

// resourcePullQueueDelete deletes an existing pull queue.
func resourcePullQueueDelete(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsMQ := config.Settings{}
	clientSettingsMQ.UseSettings(&clientSettings.MQ)

	projectID := d.State().Attributes[resourcePullQueueProjectIDKey]
	queueName := d.State().Attributes[resourcePullQueueNameKey]

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
