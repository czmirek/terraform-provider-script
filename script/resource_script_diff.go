package script

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// This functions calculates the difference between the
// result of the target_state script and the current state
// because I couldn't find any other way how to do this
// in Terraform SDK v2.
func resourceCustomDiff(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
	id := d.Id()

	// if the id is set then it's not creating
	// therefore the external resource exists
	// and the read script has set the "resource" value
	// and we need to compare it with
	// what the target_state script returns
	if len(id) > 0 {

		currentState := d.Get("resource").(string)

		l("SCRIPT: Running target state script")
		targetState, _ := runScript(&scriptOptions{
			OpList:     d.Get("target_state").([]interface{}),
			WorkingDir: d.Get("working_dir").(string),
			GetOutput:  true,
			ParamTransform: func(value *string) {
				*value = strings.Replace(*value, "##CS##", currentState, -1)
			},
		})

		// if the current state differs from target state,
		// tell terraform that targetState is the right vaule
		if currentState != targetState {
			l("SCRIPT: Difference detected")
			d.SetNew("resource", targetState)
		} else {
			l("SCRIPT: Difference not detected")
		}
	}
	return nil
}
