package main

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/version"
	"github.com/iron-io/iron_go3/config"
)

const ProviderConfigurationAuthKey = "auth"
const ProviderConfigurationCacheKey = "cache"
const ProviderConfigurationHostKey = "host"
const ProviderConfigurationLoadConfigFileKey = "load_config_file"
const ProviderConfigurationMQKey = "mq"
const ProviderConfigurationPortKey = "port"
const ProviderConfigurationProtocolKey = "protocol"
const ProviderConfigurationTokenKey = "token"
const ProviderConfigurationWorkerKey = "worker"

// ClientSettings contains the settings for each Iron.io product.
type ClientSettings struct {
	Auth   config.Settings
	Cache  config.Settings
	MQ     config.Settings
	Worker config.Settings
}

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
			ProviderConfigurationAuthKey: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						ProviderConfigurationHostKey: {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The IronAuth hostname or IP address",
						},
						ProviderConfigurationPortKey: {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The IronAuth port number",
						},
						ProviderConfigurationProtocolKey: {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The IronAuth protocol (HTTP or HTTPS)",
						},
						ProviderConfigurationTokenKey: {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The IronAuth token (OAuth)",
						},
					},
				},
				MaxItems: 1,
			},
			ProviderConfigurationCacheKey: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						ProviderConfigurationHostKey: {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The IronCache hostname or IP address",
						},
						ProviderConfigurationPortKey: {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The IronCache port number",
						},
						ProviderConfigurationProtocolKey: {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The IronCache protocol (HTTP or HTTPS)",
						},
					},
				},
				MaxItems: 1,
			},
			ProviderConfigurationLoadConfigFileKey: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to load the iron.json configuration file",
			},
			ProviderConfigurationMQKey: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						ProviderConfigurationHostKey: {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The IronMQ hostname or IP address",
						},
						ProviderConfigurationPortKey: {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The IronMQ port number",
						},
						ProviderConfigurationProtocolKey: {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The IronMQ protocol (HTTP or HTTPS)",
						},
					},
				},
				MaxItems: 1,
			},
			ProviderConfigurationWorkerKey: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						ProviderConfigurationHostKey: {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The IronWorker hostname or IP address",
						},
						ProviderConfigurationPortKey: {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The IronWorker port number",
						},
						ProviderConfigurationProtocolKey: {
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
	loadConfigFile := d.Get(ProviderConfigurationLoadConfigFileKey).(bool)

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
	auth := d.Get(ProviderConfigurationAuthKey).([]interface{})

	if len(auth) > 0 {
		authInfo := auth[0].(map[string]interface{})
		authHost := authInfo[ProviderConfigurationHostKey].(string)
		authPort := uint16(authInfo[ProviderConfigurationPortKey].(int))
		authProtocol := authInfo[ProviderConfigurationProtocolKey].(string)
		authToken := authInfo[ProviderConfigurationTokenKey].(string)

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
	cache := d.Get(ProviderConfigurationCacheKey).([]interface{})

	if len(cache) > 0 {
		cacheInfo := cache[0].(map[string]interface{})
		cacheHost := cacheInfo[ProviderConfigurationHostKey].(string)
		cachePort := uint16(cacheInfo[ProviderConfigurationPortKey].(int))
		cacheProtocol := cacheInfo[ProviderConfigurationProtocolKey].(string)

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
	mq := d.Get(ProviderConfigurationMQKey).([]interface{})

	if len(mq) > 0 {
		mqInfo := mq[0].(map[string]interface{})
		mqHost := mqInfo[ProviderConfigurationHostKey].(string)
		mqPort := uint16(mqInfo[ProviderConfigurationPortKey].(int))
		mqProtocol := mqInfo[ProviderConfigurationProtocolKey].(string)

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
	worker := d.Get(ProviderConfigurationWorkerKey).([]interface{})

	if len(worker) > 0 {
		workerInfo := worker[0].(map[string]interface{})
		workerHost := workerInfo[ProviderConfigurationHostKey].(string)
		workerPort := uint16(workerInfo[ProviderConfigurationPortKey].(int))
		workerProtocol := workerInfo[ProviderConfigurationProtocolKey].(string)

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
