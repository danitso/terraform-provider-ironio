resource "ironio_push_queue" "example" {
  name = "example-push"

  error_queue   = "${ironio_pull_queue.example_push_error.name}"
  retries       = 4
  retries_delay = 30
  multicast     = true
  project_id    = "${ironio_project.example.id}"

  subscriber {
    name = "example-push-subscriber-1"
    url  = "ironmq:///${ironio_pull_queue.example.*.name[0]}"

    headers {
      "X-Push-Queue" = "${ironio_pull_queue.example.*.name[0]}"
      "X-Push-Token" = "example-push-token-1"
    }
  }

  subscriber {
    name = "example-push-subscriber-2"
    url  = "ironmq:///${ironio_pull_queue.example.*.name[1]}"

    headers {
      "X-Push-Queue" = "${ironio_pull_queue.example.*.name[1]}"
      "X-Push-Token" = "example-push-token-2"
    }
  }
}

output "resource_ironio_push_queue_example_message_count" {
  description = "The push queue's message count"
  value       = "${ironio_push_queue.example.message_count}"
}

output "resource_ironio_push_queue_example_message_count_total" {
  description = "The push queue's total message count"
  value       = "${ironio_push_queue.example.message_count_total}"
}

output "resource_ironio_push_queue_example_subscriber" {
  description = "The push queue's subscribers"
  value       = "${ironio_push_queue.example.subscriber}"
}
