data "ironio_queues" "example_pull" {
  depends_on = [
    "ironio_pull_queue.example",
    "ironio_pull_queue.example_push_error",
  ]

  project_id = "${ironio_project.example.id}"

  filter {
    name = "${ironio_project.example.name}*"
    pull = true
    push = false
  }
}

data "ironio_queues" "example_push" {
  depends_on = [
    "ironio_push_queue.example",
  ]

  project_id = "${ironio_project.example.id}"

  filter {
    name = "${ironio_project.example.name}*"
    pull = false
    push = true
  }
}

output "data_ironio_queues_example_pull_names" {
  description = "The queue names"
  value       = "${data.ironio_queues.example_pull.names}"
}

output "data_ironio_queues_example_pull_types" {
  description = "The queue types"
  value       = "${data.ironio_queues.example_pull.types}"
}

output "data_ironio_queues_example_push_names" {
  description = "The queue names"
  value       = "${data.ironio_queues.example_push.names}"
}

output "data_ironio_queues_example_push_types" {
  description = "The queue types"
  value       = "${data.ironio_queues.example_push.types}"
}
