package main

import (
	"time"
)

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
