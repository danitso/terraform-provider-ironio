provider "ironio" {
  token = "${var.token}"
}

variable "project_id" {
  description = "The IronMQ project id"
}

variable "token" {
  description = "The IronMQ token"
}
