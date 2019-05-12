package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

// TestResourcePushQueueInstantiation() tests whether the resourcePushQueue instance can be instantiated.
func TestResourcePushQueueInstantiation(t *testing.T) {
	s := resourcePushQueue()

	if s == nil {
		t.Fatalf("Cannot instantiate resourcePushQueue")
	}
}

// TestResourcePushQueueSchema() tests the resourcePushQueue schema.
func TestResourcePushQueueSchema(t *testing.T) {
	s := resourcePushQueue()

	if s.Schema[ResourcePushQueueErrorQueueKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema: Missing argument \"%s\"", ResourcePushQueueErrorQueueKey)
	}

	if s.Schema[ResourcePushQueueErrorQueueKey].Optional != true {
		t.Fatalf("Error in resourcePushQueue.Schema: Argument \"%s\" is not optional", ResourcePushQueueErrorQueueKey)
	}

	if s.Schema[ResourcePushQueueMessageCountKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema: Missing attribute \"%s\"", ResourcePushQueueMessageCountKey)
	}

	if s.Schema[ResourcePushQueueMessageCountKey].Computed != true {
		t.Fatalf("Error in resourcePushQueue.Schema: Attribute \"%s\" is not computed", ResourcePushQueueMessageCountKey)
	}

	if s.Schema[ResourcePushQueueMessageCountTotalKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema: Missing attribute \"%s\"", ResourcePushQueueMessageCountTotalKey)
	}

	if s.Schema[ResourcePushQueueMessageCountTotalKey].Computed != true {
		t.Fatalf("Error in resourcePushQueue.Schema: Attribute \"%s\" is not computed", ResourcePushQueueMessageCountTotalKey)
	}

	if s.Schema[ResourcePushQueueMulticastKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema: Missing argument \"%s\"", ResourcePushQueueMulticastKey)
	}

	if s.Schema[ResourcePushQueueMulticastKey].Optional != true {
		t.Fatalf("Error in resourcePushQueue.Schema: Argument \"%s\" is not optional", ResourcePushQueueMulticastKey)
	}

	if s.Schema[ResourcePushQueueMulticastKey].Type != schema.TypeBool {
		t.Fatalf("Error in resourcePushQueue.Schema: Argument \"%s\" is not limited to booleans", ResourcePushQueueMulticastKey)
	}

	if s.Schema[ResourcePushQueueNameKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema: Missing argument \"%s\"", ResourcePushQueueNameKey)
	}

	if s.Schema[ResourcePushQueueNameKey].Required != true {
		t.Fatalf("Error in resourcePushQueue.Schema: Argument \"%s\" is not required", ResourcePushQueueNameKey)
	}

	if s.Schema[ResourcePushQueueProjectIDKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema: Missing argument \"%s\"", ResourcePushQueueProjectIDKey)
	}

	if s.Schema[ResourcePushQueueProjectIDKey].Required != true {
		t.Fatalf("Error in resourcePushQueue.Schema: Argument \"%s\" is not required", ResourcePushQueueNameKey)
	}

	if s.Schema[ResourcePushQueueRetriesKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema: Missing argument \"%s\"", ResourcePushQueueRetriesKey)
	}

	if s.Schema[ResourcePushQueueRetriesKey].Optional != true {
		t.Fatalf("Error in resourcePushQueue.Schema: Argument \"%s\" is not optional", ResourcePushQueueRetriesKey)
	}

	if s.Schema[ResourcePushQueueRetriesKey].Type != schema.TypeInt {
		t.Fatalf("Error in resourcePushQueue.Schema: Argument \"%s\" is not limited to integers", ResourcePushQueueRetriesKey)
	}

	if s.Schema[ResourcePushQueueRetriesDelayKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema: Missing argument \"%s\"", ResourcePushQueueRetriesDelayKey)
	}

	if s.Schema[ResourcePushQueueRetriesDelayKey].Optional != true {
		t.Fatalf("Error in resourcePushQueue.Schema: Argument \"%s\" is not optional", ResourcePushQueueRetriesDelayKey)
	}

	if s.Schema[ResourcePushQueueRetriesDelayKey].Type != schema.TypeInt {
		t.Fatalf("Error in resourcePushQueue.Schema: Argument \"%s\" is not limited to integers", ResourcePushQueueRetriesDelayKey)
	}
}

// TestResourcePushQueueSchemaSubscriber() tests the resourcePushQueue.Subscriber schema.
func TestResourcePushQueueSchemaSubscriber(t *testing.T) {
	s := resourcePushQueue()

	if s.Schema[ResourcePushQueueSubscriberKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema: Missing block \"%s\"", ResourcePushQueueSubscriberKey)
	}

	if s.Schema[ResourcePushQueueSubscriberKey].Required != true {
		t.Fatalf("Error in resourcePushQueue.Schema: Block \"%s\" is not required", ResourcePushQueueSubscriberKey)
	}

	if s.Schema[ResourcePushQueueSubscriberKey].Type != schema.TypeList {
		t.Fatalf("Error in resourcePushQueue.Schema: Block \"%s\" is not a list", ResourcePushQueueSubscriberKey)
	}

	if s.Schema[ResourcePushQueueSubscriberKey].MinItems != 1 {
		t.Fatalf("Error in resourcePushQueue.Schema: Block \"%s\" does not require any definitions", ResourcePushQueueSubscriberKey)
	}

	if s.Schema[ResourcePushQueueSubscriberKey].Elem == nil {
		t.Fatalf("Error in resourcePushQueue.Schema: Missing element for block \"%s\"", ResourcePushQueueSubscriberKey)
	}

	blockElement, blockElementCasted := s.Schema[ResourcePushQueueSubscriberKey].Elem.(*schema.Resource)

	if !blockElementCasted {
		t.Fatalf("Error in resourcePushQueue.Schema: Element for block \"%s\" is not a pointer to schema.Resource", ResourcePushQueueSubscriberKey)
	}

	if blockElement.Schema[ResourcePushQueueHeadersKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema.subscriber: Missing argument \"%s\"", ResourcePushQueueHeadersKey)
	}

	if blockElement.Schema[ResourcePushQueueHeadersKey].Optional != true {
		t.Fatalf("Error in resourcePushQueue.Schema.subscriber: Argument \"%s\" is not optional", ResourcePushQueueHeadersKey)
	}

	if blockElement.Schema[ResourcePushQueueHeadersKey].Type != schema.TypeMap {
		t.Fatalf("Error in resourcePushQueue.Schema.subscriber: Argument \"%s\" is not limited to maps", ResourcePushQueueHeadersKey)
	}

	if blockElement.Schema[ResourcePushQueueNameKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema.subscriber: Missing argument \"%s\"", ResourcePushQueueNameKey)
	}

	if blockElement.Schema[ResourcePushQueueNameKey].Optional != true {
		t.Fatalf("Error in resourcePushQueue.Schema.subscriber: Argument \"%s\" is not optional", ResourcePushQueueNameKey)
	}

	if blockElement.Schema[ResourcePushQueueURLKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema.subscriber: Missing argument \"%s\"", ResourcePushQueueURLKey)
	}

	if blockElement.Schema[ResourcePushQueueURLKey].Required != true {
		t.Fatalf("Error in resourcePushQueue.Schema.subscriber: Argument \"%s\" is not required", ResourcePushQueueURLKey)
	}
}
