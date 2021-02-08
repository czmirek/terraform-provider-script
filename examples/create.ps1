$target_state = ./target-state.ps1 | Out-String
$id = "Unique_ID_generated_during_creation"
$obj = [ordered]@{
    "id" = $id
    "resource" = $target_state
}
$obj_json = $obj | ConvertTo-Json
$obj_json | Out-File "$($id).json"
$obj_json