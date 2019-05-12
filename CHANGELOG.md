## 0.2.0 (Unreleased)

BREAKING CHANGES:

* provider: Changed configuration schema to use blocks (`auth`, `cache`, `mq` and `worker`)

FEATURES:

* **New Data Source:** `ironio_pull_queue`
* **New Data Source:** `ironio_push_queue`

ENHANCEMENTS:

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
