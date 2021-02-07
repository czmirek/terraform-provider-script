Add-Content "output.log" -Value "$(Get-Date) READ"
$obj = @{
    "a" = "ahoj";
    "b" = "ahoj";
}

$json = $obj | ConvertTo-Json
Write-Host $json