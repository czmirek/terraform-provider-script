terraform {
  required_providers {
    script = {
      version = ">= 0.1"
      source = "czmirek/script"
    }
  }
}
resource "script" "new" {
  create = ["pwsh", "${path.root}/create.ps1"]
  read = ["pwsh", "${path.root}/read.ps1 -id={{ID}}"]
  update = ["pwsh", "${path.root}/update.ps1"]
  delete = ["pwsh", "${path.root}/delete.ps1"]
  working_dir = path.root
}