terraform {
  required_providers {
    virtinfra = {
      version = "0.0.1"
      source  = "ploumpouloum.com/ploumpouloum/virtinfra"
    }
  }
}

data "virtinfra_vpcs" "all" {}

# Returns all vpcs
output "all_vpcs" {
  value = data.virtinfra_vpcs.all.vpcs
}
