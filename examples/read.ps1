param(
    [string] $id
)
Add-Content "output.log" -Value "$(Get-Date) READ"
$obj = @{
    "id" = "Unique_resource_identifier"
    "resource" = @{
        "whatever" = "asddgasdsdfasdfgasdgf"
    };
}
$json = $obj | ConvertTo-Json
Write-Host $json