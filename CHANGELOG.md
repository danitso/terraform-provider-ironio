## 0.2.0 (Unreleased)

BREAKING CHANGES:

* provider: Changed configuration schema to use blocks (`auth`, `cache`, `mq` and `worker`)
* provider: Changed the default value of `load_config_file` to `true`

ENHANCEMENTS:

* makefile: Added `example` target to simplify the testing process

## 0.1.0

FEATURES:

* **New Data Source:** `ironio_projects`
* **New Data Source:** `ironio_queues`
* **New Resource:** `ironio_project`
* **New Resource:** `ironio_pull_queue`
* **New Resource:** `ironio_push_queue`
