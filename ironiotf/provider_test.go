/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package ironiotf

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// TestProviderInstantiation tests whether the Provider instance can be instantiated.
func TestProviderInstantiation(t *testing.T) {
	s := Provider()

	if s == nil {
		t.Fatalf("Cannot instantiate Provider")
	}
}

// TestProviderConfiguration tests the Provider schema.
func TestProviderSchema(t *testing.T) {
	s := Provider()

	if s.Schema[providerConfigurationLoadConfigFileKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\"", providerConfigurationLoadConfigFileKey)
	}
}

// TestProviderSchemaAuth tests the IronAuth schema.
func TestProviderSchemaAuth(t *testing.T) {
	s := Provider()

	if s.Schema[providerConfigurationAuthKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing block \"%s\"", providerConfigurationAuthKey)
	}

	if s.Schema[providerConfigurationAuthKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not optional", providerConfigurationAuthKey)
	}

	if s.Schema[providerConfigurationAuthKey].Type != schema.TypeList {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not a list", providerConfigurationAuthKey)
	}

	if s.Schema[providerConfigurationAuthKey].MaxItems != 1 {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not limited to a single definition", providerConfigurationAuthKey)
	}

	if s.Schema[providerConfigurationAuthKey].Elem == nil {
		t.Fatalf("Error in Provider.Schema: Missing element for block \"%s\"", providerConfigurationAuthKey)
	}

	blockElement, blockElementCasted := s.Schema[providerConfigurationAuthKey].Elem.(*schema.Resource)

	if !blockElementCasted {
		t.Fatalf("Error in Provider.Schema: Element for block \"%s\" is not a pointer to schema.Resource", providerConfigurationAuthKey)
	}

	if blockElement.Schema[providerConfigurationHostKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", providerConfigurationHostKey, providerConfigurationAuthKey)
	}

	if blockElement.Schema[providerConfigurationHostKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", providerConfigurationHostKey, providerConfigurationAuthKey)
	}

	if blockElement.Schema[providerConfigurationPortKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", providerConfigurationPortKey, providerConfigurationAuthKey)
	}

	if blockElement.Schema[providerConfigurationPortKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", providerConfigurationPortKey, providerConfigurationAuthKey)
	}

	if blockElement.Schema[providerConfigurationPortKey].Type != schema.TypeInt {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not limited to integers", providerConfigurationPortKey, providerConfigurationAuthKey)
	}

	if blockElement.Schema[providerConfigurationProtocolKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", providerConfigurationProtocolKey, providerConfigurationAuthKey)
	}

	if blockElement.Schema[providerConfigurationProtocolKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", providerConfigurationProtocolKey, providerConfigurationAuthKey)
	}

	if blockElement.Schema[providerConfigurationTokenKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", providerConfigurationTokenKey, providerConfigurationAuthKey)
	}

	if blockElement.Schema[providerConfigurationTokenKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", providerConfigurationTokenKey, providerConfigurationAuthKey)
	}
}

// TestProviderSchemaCache tests the IronCache schema.
func TestProviderSchemaCache(t *testing.T) {
	s := Provider()

	if s.Schema[providerConfigurationCacheKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing block \"%s\"", providerConfigurationCacheKey)
	}

	if s.Schema[providerConfigurationCacheKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not optional", providerConfigurationCacheKey)
	}

	if s.Schema[providerConfigurationCacheKey].Type != schema.TypeList {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not a list", providerConfigurationCacheKey)
	}

	if s.Schema[providerConfigurationCacheKey].MaxItems != 1 {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not limited to a single definition", providerConfigurationCacheKey)
	}

	if s.Schema[providerConfigurationCacheKey].Elem == nil {
		t.Fatalf("Error in Provider.Schema: Missing element for block \"%s\"", providerConfigurationCacheKey)
	}

	blockElement, blockElementCasted := s.Schema[providerConfigurationCacheKey].Elem.(*schema.Resource)

	if !blockElementCasted {
		t.Fatalf("Error in Provider.Schema: Element for block \"%s\" is not a pointer to schema.Resource", providerConfigurationCacheKey)
	}

	if blockElement.Schema[providerConfigurationHostKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", providerConfigurationHostKey, providerConfigurationCacheKey)
	}

	if blockElement.Schema[providerConfigurationHostKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", providerConfigurationHostKey, providerConfigurationCacheKey)
	}

	if blockElement.Schema[providerConfigurationPortKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", providerConfigurationPortKey, providerConfigurationCacheKey)
	}

	if blockElement.Schema[providerConfigurationPortKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", providerConfigurationPortKey, providerConfigurationCacheKey)
	}

	if blockElement.Schema[providerConfigurationPortKey].Type != schema.TypeInt {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not limited to integers", providerConfigurationPortKey, providerConfigurationCacheKey)
	}

	if blockElement.Schema[providerConfigurationProtocolKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", providerConfigurationProtocolKey, providerConfigurationCacheKey)
	}

	if blockElement.Schema[providerConfigurationProtocolKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", providerConfigurationProtocolKey, providerConfigurationCacheKey)
	}
}

