package main

import (
	"testing"
)

// TestResourceProjectInstantiation tests whether the resourceProject instance can be instantiated.
func TestResourceProjectInstantiation(t *testing.T) {
	s := resourceProject()

	if s == nil {
		t.Fatalf("Cannot instantiate resourceProject")
	}
}

// TestResourceProjectSchema tests the resourceProject schema.
func TestResourceProjectSchema(t *testing.T) {
	s := resourceProject()

	if s.Schema[resourceProjectNameKey] == nil {
		t.Fatalf("Error in resourceProject.Schema: Missing argument \"%s\"", resourceProjectNameKey)
	}

	if s.Schema[resourceProjectNameKey].Required != true {
		t.Fatalf("Error in resourceProject.Schema: Argument \"%s\" is not required", resourceProjectNameKey)
	}
}
