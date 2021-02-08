package script

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScript() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOrderCreate,
		ReadContext:   resourceOrderRead,
		UpdateContext: resourceOrderUpdate,
		DeleteContext: resourceOrderDelete,
		CustomizeDiff: resourceCustomDiff,
		Schema: map[string]*schema.Schema{
			"read": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"create": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"update": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"delete": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"target_state": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"working_dir": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

type scriptModel struct {
	ID       string `json:"id"`
	Resource string `json:"resource"`
}
