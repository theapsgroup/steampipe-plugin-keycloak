package keycloak

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v12"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableUserGroup() *plugin.Table {
	return &plugin.Table{
		Name:        "keycloak_user_group",
		Description: "group membership of the Keycloak user",
		List: &plugin.ListConfig{
			Hydrate:    listUserGroups,
			KeyColumns: plugin.SingleColumn("user_id"),
		},
		Columns: userGroupColumns(),
	}
}

func userGroupColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "user_id",
			Type:        proto.ColumnType_STRING,
			Description: "Unique identifier of the user",
			Transform:   transform.FromQual("user_id"),
		},
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
func listUserGroups(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	kc, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	userId := d.EqualsQualString("user_id")
	if userId == "" {
		return nil, fmt.Errorf("keycloak_user_group List call requires an '=' qualifier for 'user_id'")
	}

	criteria := gocloak.GetGroupsParams{
		Full:                gocloak.BoolP(true),
		BriefRepresentation: gocloak.BoolP(true),
	}

	groups, err := kc.api.GetUserGroups(ctx, kc.token.AccessToken, kc.realm, userId, criteria)
	if err != nil {
		return nil, fmt.Errorf("error obtaining group memberships for userId %s: %v", userId, err)
	}

	for _, group := range groups {
		d.StreamListItem(ctx, group)
	}

	return nil, nil
}
