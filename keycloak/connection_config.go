package keycloak

import (
    "github.com/turbot/steampipe-plugin-sdk/plugin"
    "github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type PluginConfig struct {
    BaseUrl  *string `cty:"baseurl"`
    User     *string `cty:"user"`
    Password *string `cty:"password"`
    Realm    *string `cty:"realm"`
}

var ConfigSchema = map[string]*schema.Attribute{
    "baseurl": {
        Type: schema.TypeString,
    },
    "user": {
        Type: schema.TypeString,
    },
    "password": {
        Type: schema.TypeString,
    },
    "realm": {
        Type: schema.TypeString,
    },
}

func ConfigInstance() interface{} {
    return &PluginConfig{}
}

func GetConfig(connection *plugin.Connection) PluginConfig {
    if connection == nil || connection.Config == nil {
        return PluginConfig{}
    }

    config, _ := connection.Config.(PluginConfig)
    return config
}
