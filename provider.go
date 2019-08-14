/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package main

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/version"
	"github.com/iron-io/iron_go3/config"
)

const providerConfigurationAuthKey = "auth"
const providerConfigurationCacheKey = "cache"
const providerConfigurationHostKey = "host"
const providerConfigurationLoadConfigFileKey = "load_config_file"
const providerConfigurationMQKey = "mq"
const providerConfigurationPortKey = "port"
const providerConfigurationProtocolKey = "protocol"
const providerConfigurationTokenKey = "token"
const providerConfigurationWorkerKey = "worker"

// Provider returns the object for this provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureFunc: providerConfigure,
		DataSourcesMap: map[string]*schema.Resource{
			"ironio_pull_queue": dataSourcePullQueue(),
			"ironio_push_queue": dataSourcePushQueue(),
			"ironio_projects":   dataSourceProjects(),
			"ironio_queues":     dataSourceQueues(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"ironio_project":    resourceProject(),
			"ironio_pull_queue": resourcePullQueue(),
			"ironio_push_queue": resourcePushQueue(),
		},
		Schema: map[string]*schema.Schema{
			providerConfigurationAuthKey: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						providerConfigurationHostKey: {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The IronAuth hostname or IP address",
						},
						providerConfigurationPortKey: {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The IronAuth port number",
						},
						providerConfigurationProtocolKey: {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The IronAuth protocol (HTTP or HTTPS)",
						},
						providerConfigurationTokenKey: {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The IronAuth token (OAuth)",
						},
					},
				},
				MaxItems: 1,
			},
			providerConfigurationCacheKey: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						providerConfigurationHostKey: {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The IronCache hostname or IP address",
						},
						providerConfigurationPortKey: {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The IronCache port number",
						},
						providerConfigurationProtocolKey: {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The IronCache protocol (HTTP or HTTPS)",
						},
					},
				},
				MaxItems: 1,
			},
			providerConfigurationLoadConfigFileKey: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to load the iron.json configuration file",
			},
			providerConfigurationMQKey: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						providerConfigurationHostKey: {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The IronMQ hostname or IP address",
						},
						providerConfigurationPortKey: {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The IronMQ port number",
						},
						providerConfigurationProtocolKey: {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The IronMQ protocol (HTTP or HTTPS)",
						},
					},
				},
				MaxItems: 1,
			},
			providerConfigurationWorkerKey: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						providerConfigurationHostKey: {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The IronWorker hostname or IP address",
						},
						providerConfigurationPortKey: {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The IronWorker port number",
						},
						providerConfigurationProtocolKey: {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The IronWorker protocol (HTTP or HTTPS)",
						},
					},
				},
				MaxItems: 1,
			},
		},
	}
}

