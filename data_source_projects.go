package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iron-io/iron_go3/config"
)

// dataSourceProjects() retrieves information about projects.
func dataSourceProjects() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"filter_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The name filter",
				ForceNew:    true,
			},
			"ids": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"names": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},

		Read: dataSourceProjectsRead,
	}
}

// dataSourceProjectsRead reads information about available projects.
func dataSourceProjectsRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	clientSettingsAuth := config.Settings{}
	clientSettingsAuth.UseSettings(&clientSettings.Auth)

	// Prepare the filters.
	filterName := d.Get("filter_name").(string)
	filterNameMode := 0

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

	// Retrieve the list of projects.
	var out ProjectListBody

	url := getProjectsURL(clientSettingsAuth, "")
	err := url.Req("GET", nil, &out)

	if err != nil {
		return err
	}

	if out.Message == "" {
		return errors.New("Failed to retrieve the project list")
	}

	// Parse and filter the results.
	ids := make([]string, 0)
	names := make([]string, 0)

	for _, v := range out.Projects {
		if filterNameMode == 1 && !strings.Contains(v.Name, filterName) {
			continue
		} else if filterNameMode == 2 && !strings.HasSuffix(v.Name, filterName) {
			continue
		} else if filterNameMode == 3 && !strings.HasPrefix(v.Name, filterName) {
			continue
		} else if filterNameMode == 4 && strings.Compare(v.Name, filterName) != 0 {
			continue
		}

		ids = append(ids, v.ID)
		names = append(names, v.Name)
	}

	h := sha256.New()
	h.Write([]byte(strings.Join(ids, ",")))

	d.SetId(fmt.Sprintf("%x", h.Sum(nil)))
	d.Set("ids", ids)
	d.Set("names", names)

	return nil
}
