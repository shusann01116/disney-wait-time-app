locals {
  region         = "ap-northeast-1"
  aws_account_id = "863718060005"
}

generate "backend" {
  path      = "backend.tf"
  if_exists = "overwrite_terragrunt"
  contents  = <<-EOF
    terraform {
      backend "s3" {
        bucket  = "shusann-terraform-state"
        key     = "disney-wait-time-app/sandbox/${path_relative_to_include()}/terraform.tfstate"
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
        role_arn = "arn:aws:iam::${local.aws_account_id}:role/terragrunt"
      }
    }

    data "aws_ecr_authorization_token" "token" {}

    provider "docker" {
      registry_auth {
        address  = "${local.aws_account_id}.dkr.ecr.${local.region}.amazonaws.com"
        username = data.aws_ecr_authorization_token.token.user_name
        password = data.aws_ecr_authorization_token.token.password
      }
  }
  EOF
}

inputs = {
  table_name = "collector-table-sandbox"
}
