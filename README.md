# Terraform script provider

This provider contains a single `"script"` resource which delegates the resource cycle completely to your local CRUD
scripts.

**THIS IS v0.1! USE ON YOUR OWN RISK!**

*Also please note that this should be used ONLY if there is no official provider for the thing you are trying to provide
with this resource (or if it doesn't work very well yet).*

```go
resource "script" "new" {
    create = ["pwsh", "${path.root}/create.ps1", "-NoLogo"]
    read = ["pwsh", "${path.root}/read.ps1", "##ID##", "-NoLogo"]
    update = ["pwsh", "${path.root}/update.ps1", "##ID##", "##RES##"]
    delete = ["pwsh", "${path.root}/delete.ps1", "##ID##"]
    target_state = ["pwsh", "${path.root}/target-state.ps1", "-NoLogo"]
    working_dir = path.root
}
```

## Usage

Except `working_dir`, all attributes accept a list of strings
expressing the script that should run, where the first element should be the interpreter
(pwsh, bash, etc.). The code handling these command is the same code that handles the
[external data resource](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/data_source).

- **target_state**: This script provides the target state of the object you want to save in the terraform state.
The target state is really just a simple string and it can be anything but I recommend to return a serialized JSON
because terraform can detect and show parsed changes of individual properties/values inside the JSON when creating the plan.
- **create**: Script to run during resource creation
  - the script has to write into output a specific JSON structure (see below)
- **read**: Script to run when obtaining the existing resource
  - placeholder `##ID##` is replaced with the resource id in the script arguments
  - the script has to write into output a specific JSON structure (see below)
- **update**: Script to run when the target state differs from the current state
  - placeholder `##ID##` is replaced with the resource id in the script arguments
  - placeholder `##RES##` is replaced with the resource string in the script arguments
  - this script should not write anything to output
- **delete**: Script to run when the resource should be deleted
  - placeholder `##ID##` is replaced with the resource id in the script arguments
  - this script should not write anything to output
- **working_dir**: Working directory of the running script

## JSON schema of create & read outputs

```json
{ 
    "id" : "ID of your resource", 
    "resource" : "the resource string" 
}
```
