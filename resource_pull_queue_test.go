package main

import (
	"testing"
)

// TestResourcePullQueueInstantiation() tests whether the resourcePullQueue instance can be instantiated.
func TestResourcePullQueueInstantiation(t *testing.T) {
	s := resourcePullQueue()

	if s == nil {
		t.Fatalf("Cannot instantiate resourcePullQueue")
	}
}

// TestResourcePullQueueSchema() tests the resourcePullQueue schema.
func TestResourcePullQueueSchema(t *testing.T) {
	s := resourcePullQueue()

	if s.Schema[ResourcePullQueueMessageCountKey] == nil {
		t.Fatalf("Error in resourcePullQueue.Schema: Missing attribute \"%s\"", ResourcePullQueueMessageCountKey)
	}

	if s.Schema[ResourcePullQueueMessageCountKey].Computed != true {
		t.Fatalf("Error in resourcePullQueue.Schema: Attribute \"%s\" is not computed", ResourcePullQueueMessageCountKey)
	}

	if s.Schema[ResourcePullQueueMessageCountTotalKey] == nil {
		t.Fatalf("Error in resourcePullQueue.Schema: Missing attribute \"%s\"", ResourcePullQueueMessageCountTotalKey)
	}

	if s.Schema[ResourcePullQueueMessageCountTotalKey].Computed != true {
		t.Fatalf("Error in resourcePullQueue.Schema: Attribute \"%s\" is not computed", ResourcePullQueueMessageCountTotalKey)
	}

	if s.Schema[ResourcePullQueueNameKey] == nil {
		t.Fatalf("Error in resourcePullQueue.Schema: Missing argument \"%s\"", ResourcePullQueueNameKey)
	}

	if s.Schema[ResourcePullQueueNameKey].Required != true {
		t.Fatalf("Error in resourcePullQueue.Schema: Argument \"%s\" is not required", ResourcePullQueueNameKey)
	}

	if s.Schema[ResourcePullQueueProjectIDKey] == nil {
		t.Fatalf("Error in resourcePullQueue.Schema: Missing argument \"%s\"", ResourcePullQueueProjectIDKey)
	}

	if s.Schema[ResourcePullQueueProjectIDKey].Required != true {
		t.Fatalf("Error in resourcePullQueue.Schema: Argument \"%s\" is not required", ResourcePullQueueNameKey)
	}
}
