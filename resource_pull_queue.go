package main

import (
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iron-io/iron_go3/config"
	"github.com/iron-io/iron_go3/mq"
)

const ResourcePullQueueMessageCountKey = "message_count"
const ResourcePullQueueMessageCountTotalKey = "message_count_total"
const ResourcePullQueueNameKey = "name"
const ResourcePullQueueProjectIDKey = "project_id"

// resourcePullQueue() manages IronMQ pull queues.
func resourcePullQueue() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			ResourcePullQueueMessageCountKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of messages currently in the queue",
			},
			ResourcePullQueueMessageCountTotalKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of messages which have been processed by the queue",
			},
			ResourcePullQueueNameKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the queue",
				ForceNew:    true,
			},
			ResourcePullQueueProjectIDKey: &schema.Schema{
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

// resourcePullQueueCreate() creates a new pull queue.
func resourcePullQueueCreate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsMQ := config.Settings{}
	clientSettingsMQ.UseSettings(&clientSettings.MQ)

	projectID := d.Get(ResourcePullQueueProjectIDKey).(string)
	queueName := d.Get(ResourcePullQueueNameKey).(string)

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

	projectID := d.State().Attributes[ResourcePullQueueProjectIDKey]
	queueName := d.State().Attributes[ResourcePullQueueNameKey]

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

	d.Set(ResourcePullQueueMessageCountKey, queueInfo.Size)
	d.Set(ResourcePullQueueMessageCountTotalKey, queueInfo.TotalMessages)

	return nil
}

// resourcePullQueueDelete() deletes an existing pull queue.
func resourcePullQueueDelete(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsMQ := config.Settings{}
	clientSettingsMQ.UseSettings(&clientSettings.MQ)

	projectID := d.State().Attributes[ResourcePullQueueProjectIDKey]
	queueName := d.State().Attributes[ResourcePullQueueNameKey]

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
