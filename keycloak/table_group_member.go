package keycloak

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v9"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableGroupMember() *plugin.Table {
	return &plugin.Table{
		Name:        "keycloak_group_member",
		Description: "members of the Keycloak group",
		List: &plugin.ListConfig{
			Hydrate:    listGroupMembers,
			KeyColumns: plugin.SingleColumn("group_id"),
		},
		Columns: groupMemberColumns(),
	}
}

func groupMemberColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "group_id",
			Type:        proto.ColumnType_STRING,
			Description: "The group identifier",
			Transform:   transform.FromQual("group_id"),
		},
		{
			Name:        "id",
			Type:        proto.ColumnType_STRING,
			Description: "Unique identifier for the user",
		},
		{
			Name:        "username",
			Type:        proto.ColumnType_STRING,
			Description: "Login/Username of the user",
		},
		{
			Name:        "email",
			Type:        proto.ColumnType_STRING,
			Description: "Email address of the user",
		},
		{
			Name:        "first_name",
			Type:        proto.ColumnType_STRING,
			Description: "First name of the user",
		},
		{
			Name:        "last_name",
			Type:        proto.ColumnType_STRING,
			Description: "Last name of the user",
		},
		{
			Name:        "enabled",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the user is enabled",
		},
	}
}

// Hydrate Functions
func listGroupMembers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	kc, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	q := d.KeyColumnQuals
	groupId := q["group_id"].GetStringValue()
	if groupId == "" {
		return nil, fmt.Errorf("keycloak_group_member List call requires an '=' qualifier for 'group_id'")
	}

	criteria := gocloak.GetGroupsParams{
		Full:                gocloak.BoolP(true),
		BriefRepresentation: gocloak.BoolP(true),
	}

	members, err := kc.api.GetGroupMembers(ctx, kc.token.AccessToken, kc.realm, groupId, criteria)
	if err != nil {
		return nil, fmt.Errorf("error obtaining group members for groupId %s: %v", groupId, err)
	}

	for _, member := range members {
		d.StreamListItem(ctx, member)
	}

	return nil, nil
}
