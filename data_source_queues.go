package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iron-io/iron_go3/config"
	"github.com/iron-io/iron_go3/mq"
)

const dataSourceQueuesFilterKey = "filter"
const dataSourceQueuesNameKey = "name"
const dataSourceQueuesNamesKey = "names"
const dataSourceQueuesPullKey = "pull"
const dataSourceQueuesPushKey = "push"
const dataSourceQueuesProjectIDKey = "project_id"
const dataSourceQueuesTypesKey = "types"

// dataSourceQueues retrieves information about queues.
func dataSourceQueues() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			dataSourceQueuesFilterKey: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						dataSourceQueuesNameKey: &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The name filter",
							ForceNew:    true,
						},
						dataSourceQueuesPullKey: &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "Whether to include pull queues",
							ForceNew:    true,
						},
						dataSourceQueuesPushKey: &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "Whether to include push queues",
							ForceNew:    true,
						},
					},
				},
				MaxItems: 1,
			},
			dataSourceQueuesNamesKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			dataSourceQueuesProjectIDKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The project id",
				ForceNew:    true,
			},
			dataSourceQueuesTypesKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},

		Read: dataSourceQueuesRead,
	}
}

// dataSourceQueuesRead reads information about available queues.
func dataSourceQueuesRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsMQ := config.Settings{}
	clientSettingsMQ.UseSettings(&clientSettings.MQ)
	clientSettingsMQ.ProjectId = d.Get(dataSourceQueuesProjectIDKey).(string)

	// Prepare the filters.
	filter := d.Get(dataSourceQueuesFilterKey).([]interface{})
	filterName := ""
	filterNameMode := 0
	filterPull := true
	filterPush := true

	if len(filter) > 0 {
		filterData := filter[0].(map[string]interface{})
		filterName = filterData[dataSourceQueuesNameKey].(string)
		filterPull = filterData[dataSourceQueuesPullKey].(bool)
		filterPush = filterData[dataSourceQueuesPushKey].(bool)

		if filterName != "" {
			if len(filterName) >= 2 && strings.HasPrefix(filterName, "*") && strings.HasSuffix(filterName, "*") {
				filterName = filterName[1 : len(filterName)-1]
				filterNameMode = 1
			} else if strings.HasPrefix(filterName, "*") {
				filterName = filterName[1:len(filterName)]
				filterNameMode = 2
			} else if strings.HasSuffix(filterName, "*") {
				filterName = filterName[0 : len(filterName)-1]
				filterNameMode = 3
			} else {
				filterNameMode = 4
			}

			if filterNameMode > 0 && filterName == "" {
				return errors.New("The name filter cannot be an empty wildcard filter")
			}
		}
	}

	// Retrieve the list of projects.
	queues, err := mq.ListQueues(clientSettingsMQ, "", "", 1000)

	if err != nil {
		return err
	}

	// Parse and filter the results.
	names := make([]string, 0)
	types := make([]string, 0)

	for _, v := range queues {
		if filterNameMode == 1 && !strings.Contains(v.Name, filterName) {
			continue
		} else if filterNameMode == 2 && !strings.HasSuffix(v.Name, filterName) {
			continue
		} else if filterNameMode == 3 && !strings.HasPrefix(v.Name, filterName) {
			continue
		} else if filterNameMode == 4 && strings.Compare(v.Name, filterName) != 0 {
			continue
		}

		queueInfo, errInfo := v.Info()

		if errInfo != nil {
			return errInfo
		}

		if !filterPull && queueInfo.Type == "pull" {
			continue
		} else if !filterPush && queueInfo.Type != "pull" {
			continue
		}

		names = append(names, v.Name)

		if queueInfo.Type == "pull" {
			types = append(types, queueInfo.Type)
		} else {
			types = append(types, "push")
		}
	}

	h := sha256.New()
	h.Write([]byte(strings.Join(names, ",")))

	d.Set(dataSourceQueuesNamesKey, names)
	d.Set(dataSourceQueuesTypesKey, types)

	d.SetId(fmt.Sprintf("%x", h.Sum(nil)))

	return nil
}
