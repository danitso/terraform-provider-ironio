data "ironio_queues" "example" {
  depends_on = [
    "ironio_pull_queue.example",
    "ironio_pull_queue.example_push_error",
    "ironio_push_queue.example",
  ]

  filter_name = "${ironio_project.example.name}*"
  project_id  = "${ironio_project.example.id}"
}