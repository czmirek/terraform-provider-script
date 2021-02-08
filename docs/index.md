---

page_title: "script Provider"
subcategory: ""
description: |-
    This provider provides a script resource which delegates the resource cycle completely to CRUD scripts that you provide.  
---

# script Provider

This provider provides a script resource which delegates the resource cycle completely to CRUD scripts that you provide.

```hcl
resource "script" "my_resource" {
    create = ["pwsh", "${path.root}/create.ps1", "-NoLogo"]
    read = ["pwsh", "${path.root}/read.ps1", "##ID##", "-NoLogo"]
    update = ["pwsh", "${path.root}/update.ps1", "##ID##", "##RES##"]
    delete = ["pwsh", "${path.root}/delete.ps1", "##ID##"]
    target_state = ["pwsh", "${path.root}/target-state.ps1", "-NoLogo"]
    working_dir = path.root
}
```

## Schema
