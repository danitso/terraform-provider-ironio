package ironio

import (
	"fmt"
	"runtime"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/version"
	"github.com/iron-io/iron_go3/config"
)

// Provider() returns the object for this provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			"ironio_project":    resourceProject(),
			"ironio_pull_queue": resourcePullQueue(),
			"ironio_push_queue": resourcePushQueue(),
		},
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The cluster's host",
			},
			"load_config_file": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to ignore the provider properties and load the IronMQ configuration file instead",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "The cluster's port number",
			},
			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The project id",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The API protocol (HTTP or HTTPS)",
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The token",
			},
		},
	}
}

// providerConfigure() configures the provider before processing any IronMQ resources.
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var client_settings config.Settings

	load_config_file := d.Get("load_config_file").(bool)

	if load_config_file {
		// Use the settings stored in the configuration file or the environment variables.
		client_settings = config.Config("iron_mq")
	} else {
		// Initialize the settings struct with the IronMQ preset.
		preset_settings := config.Presets["mq"]
		client_settings.UseSettings(&preset_settings)

		// Retrieve the provider configuration and update the IronMQ settings accordingly.
		host := d.Get("host").(string)
		port := uint16(d.Get("port").(int))
		project_id := d.Get("project_id").(string)
		protocol := d.Get("protocol").(string)
		token := d.Get("token").(string)

		if host != "" {
			client_settings.Host = host
		}

		if port != 0 {
			client_settings.Port = port
		}

		if project_id != "" {
			client_settings.ProjectId = project_id
		} else {
			return nil, fmt.Errorf("The IronMQ project id is undefined")
		}

		if protocol != "" {
			client_settings.Scheme = protocol
		}

		if token != "" {
			client_settings.Token = token
		} else {
			return nil, fmt.Errorf("The IronMQ token is undefined")
		}
	}

	// Change the user agent in order to notify about the use of this provider.
	client_settings.UserAgent = fmt.Sprintf(
		"%s/%s Go/%s Terraform-Library/%s",
		TerraformProviderName,
		TerraformProviderVersion,
		runtime.Version(),
		version.Version,
	)

	return client_settings, nil
}
