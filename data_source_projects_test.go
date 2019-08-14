/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

// TestDataSourceProjectsInstantiation tests whether the dataSourceProjects instance can be instantiated.
func TestDataSourceProjectsInstantiation(t *testing.T) {
	s := dataSourceProjects()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceProjects")
	}
}

// TestDataSourceProjectsSchema tests the dataSourceProjects schema.
func TestDataSourceProjectsSchema(t *testing.T) {
	s := dataSourceProjects()

	if s.Schema[dataSourceProjectsIdsKey] == nil {
		t.Fatalf("Error in dataSourceProjects.Schema: Missing attribute \"%s\"", dataSourceProjectsIdsKey)
	}

	if s.Schema[dataSourceProjectsIdsKey].Computed != true {
		t.Fatalf("Error in dataSourceProjects.Schema: Attribute \"%s\" is not computed", dataSourceProjectsIdsKey)
	}

	if s.Schema[dataSourceProjectsNamesKey] == nil {
		t.Fatalf("Error in dataSourceProjects.Schema: Missing attribute \"%s\"", dataSourceProjectsNamesKey)
	}

	if s.Schema[dataSourceProjectsNamesKey].Computed != true {
		t.Fatalf("Error in dataSourceProjects.Schema: Attribute \"%s\" is not computed", dataSourceProjectsNamesKey)
	}
}

// TestDataSourceProjectsSchemaFilter tests the dataSourceProjects.Filter schema.
func TestDataSourceProjectsSchemaFilter(t *testing.T) {
	s := dataSourceProjects()

	if s.Schema[dataSourceProjectsFilterKey] == nil {
		t.Fatalf("Error in dataSourceProjects.Schema: Missing block \"%s\"", dataSourceProjectsFilterKey)
	}

	if s.Schema[dataSourceProjectsFilterKey].Optional != true {
		t.Fatalf("Error in dataSourceProjects.Schema: Block \"%s\" is not optional", dataSourceProjectsFilterKey)
	}

	if s.Schema[dataSourceProjectsFilterKey].Type != schema.TypeList {
		t.Fatalf("Error in dataSourceProjects.Schema: Block \"%s\" is not a list", dataSourceProjectsFilterKey)
	}

	if s.Schema[dataSourceProjectsFilterKey].MaxItems != 1 {
		t.Fatalf("Error in dataSourceProjects.Schema: Block \"%s\" is not limited to a single definition", dataSourceProjectsFilterKey)
	}

	if s.Schema[dataSourceProjectsFilterKey].Elem == nil {
		t.Fatalf("Error in dataSourceProjects.Schema: Missing element for block \"%s\"", dataSourceProjectsFilterKey)
	}

	blockElement, blockElementCasted := s.Schema[dataSourceProjectsFilterKey].Elem.(*schema.Resource)

	if !blockElementCasted {
		t.Fatalf("Error in dataSourceProjects.Schema: Element for block \"%s\" is not a pointer to schema.Resource", dataSourceProjectsFilterKey)
	}

	if blockElement.Schema[dataSourceProjectsNameKey] == nil {
		t.Fatalf("Error in dataSourceProjects.Schema.subscriber: Missing argument \"%s\"", dataSourceProjectsNameKey)
	}

	if blockElement.Schema[dataSourceProjectsNameKey].Optional != true {
		t.Fatalf("Error in dataSourceProjects.Schema.subscriber: Argument \"%s\" is not optional", dataSourceProjectsNameKey)
	}
}
