package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/iron-io/iron_go3/api"
	"github.com/iron-io/iron_go3/config"
)

// getProjectsURL() returns a URL for all projects or a specific one.
func getProjectsURL(cs config.Settings, id string) *api.URL {
	u := &api.URL{Settings: cs, URL: url.URL{Scheme: cs.Scheme}}

	u.URL.Host = fmt.Sprintf("%s:%d", cs.Host, cs.Port)
	u.URL.Path = fmt.Sprintf("/%s/projects", cs.ApiVersion)

	if id != "" {
		u.URL.Path = fmt.Sprintf("%s/%s", u.URL.Path, id)
	}

	return u
}

// queueNameToID converts a queue name to an identifier.
func queueNameToID(projectID string, queueName string) string {
	id := projectID + "_" + queueName

	id = strings.ReplaceAll(id, " ", "_")
	id = strings.ReplaceAll(id, "-", "_")

	return id
}
