package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

// TestProviderInstantiation() tests whether the Provider instance can be instantiated.
func TestProviderInstantiation(t *testing.T) {
	s := Provider()

	if s == nil {
		t.Fatalf("Cannot instantiate Provider")
	}
}

// TestProviderConfiguration() tests the Provider schema.
func TestProviderSchema(t *testing.T) {
	s := Provider()

	if s.Schema[ProviderConfigurationLoadConfigFileKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\"", ProviderConfigurationLoadConfigFileKey)
	}
}

// TestProviderSchemaAuth() tests the IronAuth schema.
func TestProviderSchemaAuth(t *testing.T) {
	s := Provider()

	if s.Schema[ProviderConfigurationAuthKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing block \"%s\"", ProviderConfigurationAuthKey)
	}

	if s.Schema[ProviderConfigurationAuthKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not optional", ProviderConfigurationAuthKey)
	}

	if s.Schema[ProviderConfigurationAuthKey].Type != schema.TypeList {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not a list", ProviderConfigurationAuthKey)
	}

	if s.Schema[ProviderConfigurationAuthKey].MaxItems != 1 {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not limited to a single definition", ProviderConfigurationAuthKey)
	}

	if s.Schema[ProviderConfigurationAuthKey].Elem == nil {
		t.Fatalf("Error in Provider.Schema: Missing element for block \"%s\"", ProviderConfigurationAuthKey)
	}

	blockElement, blockElementCasted := s.Schema[ProviderConfigurationAuthKey].Elem.(*schema.Resource)

	if !blockElementCasted {
		t.Fatalf("Error in Provider.Schema: Element for block \"%s\" is not a pointer to schema.Resource", ProviderConfigurationAuthKey)
	}

	if blockElement.Schema[ProviderConfigurationHostKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", ProviderConfigurationHostKey, ProviderConfigurationAuthKey)
	}

	if blockElement.Schema[ProviderConfigurationHostKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", ProviderConfigurationHostKey, ProviderConfigurationAuthKey)
	}

	if blockElement.Schema[ProviderConfigurationPortKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", ProviderConfigurationPortKey, ProviderConfigurationAuthKey)
	}

	if blockElement.Schema[ProviderConfigurationPortKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", ProviderConfigurationPortKey, ProviderConfigurationAuthKey)
	}

	if blockElement.Schema[ProviderConfigurationPortKey].Type != schema.TypeInt {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not limited to integers", ProviderConfigurationPortKey, ProviderConfigurationAuthKey)
	}

	if blockElement.Schema[ProviderConfigurationProtocolKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", ProviderConfigurationProtocolKey, ProviderConfigurationAuthKey)
	}

	if blockElement.Schema[ProviderConfigurationProtocolKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", ProviderConfigurationProtocolKey, ProviderConfigurationAuthKey)
	}

	if blockElement.Schema[ProviderConfigurationTokenKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", ProviderConfigurationTokenKey, ProviderConfigurationAuthKey)
	}

	if blockElement.Schema[ProviderConfigurationTokenKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", ProviderConfigurationTokenKey, ProviderConfigurationAuthKey)
	}
}

// TestProviderSchemaCache() tests the IronCache schema.
func TestProviderSchemaCache(t *testing.T) {
	s := Provider()

	if s.Schema[ProviderConfigurationCacheKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing block \"%s\"", ProviderConfigurationCacheKey)
	}

	if s.Schema[ProviderConfigurationCacheKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not optional", ProviderConfigurationCacheKey)
	}

	if s.Schema[ProviderConfigurationCacheKey].Type != schema.TypeList {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not a list", ProviderConfigurationCacheKey)
	}

	if s.Schema[ProviderConfigurationCacheKey].MaxItems != 1 {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not limited to a single definition", ProviderConfigurationCacheKey)
	}

	if s.Schema[ProviderConfigurationCacheKey].Elem == nil {
		t.Fatalf("Error in Provider.Schema: Missing element for block \"%s\"", ProviderConfigurationCacheKey)
	}

	blockElement, blockElementCasted := s.Schema[ProviderConfigurationCacheKey].Elem.(*schema.Resource)

	if !blockElementCasted {
		t.Fatalf("Error in Provider.Schema: Element for block \"%s\" is not a pointer to schema.Resource", ProviderConfigurationCacheKey)
	}

	if blockElement.Schema[ProviderConfigurationHostKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", ProviderConfigurationHostKey, ProviderConfigurationCacheKey)
	}

	if blockElement.Schema[ProviderConfigurationHostKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", ProviderConfigurationHostKey, ProviderConfigurationCacheKey)
	}

	if blockElement.Schema[ProviderConfigurationPortKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", ProviderConfigurationPortKey, ProviderConfigurationCacheKey)
	}

	if blockElement.Schema[ProviderConfigurationPortKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", ProviderConfigurationPortKey, ProviderConfigurationCacheKey)
	}

	if blockElement.Schema[ProviderConfigurationPortKey].Type != schema.TypeInt {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not limited to integers", ProviderConfigurationPortKey, ProviderConfigurationCacheKey)
	}

	if blockElement.Schema[ProviderConfigurationProtocolKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", ProviderConfigurationProtocolKey, ProviderConfigurationCacheKey)
	}

	if blockElement.Schema[ProviderConfigurationProtocolKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", ProviderConfigurationProtocolKey, ProviderConfigurationCacheKey)
	}
}

