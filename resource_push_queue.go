/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package main

import (
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iron-io/iron_go3/config"
	"github.com/iron-io/iron_go3/mq"
)

const resourcePushQueueErrorQueueKey = "error_queue"
const resourcePushQueueHeadersKey = "headers"
const resourcePushQueueMessageCountKey = "message_count"
const resourcePushQueueMessageCountTotalKey = "message_count_total"
const resourcePushQueueMulticastKey = "multicast"
const resourcePushQueueNameKey = "name"
const resourcePushQueueProjectIDKey = "project_id"
const resourcePushQueueRetriesDelayKey = "retries_delay"
const resourcePushQueueRetriesKey = "retries"
const resourcePushQueueSubscriberKey = "subscriber"
const resourcePushQueueURLKey = "url"

// resourcePushQueue manages IronMQ push queues.
func resourcePushQueue() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			resourcePushQueueErrorQueueKey: &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of an error queue",
				Default:     "",
			},
			resourcePushQueueMessageCountKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of messages currently in the queue",
			},
			resourcePushQueueMessageCountTotalKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of messages which have been processed by the queue",
			},
			resourcePushQueueMulticastKey: &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to create a multicast push queue",
				ForceNew:    true,
				Default:     true,
			},
			resourcePushQueueNameKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the queue",
				ForceNew:    true,
			},
			resourcePushQueueProjectIDKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The project id",
				ForceNew:    true,
			},
			resourcePushQueueRetriesKey: &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of retries before moving on to the next message",
				Default:     3,
			},
			resourcePushQueueRetriesDelayKey: &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of seconds to wait before re-sending a failed message",
				Default:     60,
			},
			resourcePushQueueSubscriberKey: &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						resourcePushQueueHeadersKey: {
							Type:     schema.TypeMap,
							Optional: true,
						},
						resourcePushQueueNameKey: {
							Type:     schema.TypeString,
							Optional: true,
						},
						resourcePushQueueURLKey: {
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

// resourcePushQueueBuildInfo builds a queue information object.
func resourcePushQueueBuildInfo(d *schema.ResourceData) (mq.QueueInfo, error) {
	queueName := d.Get(resourcePushQueueNameKey).(string)
	queueType := "unicast"

	multicast := d.Get(resourcePushQueueMulticastKey).(bool)
	subscribers := d.Get(resourcePushQueueSubscriberKey).([]interface{})

	if multicast {
		queueType = "multicast"
	}

	pushInfo := mq.PushInfo{
		RetriesDelay: d.Get(resourcePushQueueRetriesDelayKey).(int),
		Retries:      d.Get(resourcePushQueueRetriesKey).(int),
		ErrorQueue:   d.Get(resourcePushQueueErrorQueueKey).(string),
	}

	for _, v := range subscribers {
		resourceData := v.(map[string]interface{})
		url := resourceData[resourcePushQueueURLKey].(string)

		if url == "" {
			continue
		}

		pushHeaders := map[string]string{}

		for hn, hv := range resourceData[resourcePushQueueHeadersKey].(map[string]interface{}) {
			pushHeaders[hn] = hv.(string)
		}

		subscriber := mq.QueueSubscriber{
			Name:    resourceData[resourcePushQueueNameKey].(string),
			URL:     url,
			Headers: pushHeaders,
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

// resourcePushQueueCreate creates a new push queue.
func resourcePushQueueCreate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsMQ := config.Settings{}
	clientSettingsMQ.UseSettings(&clientSettings.MQ)

	projectID := d.Get(resourcePushQueueProjectIDKey).(string)
	queueName := d.Get(resourcePushQueueNameKey).(string)

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

// resourcePushQueueParseInfo parses information about an existing push queue.
func resourcePushQueueParseInfo(d *schema.ResourceData, queueInfo *mq.QueueInfo) error {
	if queueInfo.Type == "multicast" {
		d.Set(resourcePushQueueMulticastKey, true)
	} else {
		d.Set(resourcePushQueueMulticastKey, false)
	}

	if queueInfo.Push != nil {
		d.Set(resourcePushQueueErrorQueueKey, queueInfo.Push.ErrorQueue)
		d.Set(resourcePushQueueRetriesKey, queueInfo.Push.Retries)
		d.Set(resourcePushQueueRetriesDelayKey, queueInfo.Push.RetriesDelay)

		if queueInfo.Push.Subscribers != nil {
			subscribers := make([]interface{}, len(queueInfo.Push.Subscribers))

			for k, v := range queueInfo.Push.Subscribers {
				subscriberMap := make(map[string]interface{})
				headersMap := make(map[string]interface{})

				for hk, hv := range v.Headers {
					headersMap[hk] = hv
				}

				subscriberMap[resourcePushQueueNameKey] = v.Name
				subscriberMap[resourcePushQueueURLKey] = v.URL
				subscriberMap[resourcePushQueueHeadersKey] = headersMap

				subscribers[k] = subscriberMap
			}

			d.Set(resourcePushQueueSubscriberKey, subscribers)
		}
	}

	d.Set(resourcePushQueueMessageCountKey, queueInfo.Size)
	d.Set(resourcePushQueueMessageCountTotalKey, queueInfo.TotalMessages)

	return nil
}

// resourcePushQueueRead reads information about an existing push queue.
func resourcePushQueueRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsMQ := config.Settings{}
	clientSettingsMQ.UseSettings(&clientSettings.MQ)

	projectID := d.State().Attributes[resourcePushQueueProjectIDKey]
	queueName := d.State().Attributes[resourcePushQueueNameKey]

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

	if queueInfo.Type == "pull" {
		d.SetId("")

		return nil
	}

	return resourcePushQueueParseInfo(d, &queueInfo)
}

// resourcePushQueueUpdate updates an existing push queue.
func resourcePushQueueUpdate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsMQ := config.Settings{}
	clientSettingsMQ.UseSettings(&clientSettings.MQ)

	projectID := d.Get(resourcePushQueueProjectIDKey).(string)
	queueName := d.Get(resourcePushQueueNameKey).(string)

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

// resourcePushQueueDelete deletes an existing push queue.
func resourcePushQueueDelete(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsMQ := config.Settings{}
	clientSettingsMQ.UseSettings(&clientSettings.MQ)

	projectID := d.State().Attributes[resourcePushQueueProjectIDKey]
	queueName := d.State().Attributes[resourcePushQueueNameKey]

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
