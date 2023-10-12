resource "random_pet" "this" {
  length = 2
  keepers = {
    env = var.env
  }
}
