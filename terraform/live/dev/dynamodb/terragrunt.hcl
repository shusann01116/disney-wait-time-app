terraform {
  source = "../../../modules/app//dynamodb"
}

generate "provider" {
  path = "provider.tf"
  if_exists = "overwrite_terragrunt"
  contents = <<-EOF
    provider "aws" {
      assume_role {
        role_arn = "arn:aws:iam::863718060005:role/terragrunt"
      }
    }
  EOF
}

include "root" {
  path = find_in_parent_folders()
}
