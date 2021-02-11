/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package ironiotf

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// TestDataSourceQueuesInstantiation tests whether the dataSourceQueues instance can be instantiated.
func TestDataSourceQueuesInstantiation(t *testing.T) {
	s := dataSourceQueues()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceQueues")
	}
}

// TestDataSourceQueuesSchema tests the dataSourceQueues schema.
func TestDataSourceQueuesSchema(t *testing.T) {
	s := dataSourceQueues()

	if s.Schema[dataSourceQueuesNamesKey] == nil {
		t.Fatalf("Error in dataSourceQueues.Schema: Missing attribute \"%s\"", dataSourceQueuesNamesKey)
	}

	if s.Schema[dataSourceQueuesNamesKey].Computed != true {
		t.Fatalf("Error in dataSourceQueues.Schema: Attribute \"%s\" is not computed", dataSourceQueuesNamesKey)
	}

	if s.Schema[dataSourceQueuesTypesKey] == nil {
		t.Fatalf("Error in dataSourceQueues.Schema: Missing attribute \"%s\"", dataSourceQueuesNamesKey)
	}

	if s.Schema[dataSourceQueuesTypesKey].Computed != true {
		t.Fatalf("Error in dataSourceQueues.Schema: Attribute \"%s\" is not computed", dataSourceQueuesNamesKey)
	}
}

// TestDataSourceQueuesSchemaFilter tests the dataSourceQueues.Filter schema.
func TestDataSourceQueuesSchemaFilter(t *testing.T) {
	s := dataSourceQueues()

	if s.Schema[dataSourceQueuesFilterKey] == nil {
		t.Fatalf("Error in dataSourceQueues.Schema: Missing block \"%s\"", dataSourceQueuesFilterKey)
	}

	if s.Schema[dataSourceQueuesFilterKey].Optional != true {
		t.Fatalf("Error in dataSourceQueues.Schema: Block \"%s\" is not optional", dataSourceQueuesFilterKey)
	}

	if s.Schema[dataSourceQueuesFilterKey].Type != schema.TypeList {
		t.Fatalf("Error in dataSourceQueues.Schema: Block \"%s\" is not a list", dataSourceQueuesFilterKey)
	}

	if s.Schema[dataSourceQueuesFilterKey].MaxItems != 1 {
		t.Fatalf("Error in dataSourceQueues.Schema: Block \"%s\" is not limited to a single definition", dataSourceQueuesFilterKey)
	}

	if s.Schema[dataSourceQueuesFilterKey].Elem == nil {
		t.Fatalf("Error in dataSourceQueues.Schema: Missing element for block \"%s\"", dataSourceQueuesFilterKey)
	}

	blockElement, blockElementCasted := s.Schema[dataSourceQueuesFilterKey].Elem.(*schema.Resource)

	if !blockElementCasted {
		t.Fatalf("Error in dataSourceQueues.Schema: Element for block \"%s\" is not a pointer to schema.Resource", dataSourceQueuesFilterKey)
	}

	if blockElement.Schema[dataSourceQueuesNameKey] == nil {
		t.Fatalf("Error in dataSourceQueues.Schema.subscriber: Missing argument \"%s\"", dataSourceQueuesNameKey)
	}

	if blockElement.Schema[dataSourceQueuesNameKey].Optional != true {
		t.Fatalf("Error in dataSourceQueues.Schema.subscriber: Argument \"%s\" is not optional", dataSourceQueuesNameKey)
	}

	if blockElement.Schema[dataSourceQueuesPullKey] == nil {
		t.Fatalf("Error in dataSourceQueues.Schema.subscriber: Missing argument \"%s\"", dataSourceQueuesPullKey)
	}

	if blockElement.Schema[dataSourceQueuesPullKey].Optional != true {
		t.Fatalf("Error in dataSourceQueues.Schema.subscriber: Argument \"%s\" is not optional", dataSourceQueuesPullKey)
	}

	if blockElement.Schema[dataSourceQueuesPullKey].Type != schema.TypeBool {
		t.Fatalf("Error in dataSourceQueues.Schema.subscriber: Argument \"%s\" is not limited to booleans", dataSourceQueuesPullKey)
	}

	if blockElement.Schema[dataSourceQueuesPushKey] == nil {
		t.Fatalf("Error in dataSourceQueues.Schema.subscriber: Missing argument \"%s\"", dataSourceQueuesPushKey)
	}

	if blockElement.Schema[dataSourceQueuesPushKey].Optional != true {
		t.Fatalf("Error in dataSourceQueues.Schema.subscriber: Argument \"%s\" is not optional", dataSourceQueuesPushKey)
	}

	if blockElement.Schema[dataSourceQueuesPushKey].Type != schema.TypeBool {
		t.Fatalf("Error in dataSourceQueues.Schema.subscriber: Argument \"%s\" is not limited to booleans", dataSourceQueuesPushKey)
	}
}
