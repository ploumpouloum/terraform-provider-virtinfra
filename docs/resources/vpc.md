---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "virtinfra_vpc Resource - terraform-provider-virtinfra"
subcategory: ""
description: |-
  
---

# virtinfra_vpc (Resource)

Provides a VPC resource.

## Example Usage

```terraform
resource "virtinfra_vpc" "example" {
  cidr = "10.0.0.0/16"
}
```

## Argument Reference

* `cidr` (Required, String) The CIDR block for the VPC.

## Attribute Reference

* `id` (String) The ID of this resource.