// TestProviderSchemaMQ() tests the IronMQ schema.
func TestProviderSchemaMQ(t *testing.T) {
	s := Provider()

	if s.Schema[ProviderConfigurationMQKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing block \"%s\"", ProviderConfigurationMQKey)
	}

	if s.Schema[ProviderConfigurationMQKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not optional", ProviderConfigurationMQKey)
	}

	if s.Schema[ProviderConfigurationMQKey].Type != schema.TypeList {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not a list", ProviderConfigurationMQKey)
	}

	if s.Schema[ProviderConfigurationMQKey].MaxItems != 1 {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not limited to a single definition", ProviderConfigurationMQKey)
	}

	if s.Schema[ProviderConfigurationMQKey].Elem == nil {
		t.Fatalf("Error in Provider.Schema: Missing element for block \"%s\"", ProviderConfigurationMQKey)
	}

	blockElement, blockElementCasted := s.Schema[ProviderConfigurationMQKey].Elem.(*schema.Resource)

	if !blockElementCasted {
		t.Fatalf("Error in Provider.Schema: Element for block \"%s\" is not a pointer to schema.Resource", ProviderConfigurationMQKey)
	}

	if blockElement.Schema[ProviderConfigurationHostKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", ProviderConfigurationHostKey, ProviderConfigurationMQKey)
	}

	if blockElement.Schema[ProviderConfigurationHostKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", ProviderConfigurationHostKey, ProviderConfigurationMQKey)
	}

	if blockElement.Schema[ProviderConfigurationPortKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", ProviderConfigurationPortKey, ProviderConfigurationMQKey)
	}

	if blockElement.Schema[ProviderConfigurationPortKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", ProviderConfigurationPortKey, ProviderConfigurationMQKey)
	}

	if blockElement.Schema[ProviderConfigurationPortKey].Type != schema.TypeInt {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not limited to integers", ProviderConfigurationPortKey, ProviderConfigurationMQKey)
	}

	if blockElement.Schema[ProviderConfigurationProtocolKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", ProviderConfigurationProtocolKey, ProviderConfigurationMQKey)
	}

	if blockElement.Schema[ProviderConfigurationProtocolKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", ProviderConfigurationProtocolKey, ProviderConfigurationMQKey)
	}
}

// TestProviderSchemaWorker() tests the IronWorker schema.
func TestProviderSchemaWorker(t *testing.T) {
	s := Provider()

	if s.Schema[ProviderConfigurationWorkerKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing block \"%s\"", ProviderConfigurationWorkerKey)
	}

	if s.Schema[ProviderConfigurationWorkerKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not optional", ProviderConfigurationWorkerKey)
	}

	if s.Schema[ProviderConfigurationWorkerKey].Type != schema.TypeList {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not a list", ProviderConfigurationWorkerKey)
	}

	if s.Schema[ProviderConfigurationWorkerKey].MaxItems != 1 {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not limited to a single definition", ProviderConfigurationWorkerKey)
	}

	if s.Schema[ProviderConfigurationWorkerKey].Elem == nil {
		t.Fatalf("Error in Provider.Schema: Missing element for block \"%s\"", ProviderConfigurationWorkerKey)
	}

	blockElement, blockElementCasted := s.Schema[ProviderConfigurationWorkerKey].Elem.(*schema.Resource)

	if !blockElementCasted {
		t.Fatalf("Error in Provider.Schema: Element for block \"%s\" is not a pointer to schema.Resource", ProviderConfigurationWorkerKey)
	}

	if blockElement.Schema[ProviderConfigurationHostKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", ProviderConfigurationHostKey, ProviderConfigurationWorkerKey)
	}

	if blockElement.Schema[ProviderConfigurationHostKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", ProviderConfigurationHostKey, ProviderConfigurationWorkerKey)
	}

	if blockElement.Schema[ProviderConfigurationPortKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", ProviderConfigurationPortKey, ProviderConfigurationWorkerKey)
	}

	if blockElement.Schema[ProviderConfigurationPortKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", ProviderConfigurationPortKey, ProviderConfigurationWorkerKey)
	}

	if blockElement.Schema[ProviderConfigurationPortKey].Type != schema.TypeInt {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not limited to integers", ProviderConfigurationPortKey, ProviderConfigurationWorkerKey)
	}

	if blockElement.Schema[ProviderConfigurationProtocolKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", ProviderConfigurationProtocolKey, ProviderConfigurationWorkerKey)
	}

	if blockElement.Schema[ProviderConfigurationProtocolKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", ProviderConfigurationProtocolKey, ProviderConfigurationWorkerKey)
	}
}
