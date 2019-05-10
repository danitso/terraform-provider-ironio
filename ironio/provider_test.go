package ironio

import "testing"

func TestProviderDefaultValues(t *testing.T) {
	s := Provider()

	if s == nil {
		t.Fatalf("cannot instantiate Provider")
	}

	if s.Schema["host"].Default != "" {
		t.Fatalf("error in Provider.Schema, host Default value has changed")
	}

	if s.Schema["load_config_file"].Default != false {
		t.Fatalf("error in Provider.Schema, load_config_file Default value has changed")
	}

	if s.Schema["port"].Default != 0 {
		t.Fatalf("error in Provider.Schema, port Default value has changed")
	}

	if s.Schema["protocol"].Default != "" {
		t.Fatalf("error in Provider.Schema, protocol Default value has changed")
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

	if s.Schema["host"].Optional != true {
		t.Fatalf("error in Provider.Schema, host Optional value has changed")
	}

	if s.Schema["load_config_file"].Optional != true {
		t.Fatalf("error in Provider.Schema, load_config_file Optional value has changed")
	}

	if s.Schema["port"].Optional != true {
		t.Fatalf("error in Provider.Schema, port Optional value has changed")
	}

	if s.Schema["project_id"].Required != true {
		t.Fatalf("error in Provider.Schema, project_id Optional value has changed")
	}

	if s.Schema["protocol"].Optional != true {
		t.Fatalf("error in Provider.Schema, protocol Optional value has changed")
	}

	if s.Schema["token"].Required != true {
		t.Fatalf("error in Provider.Schema, token Required value has changed")
	}
}
