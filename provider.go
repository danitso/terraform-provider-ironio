package main

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/version"
	"github.com/iron-io/iron_go3/config"
)

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
			"ironio_projects": dataSourceProjects(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"ironio_project":    resourceProject(),
			"ironio_pull_queue": resourcePullQueue(),
			"ironio_push_queue": resourcePushQueue(),
		},
		Schema: map[string]*schema.Schema{
			"auth_host": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The IronAuth hostname or IP address",
			},
			"auth_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "The IronAuth port number",
			},
			"auth_protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The IronAuth protocol (HTTP or HTTPS)",
			},
			"cache_host": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The IronCache hostname or IP address",
			},
			"cache_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "The IronCache port number",
			},
			"cache_protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The IronCache protocol (HTTP or HTTPS)",
			},
			"load_config_file": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to load the iron.json configuration file",
			},
			"mq_host": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The IronMQ hostname or IP address",
			},
			"mq_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "The IronMQ port number",
			},
			"mq_protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The IronMQ protocol (HTTP or HTTPS)",
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The token",
			},
			"worker_host": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The IronWorker hostname or IP address",
			},
			"worker_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "The IronWorker port number",
			},
			"worker_protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The IronWorker protocol (HTTP or HTTPS)",
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
	loadConfigFile := d.Get("load_config_file").(bool)

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

	// Modify the authentication settings based on the configuration values.
	authHost := d.Get("auth_host").(string)
	authPort := uint16(d.Get("auth_port").(int))
	authProtocol := d.Get("auth_protocol").(string)

	if authHost != "" {
		clientSettings.Auth.Host = authHost
	}

	if authPort != 0 {
		clientSettings.Auth.Port = authPort
	}

	if authProtocol != "" {
		clientSettings.Auth.Scheme = authProtocol
	}

	// Modify the cache settings based on the configuration values.
	cacheHost := d.Get("cache_host").(string)
	cachePort := uint16(d.Get("cache_port").(int))
	cacheProtocol := d.Get("cache_protocol").(string)

	if cacheHost != "" {
		clientSettings.Cache.Host = cacheHost
	}

	if cachePort != 0 {
		clientSettings.Cache.Port = cachePort
	}

	if cacheProtocol != "" {
		clientSettings.Cache.Scheme = cacheProtocol
	}

	// Modify the MQ settings based on the configuration values.
	mqHost := d.Get("mq_host").(string)
	mqPort := uint16(d.Get("mq_port").(int))
	mqProtocol := d.Get("mq_protocol").(string)

	if mqHost != "" {
		clientSettings.MQ.Host = mqHost
	}

	if mqPort != 0 {
		clientSettings.MQ.Port = mqPort
	}

	if mqProtocol != "" {
		clientSettings.MQ.Scheme = mqProtocol
	}

	// Modify the worker settings based on the configuration values.
	workerHost := d.Get("worker_host").(string)
	workerPort := uint16(d.Get("worker_port").(int))
	workerProtocol := d.Get("worker_protocol").(string)

	if workerHost != "" {
		clientSettings.MQ.Host = workerHost
	}

	if workerPort != 0 {
		clientSettings.MQ.Port = workerPort
	}

	if workerProtocol != "" {
		clientSettings.MQ.Scheme = workerProtocol
	}

	// Specify common values for all the services.
	token := d.Get("token").(string)

	if token != "" {
		clientSettings.Auth.Token = token
		clientSettings.Cache.Token = token
		clientSettings.MQ.Token = token
		clientSettings.Worker.Token = token
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

	// Change the user agent to make it easier to analyze API requests.
	clientSettings.Cache.UserAgent = clientSettings.Auth.UserAgent
	clientSettings.MQ.UserAgent = clientSettings.Auth.UserAgent
	clientSettings.Worker.UserAgent = clientSettings.Auth.UserAgent

	return clientSettings, nil
}
