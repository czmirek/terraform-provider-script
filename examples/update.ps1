param(
    [string] $id,
    [string] $resource
)
$obj = @{
    "id" = $id
    "resource" = $resource
}
$obj | ConvertTo-Json | Out-File "$($id).json"
