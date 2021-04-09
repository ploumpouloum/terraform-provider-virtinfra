package virtinfra

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	vc "github.com/ploumpouloum/virtinfra-client-go"
)

func resourceSubnet() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSubnetCreate,
		ReadContext:   resourceSubnetRead,
		UpdateContext: resourceSubnetUpdate,
		DeleteContext: resourceSubnetDelete,
		Schema: map[string]*schema.Schema{
			"cidr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vpc_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceSubnetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*vc.Client)
	var diags diag.Diagnostics

	subnet := &vc.Subnet{
		Cidr:  (vc.Cidr)(d.Get("cidr").(string)),
		VpcId: (vc.VpcId)(d.Get("vpc_id").(string)),
	}

	err := c.SubnetAdd(subnet)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(string(subnet.Id))

	resourceSubnetRead(ctx, d, m)

	return diags
}

func resourceSubnetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*vc.Client)
	var diags diag.Diagnostics

	subnetId := d.Id()
	subnet, err := c.SubnetGet((vc.SubnetId)(subnetId))
	if err != nil {
		return diag.FromErr(err)
	}

	if subnet == nil {
		d.SetId("")
		return diags
	}

	d.Set("cidr", subnet.Cidr)
	d.Set("vpc_id", subnet.VpcId)

	return diags
}

func resourceSubnetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*vc.Client)

	subnetId := d.Id()
	subnet, err := c.SubnetGet((vc.SubnetId)(subnetId))
	if err != nil {
		return diag.FromErr(err)
	}

	subnet.Cidr = (vc.Cidr)(d.Get("cidr").(string))
	subnet.VpcId = (vc.VpcId)(d.Get("vpc_id").(string))
	err = c.SubnetUpdate(subnet)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceSubnetRead(ctx, d, m)
}

func resourceSubnetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*vc.Client)
	var diags diag.Diagnostics

	subnetId := d.Id()
	err := c.SubnetDelete((vc.SubnetId)(subnetId))
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
