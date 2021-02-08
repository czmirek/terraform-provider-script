package script

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

//runs the delete script
func resourceOrderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	_, diagReturn := runScript(&scriptOptions{
		OpList:     d.Get("delete").([]interface{}),
		WorkingDir: d.Get("working_dir").(string),
		GetOutput:  false,
		ParamTransform: func(value *string) {
			*value = strings.Replace(*value, "##ID##", d.Id(), -1)
		},
	})
	if diagReturn.HasError() {
		return diagReturn
	}
	return diag.Diagnostics{}
}
