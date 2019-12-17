/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package ironiotf

import (
	"time"

	"github.com/iron-io/iron_go3/config"
)

// ClientSettings contains the settings for each Iron.io product.
type ClientSettings struct {
	Auth   config.Settings
	Cache  config.Settings
	MQ     config.Settings
	Worker config.Settings
}

// ProjectBody describes a project payload.
type ProjectBody struct {
	Project ProjectInfo `json:"project,omitempty"`
	Message string      `json:"msg,omitempty"`
}

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

// ProjectListBody describes a project list payload.
type ProjectListBody struct {
	Projects []ProjectInfo `json:"projects,omitempty"`
	Message  string        `json:"msg,omitempty"`
}
