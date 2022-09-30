package main

import (
	"github.com/nbornstein/steampipe-plugin-rancher"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: rancher.Plugin})
}
