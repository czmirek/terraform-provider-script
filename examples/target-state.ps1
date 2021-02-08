<#
    This script outputs the target state which is just a single string.

    You should make that string a serialized JSON because terraform is able to detect the JSON
    and show the specific changes in the plan.
#>

$target_state = [ordered]@{
    "propertyA" = "some text BLABLABLA";
    "propertyB" = 5887548;
    "nested" = @{
        "something" = $TRUE
    }
}
$target_state | ConvertTo-Json -Compress