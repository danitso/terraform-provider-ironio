data "ironio_queues" "example" {
  depends_on  = ["ironio_project.example"]
  filter_name = "${ironio_project.example.name}*"
  project_id  = "${ironio_project.example.id}"
}
