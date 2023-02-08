package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"steampipe-plugin-keycloak/keycloak"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: keycloak.Plugin})
}
