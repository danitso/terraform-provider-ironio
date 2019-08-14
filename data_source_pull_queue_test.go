/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package main

import (
	"testing"
)

// TestDataSourcePullQueueInstantiation tests whether the dataSourcePullQueue instance can be instantiated.
func TestDataSourcePullQueueInstantiation(t *testing.T) {
	s := dataSourcePullQueue()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourcePullQueue")
	}
}

// TestDataSourcePullQueueSchema tests the dataSourcePullQueue schema.
func TestDataSourcePullQueueSchema(t *testing.T) {
	s := dataSourcePullQueue()

	if s.Schema[dataSourcePullQueueMessageCountKey] == nil {
		t.Fatalf("Error in dataSourcePullQueue.Schema: Missing attribute \"%s\"", dataSourcePullQueueMessageCountKey)
	}

	if s.Schema[dataSourcePullQueueMessageCountKey].Computed != true {
		t.Fatalf("Error in dataSourcePullQueue.Schema: Attribute \"%s\" is not computed", dataSourcePullQueueMessageCountKey)
	}

	if s.Schema[dataSourcePullQueueMessageCountTotalKey] == nil {
		t.Fatalf("Error in dataSourcePullQueue.Schema: Missing attribute \"%s\"", dataSourcePullQueueMessageCountTotalKey)
	}

	if s.Schema[dataSourcePullQueueMessageCountTotalKey].Computed != true {
		t.Fatalf("Error in dataSourcePullQueue.Schema: Attribute \"%s\" is not computed", dataSourcePullQueueMessageCountTotalKey)
	}

	if s.Schema[dataSourcePullQueueNameKey] == nil {
		t.Fatalf("Error in dataSourcePullQueue.Schema: Missing argument \"%s\"", dataSourcePullQueueNameKey)
	}

	if s.Schema[dataSourcePullQueueNameKey].Required != true {
		t.Fatalf("Error in dataSourcePullQueue.Schema: Argument \"%s\" is not required", dataSourcePullQueueNameKey)
	}

	if s.Schema[dataSourcePullQueueProjectIDKey] == nil {
		t.Fatalf("Error in dataSourcePullQueue.Schema: Missing argument \"%s\"", dataSourcePullQueueProjectIDKey)
	}

	if s.Schema[dataSourcePullQueueProjectIDKey].Required != true {
		t.Fatalf("Error in dataSourcePullQueue.Schema: Argument \"%s\" is not required", dataSourcePullQueueNameKey)
	}
}
