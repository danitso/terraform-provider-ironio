package main

import (
	"testing"
)

// TestDataSourcePullQueueInstantiation() tests whether the dataSourcePullQueue instance can be instantiated.
func TestDataSourcePullQueueInstantiation(t *testing.T) {
	s := dataSourcePullQueue()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourcePullQueue")
	}
}

// TestDataSourcePullQueueSchema() tests the dataSourcePullQueue schema.
func TestDataSourcePullQueueSchema(t *testing.T) {
	s := dataSourcePullQueue()

	if s.Schema[DataSourcePullQueueMessageCountKey] == nil {
		t.Fatalf("Error in dataSourcePullQueue.Schema: Missing attribute \"%s\"", DataSourcePullQueueMessageCountKey)
	}

	if s.Schema[DataSourcePullQueueMessageCountKey].Computed != true {
		t.Fatalf("Error in dataSourcePullQueue.Schema: Attribute \"%s\" is not computed", DataSourcePullQueueMessageCountKey)
	}

	if s.Schema[DataSourcePullQueueMessageCountTotalKey] == nil {
		t.Fatalf("Error in dataSourcePullQueue.Schema: Missing attribute \"%s\"", DataSourcePullQueueMessageCountTotalKey)
	}

	if s.Schema[DataSourcePullQueueMessageCountTotalKey].Computed != true {
		t.Fatalf("Error in dataSourcePullQueue.Schema: Attribute \"%s\" is not computed", DataSourcePullQueueMessageCountTotalKey)
	}

	if s.Schema[DataSourcePullQueueNameKey] == nil {
		t.Fatalf("Error in dataSourcePullQueue.Schema: Missing argument \"%s\"", DataSourcePullQueueNameKey)
	}

	if s.Schema[DataSourcePullQueueNameKey].Required != true {
		t.Fatalf("Error in dataSourcePullQueue.Schema: Argument \"%s\" is not required", DataSourcePullQueueNameKey)
	}

	if s.Schema[DataSourcePullQueueProjectIDKey] == nil {
		t.Fatalf("Error in dataSourcePullQueue.Schema: Missing argument \"%s\"", DataSourcePullQueueProjectIDKey)
	}

	if s.Schema[DataSourcePullQueueProjectIDKey].Required != true {
		t.Fatalf("Error in dataSourcePullQueue.Schema: Argument \"%s\" is not required", DataSourcePullQueueNameKey)
	}
}
