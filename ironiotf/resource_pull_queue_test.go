/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package ironiotf

import (
	"testing"
)

// TestResourcePullQueueInstantiation tests whether the resourcePullQueue instance can be instantiated.
func TestResourcePullQueueInstantiation(t *testing.T) {
	s := resourcePullQueue()

	if s == nil {
		t.Fatalf("Cannot instantiate resourcePullQueue")
	}
}

// TestResourcePullQueueSchema tests the resourcePullQueue schema.
func TestResourcePullQueueSchema(t *testing.T) {
	s := resourcePullQueue()

	if s.Schema[resourcePullQueueMessageCountKey] == nil {
		t.Fatalf("Error in resourcePullQueue.Schema: Missing attribute \"%s\"", resourcePullQueueMessageCountKey)
	}

	if s.Schema[resourcePullQueueMessageCountKey].Computed != true {
		t.Fatalf("Error in resourcePullQueue.Schema: Attribute \"%s\" is not computed", resourcePullQueueMessageCountKey)
	}

	if s.Schema[resourcePullQueueMessageCountTotalKey] == nil {
		t.Fatalf("Error in resourcePullQueue.Schema: Missing attribute \"%s\"", resourcePullQueueMessageCountTotalKey)
	}

	if s.Schema[resourcePullQueueMessageCountTotalKey].Computed != true {
		t.Fatalf("Error in resourcePullQueue.Schema: Attribute \"%s\" is not computed", resourcePullQueueMessageCountTotalKey)
	}

	if s.Schema[resourcePullQueueNameKey] == nil {
		t.Fatalf("Error in resourcePullQueue.Schema: Missing argument \"%s\"", resourcePullQueueNameKey)
	}

	if s.Schema[resourcePullQueueNameKey].Required != true {
		t.Fatalf("Error in resourcePullQueue.Schema: Argument \"%s\" is not required", resourcePullQueueNameKey)
	}

	if s.Schema[resourcePullQueueProjectIDKey] == nil {
		t.Fatalf("Error in resourcePullQueue.Schema: Missing argument \"%s\"", resourcePullQueueProjectIDKey)
	}

	if s.Schema[resourcePullQueueProjectIDKey].Required != true {
		t.Fatalf("Error in resourcePullQueue.Schema: Argument \"%s\" is not required", resourcePullQueueNameKey)
	}
}
