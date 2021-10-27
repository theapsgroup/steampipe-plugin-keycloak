package keycloak

import (
    "github.com/turbot/steampipe-plugin-sdk/plugin"
    "github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type KeycloakConfig struct {
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
    return &KeycloakConfig{}
}

func GetConfig(connection *plugin.Connection) KeycloakConfig {
    if connection == nil || connection.Config == nil {
        return KeycloakConfig{}
    }

    config, _ := connection.Config.(KeycloakConfig)
    return config
}
