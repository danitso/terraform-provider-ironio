---
layout: page
title: ironio_queues
permalink: /data-sources/queues
nav_order: 4
parent: Data Sources
---

# Data Source: ironio_queues

Retrieves information about queues.

## Example Usage

```
data "ironio_queues" "example_pull" {
  project_id = ironio_project.example.id

  filter {
    name = "${ironio_project.example.name}*"
    pull = true
    push = false
  }
}
```

## Argument Reference

* `filter` - (Optional) This is the filter block.
    * `name` - (Optional) This is the name filter. You can either do an exact match, a prefix match (`prefix*`), a suffix match (`*suffix`) or a wildcard match (`*wildcard*`).
    * `pull` - (Optional) Whether to include pull queues in the result.
    * `push` - (Optional) Whether to include push queues in the result.
* `project_id` - (Required) This is the id of the project to retrieve the queues from.

## Attribute Reference

* `names` - This is the list of queue names.
* `types` - This is the list of queue types (`pull` or `push`).
