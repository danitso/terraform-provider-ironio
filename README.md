Terraform Provider for Iron.io
==============================

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.10+
- [Go](https://golang.org/doc/install) 1.12 (to build the provider plugin)

Building the Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/danitso/terraform-provider-ironio`

```sh
$ mkdir -p $GOPATH/src/github.com/danitso; cd $GOPATH/src/github.com/danitso
$ git clone git@github.com:danitso/terraform-provider-ironio
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/danitso/terraform-provider-ironio
$ make build
```

Using the Provider
----------------------
If you're building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) After placing it into your plugins directory,  run `terraform init` to initialize it.

### Configuration

The following arguments are supported:

* `host` - (Optional) This is the address of the IronMQ cluster. Defaults to `mq-aws-us-east-1-1.iron.io`.
* `load_config_file` - (Optional) This determines whether to ignore all the attributes and load the JSON configuration file instead. Defaults to `false`.
* `port` - (Optional) This is the port number for the IronMQ cluster. Defaults to `443`.
* `project_id` - (Required) This is the IronMQ project id.
* `protocol` - (Optional) This is the protocol to use for cluster communication. Defaults to `https`.
* `token` - (Required) This is the IronMQ token.

### Resources

#### ironio_pull_queue

The following arguments are supported:

* `name` - (Required) This is the name of the queue.

Example:

```
resource "ironio_pull_queue" "example" {
    name = "example_pull"
}
```

#### ironio_push_queue

The following arguments are supported:

* `name` - (Required) This is the name of the queue.
* `error_queue` - (Optional) This is the name of an error queue.
* `multicast` - (Optional) Whether to create a multicast queue instead of a unicast queue. Defaults to `true`.
* `retries` - (Optional) This is the number of times to try to send a message to a subscriber before moving the message to the error queue. Defaults to `3`.
* `retries_delay` - (Optional) This is the number of seconds to wait before retrying a failed message. Defaults to `60`.
* `subscriber` - (Required) This the subscriber block (at least one must be specified).
    * `name` - (Optional) This is the name of the subscriber. Defaults to an empty string.
    * `url` - (Required) This is the URL for the subscriber.
    * `headers` - (Optional) This is the headers to include when sending a message to the subscriber. Defaults to `{}`.

Example:

```
resource "ironio_push_queue" "example" {
    name = "example_push"

    error_queue = "example_push_error"
    multicast = true
    retries = 3
    retries_delay = 60

    subscriber {
        name = "example_push_subscriber_1"
        url = "https://subscriber1.domain.tld"

        headers {
            "X-Project-Name" = "example"
        }
    }

    subscriber {
        name = "example_push_subscriber_2"
        url = "https://subscriber2.domain.tld"

        headers {
            "X-Project-Name" = "example"
        }
    }
}
```

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.12+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-ironio
...
```

Testing the Provider
---------------------------

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```
