param(
    [string] $id
)
Write-Host (Get-Content "$($id).json" -Raw) -NoNewline