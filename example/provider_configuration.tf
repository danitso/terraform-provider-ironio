provider "ironio" {
  token = "${var.token}"
}

variable "token" {
  description = "The token"
}
