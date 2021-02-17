locals {
  region = "eu-west-1"
}

provider "aws" {
  profile = "default"
  region = local.region
}
