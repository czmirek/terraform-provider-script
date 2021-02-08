package script

import (
	"context"

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

		currentState := d.Get("resource")

		targetState, _ := runScript(&scriptOptions{
			OpList:         d.Get("target_state").([]interface{}),
			WorkingDir:     d.Get("working_dir").(string),
			GetOutput:      true,
			ParamTransform: func(value *string) {},
		})

		// if the current state differs from target state,
		// tell terraform that targetState is the right vaule
		if currentState != targetState {
			d.SetNew("resource", targetState)
		}
	}
	return nil
}
