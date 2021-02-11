---
layout: page
title: ironio_push_queue
permalink: /resources/push_queue
nav_order: 3
parent: Resources
---

# Resource: ironio_push_queue

Manages a push queue.

## Example Usage

```
resource "ironio_push_queue" "example" {
  name = "example-push"

  error_queue   = ironio_pull_queue.example_push_error.name
  retries       = 4
  retries_delay = 30
  multicast     = true
  project_id    = ironio_project.example.id

  subscriber {
    name = "example-push-subscriber-1"
    url  = "ironmq:///${ironio_pull_queue.example.name}"

    headers = {
      X-Push-Queue = ironio_pull_queue.example.name
      X-Push-Token = "example-push-token-1"
    }
  }
}
```

## Argument Reference

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

## Attribute Reference

* `message_count` - This is the number of messages currently in the queue.
* `message_count_total` - This is the number of messages which have been processed by the queue.
