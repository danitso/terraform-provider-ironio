resource "ironio_pull_queue" "example" {
  count = 2

  name       = "example-pull-${count.index + 1}"
  project_id = "${ironio_project.example.id}"
}

resource "ironio_pull_queue" "example_push_error" {
  name       = "example-push-error"
  project_id = "${ironio_project.example.id}"
}
