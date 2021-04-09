package virtinfra

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	vc "github.com/ploumpouloum/virtinfra-client-go"
)

func resourceVpc() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVpcCreate,
		ReadContext:   resourceVpcRead,
		UpdateContext: resourceVpcUpdate,
		DeleteContext: resourceVpcDelete,
		Schema: map[string]*schema.Schema{
			"cidr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceVpcCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*vc.Client)
	var diags diag.Diagnostics

	vpc := &vc.Vpc{
		Cidr: (vc.Cidr)(d.Get("cidr").(string)),
	}

	err := c.VpcAdd(vpc)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(string(vpc.Id))

	resourceVpcRead(ctx, d, m)

	return diags
}

func resourceVpcRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*vc.Client)
	var diags diag.Diagnostics

	vpcId := d.Id()
	vpc, err := c.VpcGet((vc.VpcId)(vpcId))
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("cidr", vpc.Cidr)

	return diags
}

func resourceVpcUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*vc.Client)

	vpcId := d.Id()
	vpc, err := c.VpcGet((vc.VpcId)(vpcId))
	if err != nil {
		return diag.FromErr(err)
	}

	vpc.Cidr = (vc.Cidr)(d.Get("cidr").(string))
	err = c.VpcUpdate(vpc)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceVpcRead(ctx, d, m)
}

func resourceVpcDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*vc.Client)
	var diags diag.Diagnostics

	vpcId := d.Id()
	err := c.VpcDelete((vc.VpcId)(vpcId))
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
