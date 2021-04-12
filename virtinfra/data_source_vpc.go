package virtinfra

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	vc "github.com/ploumpouloum/virtinfra-client-go"
)

func dataSourceVpc() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVpcRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cidr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceVpcRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*vc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var vpcId string
	if cid, ok := d.GetOk("id"); ok {
		vpcId = cid.(string)
	} else {
		return diag.Errorf("Unable to parse VPC id")
	}

	vpc, err := c.VpcGet((vc.VpcId)(vpcId))
	if err != nil {
		return diag.FromErr(err)
	}
	if vpc == nil {
		return diag.Errorf("Unable to find VPC id '%s'", vpcId)
	}

	d.SetId(string(vpc.Id))
	d.Set("id", vpc.Id)
	d.Set("cidr", vpc.Cidr)

	return diags
}
