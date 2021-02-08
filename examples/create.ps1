$target_state = ./target-state.ps1
$id = "Unique_ID_generated_during_creation"
$obj = @{
    "id" = $id
    "resource" = $target_state
}
$obj_json = $obj | ConvertTo-Json
$obj_json | Out-File "$($id).json"
Write-Host $obj_json -NoNewline