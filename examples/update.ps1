<#
    This script updates the external resource in case there is some change detected between
    the external and the target state.

    The script should accept:
    - ID of the resource, don't forget to add ##ID## in the update attribute arguments
    - the resource itself, don't forget to add ##RES## in the update attribute arguments

    The example is trivial, the existing file is just overwritten with new data.
#>
param(
    [string] $id,
    [string] $resource
)
$obj = [ordered]@{
    "id" = $id
    "resource" = $resource
}
$obj | ConvertTo-Json | Out-File "$($id).json"
