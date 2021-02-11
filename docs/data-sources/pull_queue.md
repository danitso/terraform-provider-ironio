---
layout: page
title: ironio_pull_queue
permalink: /data-sources/pull_queue
nav_order: 2
parent: Data Sources
---

# Data Source: ironio_pull_queue

Retrieves information about a pull queue.

## Example Usage

```
data "ironio_pull_queue" "example" {
  name       = ironio_pull_queue.example.name
  project_id = ironio_project.example.id
}
```

## Argument Reference

* `name` - (Required) This is the name of the queue.
* `project_id` - (Required) This is the id of the project to add the queue to.

## Attribute Reference

* `message_count` - This is the number of messages currently in the queue.
* `message_count_total` - This is the number of messages which have been processed by the queue.
