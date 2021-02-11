---
layout: page
title: ironio_push_queue
permalink: /data-sources/push_queue
nav_order: 3
parent: Data Sources
---

# Data Source: ironio_push_queue

Retrieves information about a push queue.

## Example Usage

```
data "ironio_push_queue" "example" {
  name       = ironio_push_queue.example.name
  project_id = ironio_project.example.id
}
```

## Argument Reference

* `name` - (Required) This is the name of the queue.
* `project_id` - (Required) This is the id of the project to add the queue to.

## Attribute Reference

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
