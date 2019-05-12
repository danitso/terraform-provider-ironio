package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

// TestDataSourceProjectsInstantiation() tests whether the dataSourceProjects instance can be instantiated.
func TestDataSourceProjectsInstantiation(t *testing.T) {
	s := dataSourceProjects()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceProjects")
	}
}

// TestDataSourceProjectsSchema() tests the dataSourceProjects schema.
func TestDataSourceProjectsSchema(t *testing.T) {
	s := dataSourceProjects()

	if s.Schema[DataSourceProjectsIdsKey] == nil {
		t.Fatalf("Error in dataSourceProjects.Schema: Missing attribute \"%s\"", DataSourceProjectsIdsKey)
	}

	if s.Schema[DataSourceProjectsIdsKey].Computed != true {
		t.Fatalf("Error in dataSourceProjects.Schema: Attribute \"%s\" is not computed", DataSourceProjectsIdsKey)
	}

	if s.Schema[DataSourceProjectsNamesKey] == nil {
		t.Fatalf("Error in dataSourceProjects.Schema: Missing attribute \"%s\"", DataSourceProjectsNamesKey)
	}

	if s.Schema[DataSourceProjectsNamesKey].Computed != true {
		t.Fatalf("Error in dataSourceProjects.Schema: Attribute \"%s\" is not computed", DataSourceProjectsNamesKey)
	}
}

// TestDataSourceProjectsSchemaFilter() tests the dataSourceProjects.Filter schema.
func TestDataSourceProjectsSchemaFilter(t *testing.T) {
	s := dataSourceProjects()

	if s.Schema[DataSourceProjectsFilterKey] == nil {
		t.Fatalf("Error in dataSourceProjects.Schema: Missing block \"%s\"", DataSourceProjectsFilterKey)
	}

	if s.Schema[DataSourceProjectsFilterKey].Optional != true {
		t.Fatalf("Error in dataSourceProjects.Schema: Block \"%s\" is not optional", DataSourceProjectsFilterKey)
	}

	if s.Schema[DataSourceProjectsFilterKey].Type != schema.TypeList {
		t.Fatalf("Error in dataSourceProjects.Schema: Block \"%s\" is not a list", DataSourceProjectsFilterKey)
	}

	if s.Schema[DataSourceProjectsFilterKey].MaxItems != 1 {
		t.Fatalf("Error in dataSourceProjects.Schema: Block \"%s\" is not limited to a single definition", DataSourceProjectsFilterKey)
	}

	if s.Schema[DataSourceProjectsFilterKey].Elem == nil {
		t.Fatalf("Error in dataSourceProjects.Schema: Missing element for block \"%s\"", DataSourceProjectsFilterKey)
	}

	blockElement, blockElementCasted := s.Schema[DataSourceProjectsFilterKey].Elem.(*schema.Resource)

	if !blockElementCasted {
		t.Fatalf("Error in dataSourceProjects.Schema: Element for block \"%s\" is not a pointer to schema.Resource", DataSourceProjectsFilterKey)
	}

	if blockElement.Schema[DataSourceProjectsNameKey] == nil {
		t.Fatalf("Error in dataSourceProjects.Schema.subscriber: Missing argument \"%s\"", DataSourceProjectsNameKey)
	}

	if blockElement.Schema[DataSourceProjectsNameKey].Optional != true {
		t.Fatalf("Error in dataSourceProjects.Schema.subscriber: Argument \"%s\" is not optional", DataSourceProjectsNameKey)
	}
}
