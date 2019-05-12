package main

import (
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iron-io/iron_go3/config"
	"github.com/iron-io/iron_go3/mq"
)

// resourcePullQueue() manages IronMQ pull queues.
func resourcePullQueue() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the queue",
				ForceNew:    true,
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
			"project_id": &schema.Schema{
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

	projectID := d.Get("project_id").(string)
	queueName := d.Get("name").(string)

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
		}

		return err
	}

	if queueInfo.Type != "pull" {
		d.SetId("")

		return nil
	}

	d.Set("message_count", queueInfo.Size)
	d.Set("message_count_total", queueInfo.TotalMessages)

	return nil
}

// resourcePullQueueDelete() deletes an existing pull queue.
func resourcePullQueueDelete(d *schema.ResourceData, m interface{}) error {
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
