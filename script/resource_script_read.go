package script

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	readOutput, diagReturn := runScript(&scriptOptions{
		OpList:     d.Get("read").([]interface{}),
		WorkingDir: d.Get("working_dir").(string),
		GetOutput:  true,
		ParamTransform: func(value *string) {
			*value = strings.Replace(*value, "##ID##", d.Id(), -1)
		},
	})
	if diagReturn.HasError() {
		return diagReturn
	}

	model := parseOutput(readOutput)
	if len(model.ID) < 1 {
		d.SetId("")
	}

	d.Set("resource", model.Resource)

	return diag.Diagnostics{}
}
