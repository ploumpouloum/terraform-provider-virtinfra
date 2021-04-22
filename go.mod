module github.com/ploumpouloum/terraform-provider-virtinfra

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.6.0
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/ploumpouloum/virtinfra-client-go v0.0.1
	github.com/zclconf/go-cty v1.7.1 // indirect
)

// replace github.com/ploumpouloum/virtinfra-client-go v0.0.1 => ../virtinfra-client-go
