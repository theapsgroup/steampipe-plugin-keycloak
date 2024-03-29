package keycloak

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v12"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableClientRole() *plugin.Table {
	return &plugin.Table{
		Name:        "keycloak_client_role",
		Description: "roles associated to a keycloak client",
		List: &plugin.ListConfig{
			Hydrate:    listClientRoles,
			KeyColumns: plugin.SingleColumn("client_id"),
		},
		Columns: clientRoleColumns(),
	}
}

func clientRoleColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "client_id",
			Description: "Unique identifier for the client",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromQual("client_id"),
		},
		{
			Name:        "id",
			Type:        proto.ColumnType_STRING,
			Description: "Unique identifier for the client role",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Name of the client role",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Description given to the client role",
		},
		{
			Name:        "composite",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the client role is a composite role (multiple combined roles)",
		},
	}
}

// Hydrate Functions
func listClientRoles(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Debug("listClientRoles", "started")
	kc, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listClientRoles", fmt.Sprintf("unable to connect to Keycloak instance: %v", err))
		return nil, fmt.Errorf("unable to connect to Keycloak instance: %v", err)
	}

	clientId := d.EqualsQualString("client_id")
	if clientId == "" {
		plugin.Logger(ctx).Error("listClientRoles", "no qualifier provided for client_id")
		return nil, fmt.Errorf("keycloak_client_role List call requires an '=' qualifier for 'client_id'")
	}

	criteria := gocloak.GetRoleParams{
		BriefRepresentation: gocloak.BoolP(true),
	}

	clientRoles, err := kc.api.GetClientRoles(ctx, kc.token.AccessToken, kc.realm, clientId, criteria)
	if err != nil {
		plugin.Logger(ctx).Error("listClientRoles", fmt.Sprintf("error obtaining client roles for client_id %s: %v", clientId, err))
		return nil, fmt.Errorf("error obtaining client roles for client_id %s: %v", clientId, err)
	}

	plugin.Logger(ctx).Debug("listClientRoles", fmt.Sprintf("obtained %d client role(s) for clientId %s", len(clientRoles), clientId))
	for _, clientRole := range clientRoles {
		d.StreamListItem(ctx, clientRole)
	}

	plugin.Logger(ctx).Debug("listClientRoles", "completed successfully")
	return nil, nil
}
