package main

import "testing"

func TestProviderDefaultValues(t *testing.T) {
	s := Provider()

	if s == nil {
		t.Fatalf("cannot instantiate Provider")
	}

	if s.Schema["auth_host"].Default != "" {
		t.Fatalf("error in Provider.Schema, auth_host Default value has changed")
	}

	if s.Schema["auth_port"].Default != 0 {
		t.Fatalf("error in Provider.Schema, auth_port Default value has changed")
	}

	if s.Schema["auth_protocol"].Default != "" {
		t.Fatalf("error in Provider.Schema, auth_protocol Default value has changed")
	}

	if s.Schema["cache_host"].Default != "" {
		t.Fatalf("error in Provider.Schema, cache_host Default value has changed")
	}

	if s.Schema["cache_port"].Default != 0 {
		t.Fatalf("error in Provider.Schema, cache_port Default value has changed")
	}

	if s.Schema["cache_protocol"].Default != "" {
		t.Fatalf("error in Provider.Schema, cache_protocol Default value has changed")
	}

	if s.Schema["load_config_file"].Default != false {
		t.Fatalf("error in Provider.Schema, load_config_file Default value has changed")
	}

	if s.Schema["mq_host"].Default != "" {
		t.Fatalf("error in Provider.Schema, mq_host Default value has changed")
	}

	if s.Schema["mq_port"].Default != 0 {
		t.Fatalf("error in Provider.Schema, mq_port Default value has changed")
	}

	if s.Schema["mq_protocol"].Default != "" {
		t.Fatalf("error in Provider.Schema, mq_protocol Default value has changed")
	}

	if s.Schema["worker_host"].Default != "" {
		t.Fatalf("error in Provider.Schema, worker_host Default value has changed")
	}

	if s.Schema["worker_port"].Default != 0 {
		t.Fatalf("error in Provider.Schema, worker_port Default value has changed")
	}

	if s.Schema["worker_protocol"].Default != "" {
		t.Fatalf("error in Provider.Schema, worker_protocol Default value has changed")
	}
}

func TestProviderRequiredOptionalValues(t *testing.T) {
	s := Provider()

	if s == nil {
		t.Fatalf("cannot instantiate Provider")
	}

	if len(s.Schema) == 0 {
		t.Fatalf("error in Provider.Schema, list empty")
	}

	if s.Schema["auth_host"].Optional != true {
		t.Fatalf("error in Provider.Schema, auth_host Optional value has changed")
	}

	if s.Schema["auth_port"].Optional != true {
		t.Fatalf("error in Provider.Schema, auth_port Optional value has changed")
	}

	if s.Schema["auth_protocol"].Optional != true {
		t.Fatalf("error in Provider.Schema, auth_protocol Optional value has changed")
	}

	if s.Schema["cache_host"].Optional != true {
		t.Fatalf("error in Provider.Schema, cache_host Optional value has changed")
	}

	if s.Schema["cache_port"].Optional != true {
		t.Fatalf("error in Provider.Schema, cache_port Optional value has changed")
	}

	if s.Schema["cache_protocol"].Optional != true {
		t.Fatalf("error in Provider.Schema, cache_protocol Optional value has changed")
	}

	if s.Schema["load_config_file"].Optional != true {
		t.Fatalf("error in Provider.Schema, load_config_file Optional value has changed")
	}

	if s.Schema["mq_host"].Optional != true {
		t.Fatalf("error in Provider.Schema, mq_host Optional value has changed")
	}

	if s.Schema["mq_port"].Optional != true {
		t.Fatalf("error in Provider.Schema, mq_port Optional value has changed")
	}

	if s.Schema["mq_protocol"].Optional != true {
		t.Fatalf("error in Provider.Schema, mq_protocol Optional value has changed")
	}

	if s.Schema["token"].Optional != true {
		t.Fatalf("error in Provider.Schema, token Required value has changed")
	}

	if s.Schema["worker_host"].Optional != true {
		t.Fatalf("error in Provider.Schema, worker_host Optional value has changed")
	}

	if s.Schema["worker_port"].Optional != true {
		t.Fatalf("error in Provider.Schema, worker_port Optional value has changed")
	}

	if s.Schema["worker_protocol"].Optional != true {
		t.Fatalf("error in Provider.Schema, worker_protocol Optional value has changed")
	}
}
