# Terraform Provider for Iron.io
A Terraform Provider to manage IronAuth, IronCache, IronMQ and IronWorker resources from [Iron.io](https://iron.io/). Support is currently limited to IronAuth and IronMQ.

# Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.11+
- [Go](https://golang.org/doc/install) 1.12 (to build the provider plugin)

# Building the Provider
Clone repository to: `$GOPATH/src/github.com/danitso/terraform-provider-ironio`

```sh
$ mkdir -p $GOPATH/src/github.com/danitso; cd $GOPATH/src/github.com/danitso
$ git clone git@github.com:danitso/terraform-provider-ironio
```

Enter the provider directory, initialize and build the provider

```sh
$ cd $GOPATH/src/github.com/danitso/terraform-provider-ironio
$ make init
$ make build
```

# Using the Provider
If you're building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-plugins) After placing it into your plugins directory,  run `terraform init` to initialize it.

## Configuration

### Arguments

* `auth_host` - (Optional) This is the address of the IronAuth service. Defaults to `auth.iron.io`.
* `auth_port` - (Optional) This is the port number for the IronAuth service. Defaults to `443`.
* `auth_protocol` - (Optional) This is the protocol to use for IronAuth requests. Defaults to `https`.
* `cache_host` - (Optional) This is the address of the IronCache service. Defaults to `cache-aws-us-east-1.iron.io`.
* `cache_port` - (Optional) This is the port number for the IronCache service. Defaults to `443`.
* `cache_protocol` - (Optional) This is the protocol to use for IronCache requests. Defaults to `https`.
* `load_config_file` - (Optional) This determines whether to load values from the JSON configuration file (iron.json). Defaults to `false`.
* `mq_host` - (Optional) This is the address of the IronMQ service. Defaults to `mq-aws-us-east-1-1.iron.io`.
* `mq_port` - (Optional) This is the port number for the IronMQ service. Defaults to `443`.
* `mq_protocol` - (Optional) This is the protocol to use for IronMQ requests. Defaults to `https`.
* `token` - (Required) This is the IronAuth token (OAuth).
* `worker_host` - (Optional) This is the address of the IronWorker service. Defaults to `worker-aws-us-east-1.iron.io`.
* `worker_port` - (Optional) This is the port number for the IronWorker service. Defaults to `443`.
* `worker_protocol` - (Optional) This is the protocol to use for IronWorker requests. Defaults to `https`.

## Data Sources

### Projects (ironio_projects)

#### Arguments

* `filter_name` - (Optional) This is the name filter. You can either do an exact match, a prefix match (`prefix*`), a suffix match (`*suffix`) or a wildcard match (`*wildcard*`).

#### Attributes

* `ids` - This is the list of project ids.
* `names` - This is the list of project names.

### Queues (ironio_queues)

#### Arguments

* `filter_name` - (Optional) This is the name filter. You can either do an exact match, a prefix match (`prefix*`), a suffix match (`*suffix`) or a wildcard match (`*wildcard*`).
* `project_id` - (Required) This is the id of the project to retrieve the queues from.

#### Attributes

* `names` - This is the list of queue names.

## Resources

### Project (ironio_project)

#### Arguments

* `name` - (Required) This is the name of the project.

### Pull Queue (ironio_pull_queue)

#### Arguments

* `name` - (Required) This is the name of the queue.
* `project_id` - (Required) This is the id of the project to add the queue to.

### Push Queue (ironio_push_queue)

#### Arguments

* `error_queue` - (Optional) This is the name of an error queue.
* `multicast` - (Optional) Whether to create a multicast queue instead of a unicast queue. Defaults to `true`.
* `name` - (Required) This is the name of the queue.
* `project_id` - (Required) This is the id of the project to add the queue to.
* `retries` - (Optional) This is the number of times to try to send a message to a subscriber before moving the message to the error queue. Defaults to `3`.
* `retries_delay` - (Optional) This is the number of seconds to wait before retrying a failed message. Defaults to `60`.
* `subscriber` - (Required) This the subscriber block (at least one must be specified).
    * `headers` - (Optional) This is the headers to include when sending a message to the subscriber. Defaults to `{}`.
    * `name` - (Optional) This is the name of the subscriber. Defaults to an empty string.
    * `url` - (Required) This is the URL for the subscriber.

# Developing the Provider
If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.12+ is *required*).
You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-ironio
...
```

If you wish to contribute to the provider, the following requirements must be met,

* All tests must pass using `make test`
* The Go code must be formatted using Gofmt
* Dependencies are installed by `make init`

# Testing the Provider
In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

Tests are limited to regression tests, ensuring backwards compability.
