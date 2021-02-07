package main

import (
	consoleprovider "consoleprovider/consoleprovider"

	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: consoleprovider.Provider,
	})
}
