terraform {
  required_providers {
    script = {
      version = ">= 0.1"
      source = "czmirek/script"
    }
  }
}
resource "script" "new" {
  create = ["pwsh", "${path.root}/create.ps1", "-NoLogo"]
  read = ["pwsh", "${path.root}/read.ps1", "{{ID}}", "-NoLogo"]
  update = ["pwsh", "${path.root}/update.ps1", "{{ID}}", "{{RES}}"]
  delete = ["pwsh", "${path.root}/delete.ps1", "{{ID}}"]
  target_state = ["pwsh", "${path.root}/target-state.ps1", "-NoLogo"]
  working_dir = path.root
}