module github.com/ploumpouloum/terraform-provider-virtinfra

go 1.13

require (
	github.com/hashicorp/terraform-plugin-docs v0.4.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.5.0
	github.com/ploumpouloum/virtinfra-client-go v0.0.1
)

// replace github.com/ploumpouloum/virtinfra-client-go v0.0.1 => ../virtinfra-client-go
