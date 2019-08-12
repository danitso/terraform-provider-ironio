# Terraform Provider for Iron.io
A Terraform Provider to manage IronAuth, IronCache, IronMQ and IronWorker resources from [Iron.io](https://iron.io/).

*Support is currently limited to IronAuth and IronMQ.*

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.11+
- [Go](https://golang.org/doc/install) 1.12 (to build the provider plugin)

## Building the Provider
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

## Using the Provider
If you're building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-plugins) After placing it into your plugins directory,  run `terraform init` to initialize it.

### Configuration

#### Arguments

* `auth` - (Optional) This is the IronAuth configuration block
    * `host` - (Optional) This is the address of the IronAuth service. Defaults to `auth.iron.io`.
    * `port` - (Optional) This is the port number for the IronAuth service. Defaults to `443`.
    * `protocol` - (Optional) This is the protocol to use for IronAuth requests. Defaults to `https`.
    * `token` - (Optional) This is the IronAuth token (OAuth). This is only optional as long as `load_config_file` is set to `true` and when the `iron.json` configuration file contains a valid token.
* `cache` - (Optional) This is the IronCache configuration block
    * `host` - (Optional) This is the address of the IronCache service. Defaults to `cache-aws-us-east-1.iron.io`.
    * `port` - (Optional) This is the port number for the IronCache service. Defaults to `443`.
    * `protocol` - (Optional) This is the protocol to use for IronCache requests. Defaults to `https`.
* `load_config_file` - (Optional) This determines whether to load configuration values from the `iron.json` configuration file. Defaults to `false`.
* `mq` - (Optional) This is the IronMQ configuration block
    * `host` - (Optional) This is the address of the IronMQ service. Defaults to `mq-aws-us-east-1-1.iron.io`.
    * `port` - (Optional) This is the port number for the IronMQ service. Defaults to `443`.
    * `protocol` - (Optional) This is the protocol to use for IronMQ requests. Defaults to `https`.
* `worker` - (Optional) This is the IronWorker configuration block
    * `host` - (Optional) This is the address of the IronWorker service. Defaults to `worker-aws-us-east-1.iron.io`.
    * `port` - (Optional) This is the port number for the IronWorker service. Defaults to `443`.
    * `protocol` - (Optional) This is the protocol to use for IronWorker requests. Defaults to `https`.

### Data Sources

#### Pull Queue (ironio_pull_queue)

##### Arguments

* `name` - (Required) This is the name of the queue.
* `project_id` - (Required) This is the id of the project to add the queue to.

##### Attributes

* `message_count` - This is the number of messages currently in the queue.
* `message_count_total` - This is the number of messages which have been processed by the queue.

#### Push Queue (ironio_push_queue)

##### Arguments

* `name` - (Required) This is the name of the queue.
* `project_id` - (Required) This is the id of the project to add the queue to.

##### Attributes

* `error_queue` - This is the name of an error queue.
* `message_count` - This is the number of messages currently in the queue.
* `message_count_total` - This is the number of messages which have been processed by the queue.
* `multicast` - Whether to create a multicast queue instead of a unicast queue.
* `retries` - This is the number of times to try to send a message to a subscriber before moving the message to the error queue.
* `retries_delay` - This is the number of seconds to wait before retrying a failed message.
* `subscriber` - This is the list of subscribers.
    * `headers` - This is the headers to include when sending a message to the subscriber.
    * `name` - This is the name of the subscriber.
    * `url` - This is the URL for the subscriber.

#### Projects (ironio_projects)

##### Arguments

* `filter` - (Optional) This is the filter block.
    * `name` - (Optional) This is the name filter. You can either do an exact match, a prefix match (`prefix*`), a suffix match (`*suffix`) or a wildcard match (`*wildcard*`).

##### Attributes

* `ids` - This is the list of project ids.
* `names` - This is the list of project names.

#### Queues (ironio_queues)

##### Arguments

* `filter` - (Optional) This is the filter block.
    * `name` - (Optional) This is the name filter. You can either do an exact match, a prefix match (`prefix*`), a suffix match (`*suffix`) or a wildcard match (`*wildcard*`).
    * `pull` - (Optional) Whether to include pull queues in the result.
    * `push` - (Optional) Whether to include push queues in the result.
* `project_id` - (Required) This is the id of the project to retrieve the queues from.

##### Attributes

* `names` - This is the list of queue names.
* `types` - This is the list of queue types (`pull` or `push`).

### Resources

#### Project (ironio_project)

##### Arguments

* `name` - (Required) This is the name of the project.

#### Pull Queue (ironio_pull_queue)

##### Arguments

* `name` - (Required) This is the name of the queue.
* `project_id` - (Required) This is the id of the project to add the queue to.

##### Attributes

* `message_count` - This is the number of messages currently in the queue.
* `message_count_total` - This is the number of messages which have been processed by the queue.

#### Push Queue (ironio_push_queue)

##### Arguments

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

##### Attributes

* `message_count` - This is the number of messages currently in the queue.
* `message_count_total` - This is the number of messages which have been processed by the queue.

## Developing the Provider
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

## Testing the Provider
In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

Tests are limited to regression tests, ensuring backwards compability.
