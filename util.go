package main

import (
	"strings"
)

// queueNameToID converts a queue name to an identifier.
func queueNameToID(projectID string, queueName string) string {
	id := projectID + "_" + queueName

	id = strings.ReplaceAll(id, " ", "_")
	id = strings.ReplaceAll(id, "-", "_")

	return id
}
