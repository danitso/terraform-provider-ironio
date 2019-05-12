package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

// TestDataSourceQueuesInstantiation() tests whether the dataSourceQueues instance can be instantiated.
func TestDataSourceQueuesInstantiation(t *testing.T) {
	s := dataSourceQueues()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceQueues")
	}
}

// TestDataSourceQueuesSchema() tests the dataSourceQueues schema.
func TestDataSourceQueuesSchema(t *testing.T) {
	s := dataSourceQueues()

	if s.Schema[DataSourceQueuesNamesKey] == nil {
		t.Fatalf("Error in dataSourceQueues.Schema: Missing attribute \"%s\"", DataSourceQueuesNamesKey)
	}

	if s.Schema[DataSourceQueuesNamesKey].Computed != true {
		t.Fatalf("Error in dataSourceQueues.Schema: Attribute \"%s\" is not computed", DataSourceQueuesNamesKey)
	}

	if s.Schema[DataSourceQueuesTypesKey] == nil {
		t.Fatalf("Error in dataSourceQueues.Schema: Missing attribute \"%s\"", DataSourceQueuesNamesKey)
	}

	if s.Schema[DataSourceQueuesTypesKey].Computed != true {
		t.Fatalf("Error in dataSourceQueues.Schema: Attribute \"%s\" is not computed", DataSourceQueuesNamesKey)
	}
}

// TestDataSourceQueuesSchemaFilter() tests the dataSourceQueues.Filter schema.
func TestDataSourceQueuesSchemaFilter(t *testing.T) {
	s := dataSourceQueues()

	if s.Schema[DataSourceQueuesFilterKey] == nil {
		t.Fatalf("Error in dataSourceQueues.Schema: Missing block \"%s\"", DataSourceQueuesFilterKey)
	}

	if s.Schema[DataSourceQueuesFilterKey].Optional != true {
		t.Fatalf("Error in dataSourceQueues.Schema: Block \"%s\" is not optional", DataSourceQueuesFilterKey)
	}

	if s.Schema[DataSourceQueuesFilterKey].Type != schema.TypeList {
		t.Fatalf("Error in dataSourceQueues.Schema: Block \"%s\" is not a list", DataSourceQueuesFilterKey)
	}

	if s.Schema[DataSourceQueuesFilterKey].MaxItems != 1 {
		t.Fatalf("Error in dataSourceQueues.Schema: Block \"%s\" is not limited to a single definition", DataSourceQueuesFilterKey)
	}

	if s.Schema[DataSourceQueuesFilterKey].Elem == nil {
		t.Fatalf("Error in dataSourceQueues.Schema: Missing element for block \"%s\"", DataSourceQueuesFilterKey)
	}

	blockElement, blockElementCasted := s.Schema[DataSourceQueuesFilterKey].Elem.(*schema.Resource)

	if !blockElementCasted {
		t.Fatalf("Error in dataSourceQueues.Schema: Element for block \"%s\" is not a pointer to schema.Resource", DataSourceQueuesFilterKey)
	}

	if blockElement.Schema[DataSourceQueuesNameKey] == nil {
		t.Fatalf("Error in dataSourceQueues.Schema.subscriber: Missing argument \"%s\"", DataSourceQueuesNameKey)
	}

	if blockElement.Schema[DataSourceQueuesNameKey].Optional != true {
		t.Fatalf("Error in dataSourceQueues.Schema.subscriber: Argument \"%s\" is not optional", DataSourceQueuesNameKey)
	}

	if blockElement.Schema[DataSourceQueuesPullKey] == nil {
		t.Fatalf("Error in dataSourceQueues.Schema.subscriber: Missing argument \"%s\"", DataSourceQueuesPullKey)
	}

	if blockElement.Schema[DataSourceQueuesPullKey].Optional != true {
		t.Fatalf("Error in dataSourceQueues.Schema.subscriber: Argument \"%s\" is not optional", DataSourceQueuesPullKey)
	}

	if blockElement.Schema[DataSourceQueuesPullKey].Type != schema.TypeBool {
		t.Fatalf("Error in dataSourceQueues.Schema.subscriber: Argument \"%s\" is not limited to booleans", DataSourceQueuesPullKey)
	}

	if blockElement.Schema[DataSourceQueuesPushKey] == nil {
		t.Fatalf("Error in dataSourceQueues.Schema.subscriber: Missing argument \"%s\"", DataSourceQueuesPushKey)
	}

	if blockElement.Schema[DataSourceQueuesPushKey].Optional != true {
		t.Fatalf("Error in dataSourceQueues.Schema.subscriber: Argument \"%s\" is not optional", DataSourceQueuesPushKey)
	}

	if blockElement.Schema[DataSourceQueuesPushKey].Type != schema.TypeBool {
		t.Fatalf("Error in dataSourceQueues.Schema.subscriber: Argument \"%s\" is not limited to booleans", DataSourceQueuesPushKey)
	}
}
