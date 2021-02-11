/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package ironiotf

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// TestDataSourcePushQueueInstantiation tests whether the dataSourcePushQueue instance can be instantiated.
func TestDataSourcePushQueueInstantiation(t *testing.T) {
	s := dataSourcePushQueue()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourcePushQueue")
	}
}

// TestDataSourcePushQueueSchema tests the dataSourcePushQueue schema.
func TestDataSourcePushQueueSchema(t *testing.T) {
	s := dataSourcePushQueue()

	if s.Schema[dataSourcePushQueueErrorQueueKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Missing attribute \"%s\"", dataSourcePushQueueErrorQueueKey)
	}

	if s.Schema[dataSourcePushQueueErrorQueueKey].Computed != true {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Attribute \"%s\" is not computed", dataSourcePushQueueErrorQueueKey)
	}

	if s.Schema[dataSourcePushQueueMessageCountKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Missing attribute \"%s\"", dataSourcePushQueueMessageCountKey)
	}

	if s.Schema[dataSourcePushQueueMessageCountKey].Computed != true {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Attribute \"%s\" is not computed", dataSourcePushQueueMessageCountKey)
	}

	if s.Schema[dataSourcePushQueueMessageCountTotalKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Missing attribute \"%s\"", dataSourcePushQueueMessageCountTotalKey)
	}

	if s.Schema[dataSourcePushQueueMessageCountTotalKey].Computed != true {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Attribute \"%s\" is not computed", dataSourcePushQueueMessageCountTotalKey)
	}

	if s.Schema[dataSourcePushQueueMulticastKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Missing attribute \"%s\"", dataSourcePushQueueMulticastKey)
	}

	if s.Schema[dataSourcePushQueueMulticastKey].Computed != true {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Attribute \"%s\" is not computed", dataSourcePushQueueMulticastKey)
	}

	if s.Schema[dataSourcePushQueueMulticastKey].Type != schema.TypeBool {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Argument \"%s\" is not limited to booleans", dataSourcePushQueueMulticastKey)
	}

	if s.Schema[dataSourcePushQueueNameKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Missing argument \"%s\"", dataSourcePushQueueNameKey)
	}

	if s.Schema[dataSourcePushQueueNameKey].Required != true {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Argument \"%s\" is not required", dataSourcePushQueueNameKey)
	}

	if s.Schema[dataSourcePushQueueProjectIDKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Missing argument \"%s\"", dataSourcePushQueueProjectIDKey)
	}

	if s.Schema[dataSourcePushQueueProjectIDKey].Required != true {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Argument \"%s\" is not required", dataSourcePushQueueNameKey)
	}

	if s.Schema[dataSourcePushQueueRetriesKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Missing attribute \"%s\"", dataSourcePushQueueRetriesKey)
	}

	if s.Schema[dataSourcePushQueueRetriesKey].Computed != true {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Attribute \"%s\" is not optional", dataSourcePushQueueRetriesKey)
	}

	if s.Schema[dataSourcePushQueueRetriesKey].Type != schema.TypeInt {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Attribute \"%s\" is not limited to integers", dataSourcePushQueueRetriesKey)
	}

	if s.Schema[dataSourcePushQueueRetriesDelayKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Missing attribute \"%s\"", dataSourcePushQueueRetriesDelayKey)
	}

	if s.Schema[dataSourcePushQueueRetriesDelayKey].Computed != true {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Attribute \"%s\" is not computed", dataSourcePushQueueRetriesDelayKey)
	}

	if s.Schema[dataSourcePushQueueRetriesDelayKey].Type != schema.TypeInt {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Attribute \"%s\" is not limited to integers", dataSourcePushQueueRetriesDelayKey)
	}
}

// TestDataSourcePushQueueSchemaSubscriber tests the dataSourcePushQueue.Subscriber schema.
func TestDataSourcePushQueueSchemaSubscriber(t *testing.T) {
	s := dataSourcePushQueue()

	if s.Schema[dataSourcePushQueueSubscriberKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Missing block \"%s\"", dataSourcePushQueueSubscriberKey)
	}

	if s.Schema[dataSourcePushQueueSubscriberKey].Computed != true {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Block \"%s\" is not computed", dataSourcePushQueueSubscriberKey)
	}

	if s.Schema[dataSourcePushQueueSubscriberKey].Type != schema.TypeList {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Block \"%s\" is not a list", dataSourcePushQueueSubscriberKey)
	}

	if s.Schema[dataSourcePushQueueSubscriberKey].Elem == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Missing element for block \"%s\"", dataSourcePushQueueSubscriberKey)
	}

	blockElement, blockElementCasted := s.Schema[dataSourcePushQueueSubscriberKey].Elem.(*schema.Resource)

	if !blockElementCasted {
		t.Fatalf("Error in dataSourcePushQueue.Schema: Element for block \"%s\" is not a pointer to schema.Resource", dataSourcePushQueueSubscriberKey)
	}

	if blockElement.Schema[dataSourcePushQueueHeadersKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema.subscriber: Missing attribute \"%s\"", dataSourcePushQueueHeadersKey)
	}

	if blockElement.Schema[dataSourcePushQueueHeadersKey].Type != schema.TypeMap {
		t.Fatalf("Error in dataSourcePushQueue.Schema.subscriber: Attribute \"%s\" is not limited to maps", dataSourcePushQueueHeadersKey)
	}

	if blockElement.Schema[dataSourcePushQueueNameKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema.subscriber: Missing attribute \"%s\"", dataSourcePushQueueNameKey)
	}

	if blockElement.Schema[dataSourcePushQueueURLKey] == nil {
		t.Fatalf("Error in dataSourcePushQueue.Schema.subscriber: Missing attribute \"%s\"", dataSourcePushQueueURLKey)
	}
}
