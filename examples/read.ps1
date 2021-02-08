param(
    [string] $id
)

$file = "$($id).json"
if(Test-Path $file) {
    Get-Content "$($id).json" -Raw
} else {
    [ordered]@{
        "id" = "";
        "resource" = "";
    }
}
