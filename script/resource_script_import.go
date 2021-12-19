package script

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// runs the import script and returns the resource string
func resourceOrderImport(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	l("SCRIPT: Running import script")
	importOutput, diagReturn := runScript(&scriptOptions{
		OpList:     d.Get("import").([]interface{}),
		WorkingDir: d.Get("working_dir").(string),
		GetOutput:  true,
		ParamTransform: func(value *string) {
			*value = strings.Replace(*value, "##ID##", d.Id(), -1)
		},
	})
	if diagReturn.HasError() {
		return []*schema.ResourceData{}, fmt.Errorf("Failed to import %q", d.Id())
	}

	l("SCRIPT import: Parsing output...")
	model := parseOutput(importOutput)
	lf(model)
	// if the ID is empty string then it
	// signals that the external resource has been deleted
	// outside of terraform and needs to be recreated
	if len(model.ID) < 1 {
		l("SCRIPT import: Setting empty id")
		d.SetId("")
	}

	l("SCRIPT import: Setting resource...")
	d.Set("resource", model.Resource)

	return []*schema.ResourceData{d}, nil
}
