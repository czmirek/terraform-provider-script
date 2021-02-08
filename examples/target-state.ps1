$target_state_json = @{
    "propertyA" = "some text";
    "propertyB" = 5887548;
    "nested" = @{
        "something" = true
    }

}
Write-Host ($target_state_json | ConvertTo-Json) -NoNewline