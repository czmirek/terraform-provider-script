terraform {
  required_providers {
    hashicups = {
      source = "consoleprovider"
    }
  }
}
resource "console" "new" {
  read = ["pwsh", "${path.module}/read.ps1"]
  create = ["pwsh", "${path.module}/create.ps1"]
  update = ["pwsh", "${path.module}/update.ps1"]
  delete = ["pwsh", "${path.module}/delete.ps1"]
  working_dir = "${path.module}"
}