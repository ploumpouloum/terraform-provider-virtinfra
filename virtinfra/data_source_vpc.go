package virtinfra

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	vc "github.com/ploumpouloum/virtinfra-client-go"
)

func dataSourceVpcs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVpcsRead,
		Schema: map[string]*schema.Schema{
			"vpcs": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"cidr": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceVpcsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*vc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	vpcs, err := c.VpcGetList()
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("vpcs", flattenVpcsData(vpcs)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenVpcsData(vpcItems []vc.Vpc) []interface{} {
	if vpcItems != nil {
		vpcis := make([]interface{}, len(vpcItems), len(vpcItems))

		for i, vpc := range vpcItems {
			vpci := make(map[string]interface{})

			vpci["id"] = vpc.Id
			vpci["cidr"] = vpc.Cidr

			vpcis[i] = vpci
		}

		return vpcis
	}

	return make([]interface{}, 0)
}
