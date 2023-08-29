generate "backend" {
  path      = "backend.tf"
  if_exists = "overwrite_terragrunt"
  contents  = <<-EOF
    terraform {
      backend "s3" {
        bucket  = "shusann-terraform-state"
        key     = "disney-wait-time-app/dev/${path_relative_to_include()}/terraform.tfstate"
        region  = "ap-northeast-1"
        encrypt = true
      }
    }
  EOF
}

generate "provider" {
  path      = "provider.tf"
  if_exists = "overwrite_terragrunt"
  contents  = <<-EOF
    provider "aws" {
      assume_role {
        role_arn = "arn:aws:iam::863718060005:role/terragrunt"
      }
    }
  EOF
}

inputs = {
  table_name = "collector-table-dev"
}
