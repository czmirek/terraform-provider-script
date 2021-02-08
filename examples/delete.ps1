<#
    This scripts deletes the resource. 
    The script should accept ID of the resource in its parameters, don't forget to add ##ID## in the delete attribute arguments.

    This example is trivial, it just deletes the file before terraform deletes the resource from the state.
#>
param(
    [string] $id
)
Remove-Item "$($id).json"