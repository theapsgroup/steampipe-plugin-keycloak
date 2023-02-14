package keycloak

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v12"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableClient() *plugin.Table {
	return &plugin.Table{
		Name:        "keycloak_client",
		Description: "keycloak clients from current realm.",
		List: &plugin.ListConfig{
			Hydrate: listClients,
		},
		Columns: clientColumns(),
	}
}

func clientColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Description: "Unique identifier for the client",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "client_id",
			Description: "Friendly name identifier for the client",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "name",
			Description: "Display name of the client",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "description",
			Description: "Long description applied to the client",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "enabled",
			Description: "Indicates if the client is enabled",
			Type:        proto.ColumnType_BOOL,
		},
		{
			Name:        "base_url",
			Description: "Base URL of the client application",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "admin_url",
			Description: "Admin/Callback URL of the client application",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "default_roles",
			Description: "Roles automatically applied to users of the client",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "default_client_scopes",
			Description: "Client scopes automatically sent with all requests from the client",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "origin",
			Description: "",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "protocol",
			Description: "SSO Protocol used by the client (openid-connect, saml, etc)",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "web_origins",
			Description: "Allowed web origins",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "consent_required",
			Description: "Indicates if consent is required for the client",
			Type:        proto.ColumnType_BOOL,
		},
		{
			Name:        "public_client",
			Description: "Indicates if the client is a public client",
			Type:        proto.ColumnType_BOOL,
		},
	}
}

// Hydrate Functions
func listClients(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	kc, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	criteria := gocloak.GetClientsParams{}

	clients, err := kc.api.GetClients(ctx, kc.token.AccessToken, kc.realm, criteria)
	if err != nil {
		return nil, fmt.Errorf("error obtaining clients: %v", err)
	}

	for _, client := range clients {
		d.StreamListItem(ctx, client)
	}

	return nil, nil
}
