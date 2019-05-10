default: build

build: install
	go build .

test:
	go test -v

install:
	go get -u github.com/hashicorp/terraform/plugin
	go get -u github.com/hashicorp/terraform/terraform


.PHONY: build test install

