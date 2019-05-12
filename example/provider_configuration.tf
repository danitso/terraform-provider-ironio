provider "ironio" {
  auth {
    token = "${var.token}"
  }
}

variable "token" {
  description = "The token"
}
