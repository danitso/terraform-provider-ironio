/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package ironiotf

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/iron-io/iron_go3/config"
)

const resourceProjectNameKey = "name"

// resourceProject manages projects.
func resourceProject() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			resourceProjectNameKey: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the project",
			},
		},

		Create: resourceProjectCreate,
		Read:   resourceProjectRead,
		Update: resourceProjectUpdate,
		Delete: resourceProjectDelete,
	}
}

// resourceProjectCreate creates a project.
func resourceProjectCreate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsAuth := config.Settings{}
	clientSettingsAuth.UseSettings(&clientSettings.Auth)

	in := ProjectInfo{
		Name: d.Get(resourceProjectNameKey).(string),
	}

	var out ProjectBody

	url := getProjectsURL(clientSettingsAuth, "")
	err := url.Req("POST", in, &out)

	if err != nil {
		return err
	}

	if out.Project.ID == "" {
		return fmt.Errorf("Failed to retrieve the project id for \"%s\"", in.Name)
	}

	d.SetId(out.Project.ID)

	return nil
}

// resourceProjectRead reads information about an existing project.
func resourceProjectRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsAuth := config.Settings{}
	clientSettingsAuth.UseSettings(&clientSettings.Auth)

	var out ProjectBody

	url := getProjectsURL(clientSettingsAuth, d.Id())
	err := url.Req("GET", nil, &out)

	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")

			return nil
		}

		return err
	}

	if out.Project.Name == "" {
		return errors.New("Failed to retrieve the project name")
	}

	d.Set(resourceProjectNameKey, out.Project.Name)

	return nil
}

// resourceProjectUpdate updates an existing project.
func resourceProjectUpdate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsAuth := config.Settings{}
	clientSettingsAuth.UseSettings(&clientSettings.Auth)

	in := ProjectInfo{
		Name: d.Get(resourceProjectNameKey).(string),
	}

	var out ProjectBody

	url := getProjectsURL(clientSettingsAuth, d.Id())
	err := url.Req("PATCH", in, &out)

	if err != nil {
		return err
	}

	return nil
}

// resourceProjectDelete deletes an existing project.
func resourceProjectDelete(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsAuth := config.Settings{}
	clientSettingsAuth.UseSettings(&clientSettings.Auth)

	var out ProjectBody

	url := getProjectsURL(clientSettingsAuth, d.Id())
	err := url.Req("DELETE", nil, &out)

	if err != nil {
		if !strings.Contains(err.Error(), "404") {
			return err
		}
	}

	d.SetId("")

	return nil
}
