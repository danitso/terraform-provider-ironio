/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

// TestResourcePushQueueInstantiation tests whether the resourcePushQueue instance can be instantiated.
func TestResourcePushQueueInstantiation(t *testing.T) {
	s := resourcePushQueue()

	if s == nil {
		t.Fatalf("Cannot instantiate resourcePushQueue")
	}
}

// TestResourcePushQueueSchema tests the resourcePushQueue schema.
func TestResourcePushQueueSchema(t *testing.T) {
	s := resourcePushQueue()

	if s.Schema[resourcePushQueueErrorQueueKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema: Missing argument \"%s\"", resourcePushQueueErrorQueueKey)
	}

	if s.Schema[resourcePushQueueErrorQueueKey].Optional != true {
		t.Fatalf("Error in resourcePushQueue.Schema: Argument \"%s\" is not optional", resourcePushQueueErrorQueueKey)
	}

	if s.Schema[resourcePushQueueMessageCountKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema: Missing attribute \"%s\"", resourcePushQueueMessageCountKey)
	}

	if s.Schema[resourcePushQueueMessageCountKey].Computed != true {
		t.Fatalf("Error in resourcePushQueue.Schema: Attribute \"%s\" is not computed", resourcePushQueueMessageCountKey)
	}

	if s.Schema[resourcePushQueueMessageCountTotalKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema: Missing attribute \"%s\"", resourcePushQueueMessageCountTotalKey)
	}

	if s.Schema[resourcePushQueueMessageCountTotalKey].Computed != true {
		t.Fatalf("Error in resourcePushQueue.Schema: Attribute \"%s\" is not computed", resourcePushQueueMessageCountTotalKey)
	}

	if s.Schema[resourcePushQueueMulticastKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema: Missing argument \"%s\"", resourcePushQueueMulticastKey)
	}

	if s.Schema[resourcePushQueueMulticastKey].Optional != true {
		t.Fatalf("Error in resourcePushQueue.Schema: Argument \"%s\" is not optional", resourcePushQueueMulticastKey)
	}

	if s.Schema[resourcePushQueueMulticastKey].Type != schema.TypeBool {
		t.Fatalf("Error in resourcePushQueue.Schema: Argument \"%s\" is not limited to booleans", resourcePushQueueMulticastKey)
	}

	if s.Schema[resourcePushQueueNameKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema: Missing argument \"%s\"", resourcePushQueueNameKey)
	}

	if s.Schema[resourcePushQueueNameKey].Required != true {
		t.Fatalf("Error in resourcePushQueue.Schema: Argument \"%s\" is not required", resourcePushQueueNameKey)
	}

	if s.Schema[resourcePushQueueProjectIDKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema: Missing argument \"%s\"", resourcePushQueueProjectIDKey)
	}

	if s.Schema[resourcePushQueueProjectIDKey].Required != true {
		t.Fatalf("Error in resourcePushQueue.Schema: Argument \"%s\" is not required", resourcePushQueueNameKey)
	}

	if s.Schema[resourcePushQueueRetriesKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema: Missing argument \"%s\"", resourcePushQueueRetriesKey)
	}

	if s.Schema[resourcePushQueueRetriesKey].Optional != true {
		t.Fatalf("Error in resourcePushQueue.Schema: Argument \"%s\" is not optional", resourcePushQueueRetriesKey)
	}

	if s.Schema[resourcePushQueueRetriesKey].Type != schema.TypeInt {
		t.Fatalf("Error in resourcePushQueue.Schema: Argument \"%s\" is not limited to integers", resourcePushQueueRetriesKey)
	}

	if s.Schema[resourcePushQueueRetriesDelayKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema: Missing argument \"%s\"", resourcePushQueueRetriesDelayKey)
	}

	if s.Schema[resourcePushQueueRetriesDelayKey].Optional != true {
		t.Fatalf("Error in resourcePushQueue.Schema: Argument \"%s\" is not optional", resourcePushQueueRetriesDelayKey)
	}

	if s.Schema[resourcePushQueueRetriesDelayKey].Type != schema.TypeInt {
		t.Fatalf("Error in resourcePushQueue.Schema: Argument \"%s\" is not limited to integers", resourcePushQueueRetriesDelayKey)
	}
}

// TestResourcePushQueueSchemaSubscriber tests the resourcePushQueue.Subscriber schema.
func TestResourcePushQueueSchemaSubscriber(t *testing.T) {
	s := resourcePushQueue()

	if s.Schema[resourcePushQueueSubscriberKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema: Missing block \"%s\"", resourcePushQueueSubscriberKey)
	}

	if s.Schema[resourcePushQueueSubscriberKey].Required != true {
		t.Fatalf("Error in resourcePushQueue.Schema: Block \"%s\" is not required", resourcePushQueueSubscriberKey)
	}

	if s.Schema[resourcePushQueueSubscriberKey].Type != schema.TypeList {
		t.Fatalf("Error in resourcePushQueue.Schema: Block \"%s\" is not a list", resourcePushQueueSubscriberKey)
	}

	if s.Schema[resourcePushQueueSubscriberKey].MinItems != 1 {
		t.Fatalf("Error in resourcePushQueue.Schema: Block \"%s\" does not require any definitions", resourcePushQueueSubscriberKey)
	}

	if s.Schema[resourcePushQueueSubscriberKey].Elem == nil {
		t.Fatalf("Error in resourcePushQueue.Schema: Missing element for block \"%s\"", resourcePushQueueSubscriberKey)
	}

	blockElement, blockElementCasted := s.Schema[resourcePushQueueSubscriberKey].Elem.(*schema.Resource)

	if !blockElementCasted {
		t.Fatalf("Error in resourcePushQueue.Schema: Element for block \"%s\" is not a pointer to schema.Resource", resourcePushQueueSubscriberKey)
	}

	if blockElement.Schema[resourcePushQueueHeadersKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema.subscriber: Missing argument \"%s\"", resourcePushQueueHeadersKey)
	}

	if blockElement.Schema[resourcePushQueueHeadersKey].Optional != true {
		t.Fatalf("Error in resourcePushQueue.Schema.subscriber: Argument \"%s\" is not optional", resourcePushQueueHeadersKey)
	}

	if blockElement.Schema[resourcePushQueueHeadersKey].Type != schema.TypeMap {
		t.Fatalf("Error in resourcePushQueue.Schema.subscriber: Argument \"%s\" is not limited to maps", resourcePushQueueHeadersKey)
	}

	if blockElement.Schema[resourcePushQueueNameKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema.subscriber: Missing argument \"%s\"", resourcePushQueueNameKey)
	}

	if blockElement.Schema[resourcePushQueueNameKey].Optional != true {
		t.Fatalf("Error in resourcePushQueue.Schema.subscriber: Argument \"%s\" is not optional", resourcePushQueueNameKey)
	}

	if blockElement.Schema[resourcePushQueueURLKey] == nil {
		t.Fatalf("Error in resourcePushQueue.Schema.subscriber: Missing argument \"%s\"", resourcePushQueueURLKey)
	}

	if blockElement.Schema[resourcePushQueueURLKey].Required != true {
		t.Fatalf("Error in resourcePushQueue.Schema.subscriber: Argument \"%s\" is not required", resourcePushQueueURLKey)
	}
}
