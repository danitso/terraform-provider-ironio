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
	client_settings, ok := m.(config.Settings)

	if !ok {
		return fmt.Errorf("Failed to retrieve the client settings")
	}

	queue_name := d.Get("name").(string)
	queue_info := mq.QueueInfo{
		Name: queue_name,
		Type: "pull",
	}
	_, err := mq.ConfigCreateQueue(queue_info, &client_settings)

	if err != nil {
		return err
	}

	d.SetId(queueNameToId(client_settings.ProjectId, queue_name))

	return resourcePullQueueRead(d, m)
}

// resourcePullQueueRead() reads information about an existing pull queue.
func resourcePullQueueRead(d *schema.ResourceData, m interface{}) error {
	client_settings, ok := m.(config.Settings)

	if !ok {
		return fmt.Errorf("Failed to retrieve the client settings")
	}

	queue_name := d.Get("name").(string)
	queue := mq.ConfigNew(queue_name, &client_settings)
	_, err := queue.Info()

	if err != nil {
		if strings.Contains(err.Error(), "Queue not found") {
			d.SetId("")

			return nil
		} else {
			return err
		}
	}

	return nil
}

// resourcePullQueueDelete() deletes an existing pull queue.
func resourcePullQueueDelete(d *schema.ResourceData, m interface{}) error {
	client_settings, ok := m.(config.Settings)

	if !ok {
		return fmt.Errorf("Failed to retrieve the client settings")
	}

	queue_name := d.Get("name").(string)
	queue := mq.ConfigNew(queue_name, &client_settings)
	err := queue.Delete()

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
