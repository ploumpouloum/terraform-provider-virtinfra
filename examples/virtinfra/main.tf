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

# Returns all vpcs
output "all_vpcs" {
  value = data.virtinfra_vpcs.all.vpcs
}
