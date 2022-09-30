package rancher

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/schema"
)

type rancherConfig struct {
	ConfigPaths   []string `cty:"config_paths"`
	ConfigPath    *string  `cty:"config_path"`
	ConfigContext *string  `cty:"config_context"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"config_paths": {
		Type: schema.TypeList,
		Elem: &schema.Attribute{Type: schema.TypeString},
	},
	"config_path": {
		Type: schema.TypeString,
	},
	"config_context": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &rancherConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) rancherConfig {
	if connection == nil || connection.Config == nil {
		return rancherConfig{}
	}
	config, _ := connection.Config.(rancherConfig)
	return config
}
