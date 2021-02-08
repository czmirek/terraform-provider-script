<#
    This script reads the current external state of the resource.
    The script should accept ID of the resource in its parameters, don't forget to add ##ID## in the read attribute arguments.

    READ SCRIPT SPECIFICS: 
    - If the external resource DOES NO LONGER EXIST, it means it has been removed outside the terraform module
      which conflicts the target state. In this case you are expected to return the empty resource JSON
      which has empty ID. The provider can detect this and it'll add a task to the plan to recreate 
      the resource.
#>
param(
    [string] $id
)
$file = "$($id).json"

# if the file exists, simply return content
if(Test-Path $file) {
    Get-Content "$($id).json" -Raw
} else {
    # if the file does NOT exists, return the empty JSON with empty strings
    # so the provider can detect this and notify terraform that the external
    # resource is missing and should be recreated
    $empty_resource = [ordered]@{
        "id" = "";
        "resource" = "";
    } 
    $empty_resource | ConvertTo-Json
}
