data "ironio_push_queue" "example" {
  depends_on = ["ironio_push_queue.example"]

  name       = "${ironio_push_queue.example.name}"
  project_id = "${ironio_project.example.id}"
}

output "data_ironio_push_queue_example_message_count" {
  description = "The push queue's message count"
  value       = "${data.ironio_push_queue.example.message_count}"
}

output "data_ironio_push_queue_example_message_count_total" {
  description = "The push queue's total message count"
  value       = "${data.ironio_push_queue.example.message_count_total}"
}

output "data_ironio_push_queue_example_subscriber" {
  description = "The push queue's subscribers"
  value       = "${data.ironio_push_queue.example.subscriber}"
}
