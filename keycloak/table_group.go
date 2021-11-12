package keycloak

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v9"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableGroup() *plugin.Table {
	return &plugin.Table{
		Name:        "keycloak_group",
		Description: "Keycloak groups from current realm.",
		List: &plugin.ListConfig{
			Hydrate: listGroups,
		},
		Columns: groupColumns(),
	}
}

func groupColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Type:        proto.ColumnType_STRING,
			Description: "Unique identifier for the group",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Name of the group",
		},
		{
			Name:        "path",
			Type:        proto.ColumnType_STRING,
			Description: "Path of the group",
		},
	}
}

// Hydrate Functions
func listGroups(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	kc, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	criteria := gocloak.GetGroupsParams{
		Full: gocloak.BoolP(true),
	}

	groups, err := kc.api.GetGroups(ctx, kc.token.AccessToken, kc.realm, criteria)
	if err != nil {
		return nil, fmt.Errorf("error obtaining groups: %v", err)
	}

	for _, group := range groups {
		d.StreamListItem(ctx, group)
	}

	return nil, nil
}
