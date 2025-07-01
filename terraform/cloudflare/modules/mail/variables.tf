variable "zone" {
  type = object({
    id = string
    name = string
  })
}

variable "ip" {
  type = string
}