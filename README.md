[![Build Status](https://api.travis-ci.com/danitso/terraform-provider-ironio.svg?branch=master)](https://travis-ci.com/danitso/terraform-provider-ironio)
[![Go Report Card](https://goreportcard.com/badge/github.com/danitso/terraform-provider-ironio)](https://goreportcard.com/report/github.com/danitso/terraform-provider-ironio)
[![GoDoc](https://godoc.org/github.com/danitso/terraform-provider-ironio?status.svg)](http://godoc.org/github.com/danitso/terraform-provider-ironio)

# Terraform Provider for Iron.io
A Terraform Provider to manage IronAuth, IronCache, IronMQ and IronWorker resources from [Iron.io](https://iron.io/).

*Support is currently limited to IronAuth and IronMQ.*

## Requirements
- [Terraform](https://www.terraform.io/downloads.html) 0.13+
- [Go](https://golang.org/doc/install) 1.15+ (to build the provider plugin)
- [GoReleaser](https://goreleaser.com/install/) 0.155+ (to build the provider plugin)

## Table of Contents
- [Building the provider](#building-the-provider)
- [Using the provider](#using-the-provider)
- [Testing the provider](#testing-the-provider)

## Building the provider
- Clone the repository to `$GOPATH/src/github.com/danitso/terraform-provider-ironio`:

    ```sh
    $ mkdir -p "${GOPATH}/src/github.com/danitso"
    $ cd "${GOPATH}/src/github.com/danitso"
    $ git clone git@github.com:danitso/terraform-provider-ironio
    ```

- Enter the provider directory and build it:

    ```sh
    $ cd "${GOPATH}/src/github.com/danitso/terraform-provider-ironio"
    $ make build
    ```

## Using the provider
You can find the latest release and its documentation in the [Terraform Registry](https://registry.terraform.io/providers/danitso/ironio/latest).

## Testing the provider
In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

Tests are limited to regression tests, ensuring backwards compability.
