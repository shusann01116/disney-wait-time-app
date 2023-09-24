terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~>5.11.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~>3.5.0"
    }
    docker = {
      source  = "kreuzwerker/docker"
      version = "~>3.0"
    }
  }
}
