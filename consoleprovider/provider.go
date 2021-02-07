package consoleprovider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Provider returns a schema.Provider for my provider
func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema:         map[string]*schema.Schema{},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap: map[string]*schema.Resource{
			"console": resourceConsole(),
		},
	}

	return p
}
