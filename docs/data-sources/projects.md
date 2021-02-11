---
layout: page
title: ironio_projects
permalink: /data-sources/projects
nav_order: 1
parent: Data Sources
---

# Data Source: ironio_projects

Retrieves information about projects.

## Example Usage

```
data "ironio_projects" "example" {
  depends_on = [ironio_project.example]

  filter {
    name = ironio_project.example.name
  }
}
```

## Argument Reference

* `filter` - (Optional) This is the filter block.
    * `name` - (Optional) This is the name filter. You can either do an exact match, a prefix match (`prefix*`), a suffix match (`*suffix`) or a wildcard match (`*wildcard*`).

## Attribute Reference

* `ids` - This is the list of project ids.
* `names` - This is the list of project names.
