package main

import (
	"context"
	"net/url"

	"github.com/evertrust/horizon-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"x_api_id": {
				Description: "identifiant de compte local ou nom du profil",
				Type:        schema.TypeString,
				Required:    true,
			},
			"x_api_key": {
				Description: "mot de passe compte local",
				Type:        schema.TypeString,
				Required:    true,
			},
			"endpoint": {
				Description: "url de horizon",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"horizon_certificate": resourceCertificate(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	x_api_id := d.Get("x_api_id").(string)
	x_api_key := d.Get("x_api_key").(string)
	endpoint, _ := url.Parse(d.Get("endpoint").(string))

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	c := new(horizon.Horizon)
	if (x_api_id != "") && (x_api_key != "") {
		c.Init(*endpoint, x_api_id, x_api_key)
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Invalid credentials",
		})
		return nil, diags
	}

	return c, diags
}
