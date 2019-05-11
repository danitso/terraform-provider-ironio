output "ironio_projects_example_ids" {
  description = "The project ids"
  value       = "${data.ironio_projects.example.ids}"
}

output "ironio_projects_example_names" {
  description = "The project names"
  value       = "${data.ironio_projects.example.names}"
}
