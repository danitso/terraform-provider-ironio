package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

// TestDataSourcePushQueueInstantiation() tests whether the dataSourcePushQueue instance can be instantiated.
func TestDataSourcePushQueueInstantiation(t *testing.T) {
	s := dataSourcePushQueue()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourcePushQueue")
	}
}

// TestDataSourcePushQueueSchema() tests the dataSourcePushQueue schema.
func TestDataSourcePushQueueSchema(t *testing.T) {
	s := dataSourcePushQueue()

	if s.Schema[DataSourcePushQueueErrorQueueKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Missing attribute \"%s\"", DataSourcePushQueueErrorQueueKey)
	}

	if s.Schema[DataSourcePushQueueErrorQueueKey].Computed != true {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Attribute \"%s\" is not computed", DataSourcePushQueueErrorQueueKey)
	}

	if s.Schema[DataSourcePushQueueMessageCountKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Missing attribute \"%s\"", DataSourcePushQueueMessageCountKey)
	}

	if s.Schema[DataSourcePushQueueMessageCountKey].Computed != true {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Attribute \"%s\" is not computed", DataSourcePushQueueMessageCountKey)
	}

	if s.Schema[DataSourcePushQueueMessageCountTotalKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Missing attribute \"%s\"", DataSourcePushQueueMessageCountTotalKey)
	}

	if s.Schema[DataSourcePushQueueMessageCountTotalKey].Computed != true {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Attribute \"%s\" is not computed", DataSourcePushQueueMessageCountTotalKey)
	}

	if s.Schema[DataSourcePushQueueMulticastKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Missing attribute \"%s\"", DataSourcePushQueueMulticastKey)
	}

	if s.Schema[DataSourcePushQueueMulticastKey].Computed != true {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Attribute \"%s\" is not computed", DataSourcePushQueueMulticastKey)
	}

	if s.Schema[DataSourcePushQueueMulticastKey].Type != schema.TypeBool {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Argument \"%s\" is not limited to booleans", DataSourcePushQueueMulticastKey)
	}

	if s.Schema[DataSourcePushQueueNameKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Missing argument \"%s\"", DataSourcePushQueueNameKey)
	}

	if s.Schema[DataSourcePushQueueNameKey].Required != true {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Argument \"%s\" is not required", DataSourcePushQueueNameKey)
	}

	if s.Schema[DataSourcePushQueueProjectIDKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Missing argument \"%s\"", DataSourcePushQueueProjectIDKey)
	}

	if s.Schema[DataSourcePushQueueProjectIDKey].Required != true {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Argument \"%s\" is not required", DataSourcePushQueueNameKey)
	}

	if s.Schema[DataSourcePushQueueRetriesKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Missing attribute \"%s\"", DataSourcePushQueueRetriesKey)
	}

	if s.Schema[DataSourcePushQueueRetriesKey].Computed != true {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Attribute \"%s\" is not optional", DataSourcePushQueueRetriesKey)
	}

	if s.Schema[DataSourcePushQueueRetriesKey].Type != schema.TypeInt {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Attribute \"%s\" is not limited to integers", DataSourcePushQueueRetriesKey)
	}

	if s.Schema[DataSourcePushQueueRetriesDelayKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Missing attribute \"%s\"", DataSourcePushQueueRetriesDelayKey)
	}

	if s.Schema[DataSourcePushQueueRetriesDelayKey].Computed != true {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Attribute \"%s\" is not computed", DataSourcePushQueueRetriesDelayKey)
	}

	if s.Schema[DataSourcePushQueueRetriesDelayKey].Type != schema.TypeInt {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Attribute \"%s\" is not limited to integers", DataSourcePushQueueRetriesDelayKey)
	}
}

// TestDataSourcePushQueueSchemaSubscriber() tests the dataSourcePushQueue.Subscriber schema.
func TestDataSourcePushQueueSchemaSubscriber(t *testing.T) {
	s := dataSourcePushQueue()

	if s.Schema[DataSourcePushQueueSubscriberKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Missing block \"%s\"", DataSourcePushQueueSubscriberKey)
	}

	if s.Schema[DataSourcePushQueueSubscriberKey].Computed != true {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Block \"%s\" is not computed", DataSourcePushQueueSubscriberKey)
	}

	if s.Schema[DataSourcePushQueueSubscriberKey].Type != schema.TypeList {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Block \"%s\" is not a list", DataSourcePushQueueSubscriberKey)
	}

	if s.Schema[DataSourcePushQueueSubscriberKey].Elem == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Missing element for block \"%s\"", DataSourcePushQueueSubscriberKey)
	}

	blockElement, blockElementCasted := s.Schema[DataSourcePushQueueSubscriberKey].Elem.(*schema.Resource)

	if !blockElementCasted {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Element for block \"%s\" is not a pointer to schema.Resource", DataSourcePushQueueSubscriberKey)
	}

	if blockElement.Schema[DataSourcePushQueueHeadersKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema.subscriber: Missing attribute \"%s\"", DataSourcePushQueueHeadersKey)
	}

	if blockElement.Schema[DataSourcePushQueueHeadersKey].Type != schema.TypeMap {
		t.Fatalf("Error in dataSourcePushQueue.Schema.subscriber: Attribute \"%s\" is not limited to maps", DataSourcePushQueueHeadersKey)
	}

	if blockElement.Schema[DataSourcePushQueueNameKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema.subscriber: Missing attribute \"%s\"", DataSourcePushQueueNameKey)
	}

	if blockElement.Schema[DataSourcePushQueueURLKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema.subscriber: Missing attribute \"%s\"", DataSourcePushQueueURLKey)
	}
}
