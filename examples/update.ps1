param(
    [string] $id,
    [string] $resource
)
$obj = [ordered]@{
    "id" = $id
    "resource" = $resource
}
$obj | ConvertTo-Json | Out-File "$($id).json"
