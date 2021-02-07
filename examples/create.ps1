Add-Content "output.log" -Value "$(Get-Date) CREATE"
$obj = @{
    "id" = "Unique_resource_identifier"
    "resource" = @{
        "whatever" = "asdf"
    };
}
$json = $obj | ConvertTo-Json
Write-Host $json