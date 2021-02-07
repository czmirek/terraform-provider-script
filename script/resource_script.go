package script

import (
	"context"
	"encoding/json"
	"log"

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

type scriptModel struct {
	ID       string      `json:"id"`
	Resource interface{} `json:"resource"`
}

func resourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	result, diagReturn := runScript(d, true, "read")
	if diagReturn.HasError() {
		return diagReturn
	}

	currentResult := d.Get("result").(string)
	log.Printf("[INFO] READ RESULT %v\r\n", result)
	log.Printf("[INFO] CURRENT RESULT %v\r\n", currentResult)
	log.Printf("[INFO] RESULT EQUAL %v\r\n", result == currentResult)
	if result != currentResult {
		_, diagReturn := runScript(d, false, "update")
		if diagReturn.HasError() {
			return diagReturn
		}
	}
	return resourceOrderRead(ctx, d, m)
}

func resourceOrderCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	result, diagReturn := runScript(d, true, "create")
	if diagReturn.HasError() {
		return diagReturn
	}
	model := &scriptModel{}
	err := json.Unmarshal([]byte(result), model)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(model.ID)
	d.Set("result", result)
	return diags
}

func resourceOrderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	_, diagReturn := runScript(d, false, "delete")
	if diagReturn.HasError() {
		return diagReturn
	}
	d.SetId("")
	return diag.Diagnostics{}
}
