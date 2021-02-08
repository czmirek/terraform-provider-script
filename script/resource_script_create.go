package script

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOrderCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	createOutput, diagReturn := runScript(&scriptOptions{
		OpList:         d.Get("create").([]interface{}),
		GetOutput:      true,
		WorkingDir:     d.Get("working_dir").(string),
		ParamTransform: func(value *string) {},
	})
	if diagReturn.HasError() {
		return diagReturn
	}
	model := parseOutput(createOutput)
	d.SetId(model.ID)
	d.Set("resource", model.Resource)
	return diags
}
