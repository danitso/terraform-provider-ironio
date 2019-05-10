default: build

build:
	go build .

test:
	go test -v

init:
	go get -u github.com/hashicorp/terraform/plugin
	go get -u github.com/hashicorp/terraform/terraform


.PHONY: build test init

