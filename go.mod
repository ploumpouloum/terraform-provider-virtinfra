module github.com/ploumpouloum/terraform-provider-virtinfra

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.0.0-rc.2
	github.com/ploumpouloum/virtinfra-client-go v0.0.0-20210406204225-ce1d4934a53a
)

replace github.com/ploumpouloum/virtinfra-client-go v0.0.0-20210406204225-ce1d4934a53a => ../virtinfra-client-go