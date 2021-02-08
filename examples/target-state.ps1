
$target_state = [ordered]@{
    "propertyA" = "some text BLABLABLA";
    "propertyB" = 5887548;
    "nested" = @{
        "something" = $TRUE
    }
}
$target_state | ConvertTo-Json -Compress