// TestProviderSchemaMQ tests the IronMQ schema.
func TestProviderSchemaMQ(t *testing.T) {
	s := Provider()

	if s.Schema[providerConfigurationMQKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing block \"%s\"", providerConfigurationMQKey)
	}

	if s.Schema[providerConfigurationMQKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not optional", providerConfigurationMQKey)
	}

	if s.Schema[providerConfigurationMQKey].Type != schema.TypeList {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not a list", providerConfigurationMQKey)
	}

	if s.Schema[providerConfigurationMQKey].MaxItems != 1 {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not limited to a single definition", providerConfigurationMQKey)
	}

	if s.Schema[providerConfigurationMQKey].Elem == nil {
		t.Fatalf("Error in Provider.Schema: Missing element for block \"%s\"", providerConfigurationMQKey)
	}

	blockElement, blockElementCasted := s.Schema[providerConfigurationMQKey].Elem.(*schema.Resource)

	if !blockElementCasted {
		t.Fatalf("Error in Provider.Schema: Element for block \"%s\" is not a pointer to schema.Resource", providerConfigurationMQKey)
	}

	if blockElement.Schema[providerConfigurationHostKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", providerConfigurationHostKey, providerConfigurationMQKey)
	}

	if blockElement.Schema[providerConfigurationHostKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", providerConfigurationHostKey, providerConfigurationMQKey)
	}

	if blockElement.Schema[providerConfigurationPortKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", providerConfigurationPortKey, providerConfigurationMQKey)
	}

	if blockElement.Schema[providerConfigurationPortKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", providerConfigurationPortKey, providerConfigurationMQKey)
	}

	if blockElement.Schema[providerConfigurationPortKey].Type != schema.TypeInt {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not limited to integers", providerConfigurationPortKey, providerConfigurationMQKey)
	}

	if blockElement.Schema[providerConfigurationProtocolKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", providerConfigurationProtocolKey, providerConfigurationMQKey)
	}

	if blockElement.Schema[providerConfigurationProtocolKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", providerConfigurationProtocolKey, providerConfigurationMQKey)
	}
}

// TestProviderSchemaWorker tests the IronWorker schema.
func TestProviderSchemaWorker(t *testing.T) {
	s := Provider()

	if s.Schema[providerConfigurationWorkerKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing block \"%s\"", providerConfigurationWorkerKey)
	}

	if s.Schema[providerConfigurationWorkerKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not optional", providerConfigurationWorkerKey)
	}

	if s.Schema[providerConfigurationWorkerKey].Type != schema.TypeList {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not a list", providerConfigurationWorkerKey)
	}

	if s.Schema[providerConfigurationWorkerKey].MaxItems != 1 {
		t.Fatalf("Error in Provider.Schema: Block \"%s\" is not limited to a single definition", providerConfigurationWorkerKey)
	}

	if s.Schema[providerConfigurationWorkerKey].Elem == nil {
		t.Fatalf("Error in Provider.Schema: Missing element for block \"%s\"", providerConfigurationWorkerKey)
	}

	blockElement, blockElementCasted := s.Schema[providerConfigurationWorkerKey].Elem.(*schema.Resource)

	if !blockElementCasted {
		t.Fatalf("Error in Provider.Schema: Element for block \"%s\" is not a pointer to schema.Resource", providerConfigurationWorkerKey)
	}

	if blockElement.Schema[providerConfigurationHostKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", providerConfigurationHostKey, providerConfigurationWorkerKey)
	}

	if blockElement.Schema[providerConfigurationHostKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", providerConfigurationHostKey, providerConfigurationWorkerKey)
	}

	if blockElement.Schema[providerConfigurationPortKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", providerConfigurationPortKey, providerConfigurationWorkerKey)
	}

	if blockElement.Schema[providerConfigurationPortKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", providerConfigurationPortKey, providerConfigurationWorkerKey)
	}

	if blockElement.Schema[providerConfigurationPortKey].Type != schema.TypeInt {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not limited to integers", providerConfigurationPortKey, providerConfigurationWorkerKey)
	}

	if blockElement.Schema[providerConfigurationProtocolKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\" for block \"%s\"", providerConfigurationProtocolKey, providerConfigurationWorkerKey)
	}

	if blockElement.Schema[providerConfigurationProtocolKey].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" for block \"%s\" is not optional", providerConfigurationProtocolKey, providerConfigurationWorkerKey)
	}
}
