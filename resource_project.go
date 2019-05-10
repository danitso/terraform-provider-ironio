package main

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iron-io/iron_go3/api"
	"github.com/iron-io/iron_go3/config"
)

// ProjectInfo describes a project.
type ProjectInfo struct {
	ID        string     `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	TenantID  int        `json:"tenant_id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Status    string     `json:"status,omitempty"`
	UserID    string     `json:"user_id,omitempty"`
}

// ProjectBody describes a project request payload.
type ProjectBody struct {
	Project ProjectInfo `json:"project,omitempty"`
	Message string      `json:"msg,omitempty"`
}

// resourceProject() manages projects.
func resourceProject() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
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

// resourceProjectCreate() creates a project.
func resourceProjectCreate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsAuth := config.Settings{}
	clientSettingsAuth.UseSettings(&clientSettings.Auth)

	in := ProjectInfo{
		Name: d.Get("name").(string),
	}

	var out ProjectBody

	url := resourceProjectGetEndpoint(clientSettingsAuth, "")
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

// resourceProjectGetEndpoint() returns an endpoint for a project.
func resourceProjectGetEndpoint(cs config.Settings, id string) *api.URL {
	u := &api.URL{Settings: cs, URL: url.URL{Scheme: cs.Scheme}}

	u.URL.Host = fmt.Sprintf("%s:%d", cs.Host, cs.Port)
	u.URL.Path = fmt.Sprintf("/%s/projects", cs.ApiVersion)

	if id != "" {
		u.URL.Path = fmt.Sprintf("%s/%s", u.URL.Path, id)
	}

	return u
}

// resourceProjectRead reads information about an existing project.
func resourceProjectRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsAuth := config.Settings{}
	clientSettingsAuth.UseSettings(&clientSettings.Auth)

	var out ProjectBody

	url := resourceProjectGetEndpoint(clientSettingsAuth, d.Id())
	err := url.Req("GET", nil, &out)

	if err != nil {
		if strings.Contains(err.Error(), " 404 ") {
			d.SetId("")

			return nil
		}
		return err
	}

	if out.Project.Name == "" {
		return errors.New("Failed to retrieve the project name")
	}

	d.Set("name", out.Project.Name)

	return nil
}

// resourceProjectUpdate updates an existing project.
func resourceProjectUpdate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsAuth := config.Settings{}
	clientSettingsAuth.UseSettings(&clientSettings.Auth)

	in := ProjectInfo{
		Name: d.Get("name").(string),
	}

	var out ProjectBody

	url := resourceProjectGetEndpoint(clientSettingsAuth, d.Id())
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

	url := resourceProjectGetEndpoint(clientSettingsAuth, d.Id())
	err := url.Req("DELETE", nil, &out)

	if err != nil {
		if !strings.Contains(err.Error(), " 404 ") {
			return err
		}
	}

	if out.Message != "success" {
		return fmt.Errorf("ERROR: Failed to delete the project due to an unknown error")
	}

	d.SetId("")

	return nil
}
