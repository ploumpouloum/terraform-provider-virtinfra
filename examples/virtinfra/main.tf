terraform {
  required_providers {
    virtinfra = {
      version = "0.0.1"
      source  = "ploumpouloum.com/ploumpouloum/virtinfra"
    }
  }
}

provider "virtinfra" {
      local_file_location = "/Users/benoit/Repos/virtinfra/terraform-provider-virtinfra/examples/datasets/account_1.json"
}

data "virtinfra_vpcs" "all" {}

data "virtinfra_vpc" "vpc_1234" {
  id = "1234"
}

# Returns all vpcs
output "all_vpcs" {
  value = data.virtinfra_vpcs.all.vpcs
}

output "vpc_1234" {
  value = data.virtinfra_vpc.vpc_1234
}
