package script

import (
	"encoding/json"
	"log"
	"os/exec"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func runScript(d *schema.ResourceData, getOutput bool, op string) (map[string]string, diag.Diagnostics) {
	var diags diag.Diagnostics

	opList := d.Get(op).([]interface{})
	workingDir := d.Get("working_dir").(string)

	if err := validateProgramAttr(opList); err != nil {
		return nil, diag.FromErr(err)
	}

	program := make([]string, len(opList))

	for i, vI := range opList {
		program[i] = vI.(string)
	}

	cmd := exec.Command(program[0], program[1:]...)
	cmd.Dir = workingDir

	if getOutput {
		resultJSON, err := cmd.Output()
		log.Printf("[TRACE] JSON output: %+v\n", string(resultJSON))
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				if exitErr.Stderr != nil && len(exitErr.Stderr) > 0 {
					return nil, diag.Errorf("failed to execute %q: %s", program[0], string(exitErr.Stderr))
				}
				return nil, diag.Errorf("command %q failed with no error message", program[0])
			} else {
				return nil, diag.Errorf("failed to execute %q: %s", program[0], err)
			}
		}
		result := map[string]string{}
		err = json.Unmarshal(resultJSON, &result)
		if err != nil {
			return nil, diag.Errorf("command %q produced invalid JSON: %s", program[0], err)
		}
		return result, diags
	}

	if err := cmd.Run(); err != nil {
		return nil, diag.FromErr(err)
	}
	return nil, diags
}
