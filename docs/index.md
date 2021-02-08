---

page_title: "script Provider"
subcategory: ""
description: |-
    This provider provides a script resource which delegates the resource cycle completely to scripts that you provide.
    Your target state is defined by a single string, it can be anything but I recommend to use a serialized JSON.The provider excepts your scripts to work with a simple JSON structure.  
---

# script Provider
This provider provides a script resource which delegates the resource cycle completely to scripts that you provide.
Your target state is defined by a single string, it can be anything but I recommend to use a serialized JSON.The provider excepts your scripts to work with a simple JSON structure:

{ 
"id" : "ID of your resource", 
"resource" : "the resource string" 
}
See the powershell example https://github.com/czmirek/terraform-provider-script/tree/main/examples


## Schema
