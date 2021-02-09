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
		Description: "This provider provides a `script` resource which delegates the resource cycle completely to scripts that you provide.\n\n" +
			"- Your target state is defined by a single string, it can be anything but I recommend to use a serialized JSON.\n" +
			"- The provider excepts your scripts to work with a simple JSON structure:\n" +
			"```" +
			`
{ 
	"id" : "ID of your resource", 
	"resource" : "the resource string" 
}` + "\n```\n\n" +

			"See the [powershell example](https://github.com/czmirek/terraform-provider-script/tree/main/examples)",
		Schema: map[string]*schema.Schema{
			"read": {
				Type: schema.TypeList,
				Description: "Command to run a read script which can accept the ID of the resource and returns the external state." +
					"Placeholder ##ID## is replaced with the resource id. Example: `read = [\"pwsh\", \"${path.root}/read-my-resource.ps1\", \"##ID##\", \"-NoLogo\"]`" +
					"The read is expected to write to output a JSON with an id and resource property, see the resource description",
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"create": {
				Type:        schema.TypeList,
				Description: "Command to run a create script. Example: `create = [\"pwsh\", \"${path.root}/create.ps1\", \"-NoLogo\"]`",
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"update": {
				Type: schema.TypeList,
				Description: "Command to run an update script.\n" +
					"Placeholder `##ID##` is replaced with the resource id, `##CS##` by string resource itself.\n" +
					"Example: `update = [\"pwsh\", \"${path.root}/update.ps1\", \"##ID##\", \"##CS##\"]`\n",
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"delete": {
				Type:        schema.TypeList,
				Description: "Command to run a delete script. Placeholder `##ID##` is replaced with the resource id. Example: `delete = [\"pwsh\", \"${path.root}/delete.ps1\", \"##ID##\"]`",
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"target_state": {
				Description: "Command to output the resource string which is considered to be a target state.\n" +
					"The string can be anything but you should make it a serialized JSON of whatever data you need\n" +
					"because terraform can detect the serialized JSON and show individual property/value changes in your plan.\n\n" +
					"Example: `target_state = [\"pwsh\", \"${path.root}/target-state.ps1\", \"-NoLogo\"]`",
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"working_dir": {
				Description: "Working directory of the running scripts",
				Type:        schema.TypeString,
				Required:    true,
			},
			"resource": {
				Description: "The resource string itself. This is whatever you provide in the `target_state` script.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

type scriptModel struct {
	ID       string `json:"id"`
	Resource string `json:"resource"`
}
