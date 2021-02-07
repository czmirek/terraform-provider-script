package script

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScript() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOrderCreate,
		ReadContext:   resourceOrderRead,
		UpdateContext: resourceOrderUpdate,
		DeleteContext: resourceOrderDelete,
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
			"working_dir": {
				Type:     schema.TypeString,
				Required: true,
			},
			"result": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Println("[INFO] resourceOrderRead HIT")

	result, diagReturn := runScript(d, true, "read")

	if diagReturn.HasError() {
		return diagReturn
	}

	if err := d.Set("result", result); err != nil {
		return diag.FromErr(err)
	}

	return diag.Diagnostics{}
}

func resourceOrderUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Println("[INFO] resourceOrderUpdate HIT")

	if d.HasChange("result") {
		_, diagReturn := runScript(d, false, "update")
		if diagReturn.HasError() {
			return diagReturn
		}
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}
	return resourceOrderRead(ctx, d, m)
}
func resourceOrderCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	_, diagReturn := runScript(d, false, "create")
	if diagReturn.HasError() {
		return diagReturn
	}
	uuid, err := uuid.GenerateUUID()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid)
	resourceOrderRead(ctx, d, m)
	return diag.Diagnostics{}
}

func resourceOrderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Println("[INFO] resourceOrderDelete HIT")

	_, diagReturn := runScript(d, false, "delete")
	if diagReturn.HasError() {
		return diagReturn
	}
	d.SetId("")
	return diag.Diagnostics{}
}
