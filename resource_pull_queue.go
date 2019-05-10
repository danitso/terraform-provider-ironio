package ironio

import (
	"fmt"
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
		},

		Create: resourcePullQueueCreate,
		Read:   resourcePullQueueRead,
		Delete: resourcePullQueueDelete,
	}
}

// resourcePullQueueCreate() creates a new pull queue.
func resourcePullQueueCreate(d *schema.ResourceData, m interface{}) error {
	clientSettings, ok := m.(config.Settings)

	if !ok {
		return fmt.Errorf("Failed to retrieve the client settings")
	}

	queueName := d.Get("name").(string)
	queueInfo := mq.QueueInfo{
		Name: queueName,
		Type: "pull",
	}
	_, err := mq.ConfigCreateQueue(queueInfo, &clientSettings)

	if err != nil {
		return err
	}

	d.SetId(queueNameToID(clientSettings.ProjectId, queueName))

	return resourcePullQueueRead(d, m)
}

// resourcePullQueueRead reads information about an existing pull queue.
func resourcePullQueueRead(d *schema.ResourceData, m interface{}) error {
	clientSettings, ok := m.(config.Settings)

	if !ok {
		return fmt.Errorf("Failed to retrieve the client settings")
	}

	queueName := d.Get("name").(string)
	queue := mq.ConfigNew(queueName, &clientSettings)
	_, err := queue.Info()

	if err != nil {
		if strings.Contains(err.Error(), "Queue not found") {
			d.SetId("")

			return nil
		}
		return err
	}

	return nil
}

// resourcePullQueueDelete() deletes an existing pull queue.
func resourcePullQueueDelete(d *schema.ResourceData, m interface{}) error {
	clientSettings, ok := m.(config.Settings)

	if !ok {
		return fmt.Errorf("Failed to retrieve the client settings")
	}

	queueName := d.Get("name").(string)
	queue := mq.ConfigNew(queueName, &clientSettings)
	err := queue.Delete()

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
