module github.com/ploumpouloum/terraform-provider-virtinfra

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/ploumpouloum/virtinfra-client-go v0.0.2
)

// replace github.com/ploumpouloum/virtinfra-client-go v0.0.2 => ../virtinfra-client-go
