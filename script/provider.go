package script

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a schema.Provider for my provider
func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema:         map[string]*schema.Schema{},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap: map[string]*schema.Resource{
			"script": resourceScript(),
		},
	}

	return p
}
