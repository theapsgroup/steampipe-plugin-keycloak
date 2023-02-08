package keycloak

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
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
			"keycloak_user":         tableUser(),
			"keycloak_user_group":   tableUserGroup(),
			"keycloak_group":        tableGroup(),
			"keycloak_group_member": tableGroupMember(),
			"keycloak_client":       tableClient(),
			"keycloak_client_role":  tableClientRole(),
			"keycloak_idp":          tableIdp(),
		},
	}

	return p
}
