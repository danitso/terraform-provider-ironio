## 0.2.1

BUG FIXES:

* provider: Fixed compatibility issues with Terraform 0.12+

## 0.2.0

BREAKING CHANGES:

* data-source/ironio_projects: Renamed `filter_name` argument to `name` and moved it into the new `filter` block
* data-source/ironio_queues: Renamed `filter_name` argument to `name` and moved it into the new `filter` block
* provider: Changed configuration schema to use blocks (`auth`, `cache`, `mq` and `worker`)

FEATURES:

* **New Data Source:** `ironio_pull_queue`
* **New Data Source:** `ironio_push_queue`

ENHANCEMENTS:

* data-source/ironio_queues: Added `filter` block and two new filters (`pull` and `push`)
* data-source/ironio_queues: Added `types` attribute
* makefile: Added `example` target to simplify the testing process
* resource/ironio_pull_queue: Added `message_count` and `message_count_total` attributes
* resource/ironio_push_queue: Added `message_count` and `message_count_total` attributes

## 0.1.0

FEATURES:

* **New Data Source:** `ironio_projects`
* **New Data Source:** `ironio_queues`
* **New Resource:** `ironio_project`
* **New Resource:** `ironio_pull_queue`
* **New Resource:** `ironio_push_queue`
