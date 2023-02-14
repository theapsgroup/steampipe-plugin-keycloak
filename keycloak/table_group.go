package keycloak

import (
	"context"
	"fmt"

	"github.com/Nerzal/gocloak/v12"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
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
	plugin.Logger(ctx).Debug("listGroups", "started")
	kc, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listGroups", fmt.Sprintf("unable to connect to Keycloak instance: %v", err))
		return nil, fmt.Errorf("unable to connect to Keycloak instance: %v", err)
	}

	criteria := gocloak.GetGroupsParams{
		Full: gocloak.BoolP(true),
	}

	groups, err := kc.api.GetGroups(ctx, kc.token.AccessToken, kc.realm, criteria)
	if err != nil {
		plugin.Logger(ctx).Error("listGroups", fmt.Sprintf("error obtaining groups: %v", err))
		return nil, fmt.Errorf("error obtaining groups: %v", err)
	}

	plugin.Logger(ctx).Debug("listGroups", fmt.Sprintf("obtained %d group(s)", len(groups)))
	for _, group := range groups {
		d.StreamListItem(ctx, group)
	}

	plugin.Logger(ctx).Debug("listGroups", "completed successfully")
	return nil, nil
}
