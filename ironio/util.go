package ironio

import (
	"strings"
)

// queueNameToId() converts a queue name to an identifier.
func queueNameToId(projectId string, queueName string) string {
	id := projectId + "_" + queueName

	id = strings.ReplaceAll(id, " ", "_")
	id = strings.ReplaceAll(id, "-", "_")

	return id
}
