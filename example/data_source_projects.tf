data "ironio_projects" "example" {
  depends_on  = ["ironio_project.example"]
  filter_name = "${ironio_project.example.name}"
}
