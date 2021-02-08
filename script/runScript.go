package script

import (
	"log"
	"os/exec"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

type scriptOptions struct {
	OpList         []interface{}
	WorkingDir     string
	GetOutput      bool
	ParamTransform func(*string)
}

func runScript(o *scriptOptions) (string, diag.Diagnostics) {
	var diags diag.Diagnostics
	l("RUNNING SCRIPT")

	opList := o.OpList
	workingDir := o.WorkingDir

	if err := validateProgramAttr(opList); err != nil {
		return "", diag.FromErr(err)
	}

	program := make([]string, len(opList))

	for i, vI := range opList {
		program[i] = vI.(string)
		o.ParamTransform(&program[i])
	}

	cmd := exec.Command(program[0], program[1:]...)
	cmd.Dir = workingDir
	lf(cmd)

	if o.GetOutput {
		resultBytes, err := cmd.Output()
		resultJSON := string(resultBytes)
		log.Printf("[TRACE] JSON output: %+v\r\n", resultJSON)
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				if exitErr.Stderr != nil && len(exitErr.Stderr) > 0 {
					return "", diag.Errorf("failed to execute %q: %s", program[0], string(exitErr.Stderr))
				}
				return "", diag.Errorf("command %q failed with no error message", program[0])
			} else {
				return "", diag.Errorf("failed to execute %q: %s", program[0], err)
			}
		}
		return resultJSON, diags
	}

	if err := cmd.Run(); err != nil {
		return "", diag.FromErr(err)
	}
	return "", diags
}
