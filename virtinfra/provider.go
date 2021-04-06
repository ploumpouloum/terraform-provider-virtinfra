package virtinfra

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	vc "github.com/ploumpouloum/virtinfra-client-go"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"local_file_location": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VIRTINFRA_LOCAL_FILE_LOCATION", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"virtinfra_vpcs": dataSourceVpcs(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	localFileLocation := d.Get("local_file_location").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if localFileLocation == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create VirtInfra client",
			Detail:   "Unable to create VirtInfra client if local_file_location is empty",
		})
		return nil, diags
	}

	client, err := vc.OpenClientFromLocalStorage(localFileLocation)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to open account in VirtInfra client",
			Detail:   err.Error(),
		})
		return nil, diags
	}

	return client, diags
}
