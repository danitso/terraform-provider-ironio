---
layout: home
title: Introduction
nav_order: 1
---

# Iron.io Provider

This provider for [Terraform](https://www.terraform.io/) is used for interacting with the resources supported by [Iron.io](https://iron.io). The provider needs to be configured with the proper endpoints and tokens before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```
provider "ironio" {
  auth {
    token = "your-api-token-here"
  }
}
```

## Argument Reference

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
