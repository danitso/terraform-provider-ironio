data "ironio_pull_queue" "example" {
  count      = 2
  depends_on = ["ironio_pull_queue.example"]

  name       = "${ironio_pull_queue.example.*.name[count.index]}"
  project_id = "${ironio_project.example.id}"
}

data "ironio_pull_queue" "example_push_error" {
  depends_on = ["ironio_pull_queue.example_push_error"]

  name       = "${ironio_pull_queue.example_push_error.name}"
  project_id = "${ironio_project.example.id}"
}

output "data_ironio_pull_queue_example_message_count" {
  description = "The pull queues message counts"
  value       = "${data.ironio_pull_queue.example.*.message_count}"
}

output "data_ironio_pull_queue_example_message_count_total" {
  description = "The pull queues total message counts"
  value       = "${data.ironio_pull_queue.example.*.message_count_total}"
}

output "data_ironio_pull_queue_example_push_error_message_count" {
  description = "The error queue's message count"
  value       = "${data.ironio_pull_queue.example_push_error.message_count}"
}

output "data_ironio_pull_queue_example_push_error_message_count_total" {
  description = "The error queue's total message count"
  value       = "${data.ironio_pull_queue.example_push_error.message_count_total}"
}
