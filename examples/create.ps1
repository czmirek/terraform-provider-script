<#
    This scripts creates the resource. 
    The provider does not add any additional parameters to this resource other than what you specify in the "create" attribute.
    The create script is expected to return the full JSON object with "id" and "resource"

    This example is simple: we simply read the target state, construct the JSON in the format the provider expects
    and save the JSON into a file. The "read" script is then able to pick up the file.

    You can change the file to be any external resource on the cloud or whatever you need
    that has no official provider in the terraform registry.
#>

# read target state
$target_state = ./target-state.ps1 | Out-String
$id = "Unique_ID_generated_during_creation"
$obj = [ordered]@{
    "id" = $id
    "resource" = $target_state # all you need must be saved inside this string
                               # note that even if this is a JSON serialized string,
                               # terraform can actually pick it up and show in a plan
                               # changes inside that JSON
}
$obj_json = $obj | ConvertTo-Json

# save to file
$obj_json | Out-File "$($id).json"

# write the json to output so the provider can pick it up and update the id and the state string
$obj_json