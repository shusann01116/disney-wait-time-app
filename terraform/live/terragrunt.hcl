generate "backend" {
  path      = "backend.tf"
  if_exists = "overwrite_terragrunt"
  contents  = <<-EOF
    terraform {
      backend "s3" {
        bucket = "shusann-terraform-state"
        key = "disney-wait-time-app/${path_relative_to_include()}/terraform.tfstate"
        region = "ap-northeast-1"
        encrypt = true
      }
    }
  EOF
}