// providerConfigure configures the provider before processing any IronMQ resources.
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	clientSettings := ClientSettings{
		Auth: config.Settings{
			Scheme:     "https",
			Port:       443,
			ApiVersion: "1",
			Host:       "auth.iron.io",
			UserAgent: fmt.Sprintf(
				"%s/%s Go/%s Terraform-Library/%s",
				TerraformProviderName,
				TerraformProviderVersion,
				runtime.Version(),
				version.Version,
			),
		},
	}
	loadConfigFile := d.Get(providerConfigurationLoadConfigFileKey).(bool)

	if loadConfigFile {
		// Use the settings stored in the configuration file or the environment variables.
		clientSettings.Cache = config.Config("iron_cache")
		clientSettings.MQ = config.Config("iron_mq")
		clientSettings.Worker = config.Config("iron_worker")
	} else {
		// Initialize the settings struct with the IronMQ preset.
		clientSettingsPresetCache := config.Presets["cache"]
		clientSettingsPresetMQ := config.Presets["mq"]
		clientSettingsPresetWorker := config.Presets["worker"]

		clientSettings.Cache.UseSettings(&clientSettingsPresetCache)
		clientSettings.MQ.UseSettings(&clientSettingsPresetMQ)
		clientSettings.Worker.UseSettings(&clientSettingsPresetWorker)
	}

	// Change the user agent to make it easier to analyze API requests.
	clientSettings.Cache.UserAgent = clientSettings.Auth.UserAgent
	clientSettings.MQ.UserAgent = clientSettings.Auth.UserAgent
	clientSettings.Worker.UserAgent = clientSettings.Auth.UserAgent

	// Modify the authentication settings based on the configuration values.
	auth := d.Get(providerConfigurationAuthKey).([]interface{})

	if len(auth) > 0 {
		authInfo := auth[0].(map[string]interface{})
		authHost := authInfo[providerConfigurationHostKey].(string)
		authPort := uint16(authInfo[providerConfigurationPortKey].(int))
		authProtocol := authInfo[providerConfigurationProtocolKey].(string)
		authToken := authInfo[providerConfigurationTokenKey].(string)

		if authHost != "" {
			clientSettings.Auth.Host = authHost
		}

		if authPort != 0 {
			clientSettings.Auth.Port = authPort
		}

		if authProtocol != "" {
			clientSettings.Auth.Scheme = authProtocol
		}

		if authToken != "" {
			clientSettings.Auth.Token = authToken
			clientSettings.Cache.Token = clientSettings.Auth.Token
			clientSettings.MQ.Token = clientSettings.Auth.Token
			clientSettings.Worker.Token = clientSettings.Auth.Token
		}
	}

	// Modify the cache settings based on the configuration values.
	cache := d.Get(providerConfigurationCacheKey).([]interface{})

	if len(cache) > 0 {
		cacheInfo := cache[0].(map[string]interface{})
		cacheHost := cacheInfo[providerConfigurationHostKey].(string)
		cachePort := uint16(cacheInfo[providerConfigurationPortKey].(int))
		cacheProtocol := cacheInfo[providerConfigurationProtocolKey].(string)

		if cacheHost != "" {
			clientSettings.Cache.Host = cacheHost
		}

		if cachePort != 0 {
			clientSettings.Cache.Port = cachePort
		}

		if cacheProtocol != "" {
			clientSettings.Cache.Scheme = cacheProtocol
		}
	}

	// Modify the MQ settings based on the configuration values.
	mq := d.Get(providerConfigurationMQKey).([]interface{})

	if len(mq) > 0 {
		mqInfo := mq[0].(map[string]interface{})
		mqHost := mqInfo[providerConfigurationHostKey].(string)
		mqPort := uint16(mqInfo[providerConfigurationPortKey].(int))
		mqProtocol := mqInfo[providerConfigurationProtocolKey].(string)

		if mqHost != "" {
			clientSettings.MQ.Host = mqHost
		}

		if mqPort != 0 {
			clientSettings.MQ.Port = mqPort
		}

		if mqProtocol != "" {
			clientSettings.MQ.Scheme = mqProtocol
		}
	}

	// Modify the worker settings based on the configuration values.
	worker := d.Get(providerConfigurationWorkerKey).([]interface{})

	if len(worker) > 0 {
		workerInfo := worker[0].(map[string]interface{})
		workerHost := workerInfo[providerConfigurationHostKey].(string)
		workerPort := uint16(workerInfo[providerConfigurationPortKey].(int))
		workerProtocol := workerInfo[providerConfigurationProtocolKey].(string)

		if workerHost != "" {
			clientSettings.Worker.Host = workerHost
		}

		if workerPort != 0 {
			clientSettings.Worker.Port = workerPort
		}

		if workerProtocol != "" {
			clientSettings.Worker.Scheme = workerProtocol
		}
	}

	// Verify that the token has been specified for all services as we cannot proceed without it.
	if clientSettings.Auth.Token == "" {
		return clientSettings, errors.New("The token for IronAuth is undefined")
	} else if clientSettings.Cache.Token == "" {
		return clientSettings, errors.New("The token for IronCache is undefined")
	} else if clientSettings.MQ.Token == "" {
		return clientSettings, errors.New("The token for IronMQ is undefined")
	} else if clientSettings.Worker.Token == "" {
		return clientSettings, errors.New("The token for IronWorker is undefined")
	}

	return clientSettings, nil
}
