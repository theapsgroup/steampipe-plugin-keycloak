package keycloak

import (
    "context"
    "github.com/turbot/steampipe-plugin-sdk/plugin"
    "github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
    p := &plugin.Plugin{
        Name: "steampipe-plugin-keycloak",
        ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
            NewInstance: ConfigInstance,
            Schema:      ConfigSchema,
        },
        DefaultTransform: transform.FromGo().NullIfZero(),
        TableMap: map[string]*plugin.Table{
            "keycloak_user":  tableUser(),
            "keycloak_group": tableGroup(),
        },
    }

    return p
}
