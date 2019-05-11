GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
NAME=$$(grep TerraformProviderName version.go | grep -o -P 'terraform-provider-[a-z]+')
TARGETS=darwin linux windows
VERSION=$$(grep TerraformProviderVersion version.go | grep -o -P '\d\.\d\.\d')

default: build

build:
	go build \
		-o "bin/$(NAME)_v$(VERSION)_x4"

fmt:
	gofmt -w $(GOFMT_FILES)

test:
	go test -v

init:
	go get ./...

targets: $(TARGETS)

$(TARGETS):
	GOOS=$@ GOARCH=amd64 CGO_ENABLED=0 go build \
		-o "dist/$@/$(NAME)_v$(VERSION)_x4" \
		-a -ldflags '-extldflags "-static"'

.PHONY: build fmt test init targets $(TARGETS)
