---
layout: page
title: ironio_pull_queue
permalink: /resources/pull_queue
nav_order: 2
parent: Resources
---

# Resource: ironio_pull_queue

Manages a pull queue.

## Example Usage

```
resource "ironio_pull_queue" "example" {
  name       = "example-pull-1"
  project_id = ironio_project.example.id
}
```

## Argument Reference

* `name` - (Required) This is the name of the queue.
* `project_id` - (Required) This is the id of the project to add the queue to.

## Attribute Reference

* `message_count` - This is the number of messages currently in the queue.
* `message_count_total` - This is the number of messages which have been processed by the queue.
