resource "ironio_pull_queue" "example" {
  count = 2

  name       = "example-pull-${count.index + 1}"
  project_id = "${ironio_project.example.id}"
}

resource "ironio_pull_queue" "example_push_error" {
  name       = "example-push-error"
  project_id = "${ironio_project.example.id}"
}

output "resource_ironio_pull_queue_example_message_count" {
  description = "The pull queues message counts"
  value       = "${ironio_pull_queue.example.*.message_count}"
}

output "resource_ironio_pull_queue_example_message_count_total" {
  description = "The pull queues total message counts"
  value       = "${ironio_pull_queue.example.*.message_count_total}"
}

output "resource_ironio_pull_queue_example_push_error_message_count" {
  description = "The error queue's message counts"
  value       = "${ironio_pull_queue.example_push_error.message_count}"
}

output "resource_ironio_pull_queue_example_push_error_message_count_total" {
  description = "The error queue's total message counts"
  value       = "${ironio_pull_queue.example_push_error.message_count_total}"
}
