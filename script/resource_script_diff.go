package script

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCustomDiff(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
	id := d.Id()

	if len(id) > 0 {

		currentState := d.Get("resource")

		targetState, _ := runScript(&scriptOptions{
			OpList:         d.Get("target_state").([]interface{}),
			WorkingDir:     d.Get("working_dir").(string),
			GetOutput:      true,
			ParamTransform: func(value *string) {},
		})

		if currentState != targetState {
			d.SetNew("resource", targetState)
		}
	}
	return nil
}
