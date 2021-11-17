package keycloak

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v9"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

// TODO: Optional KeyColumns to Filter more efficiently
// TODO: Increase Perf...

func tableUser() *plugin.Table {
	return &plugin.Table{
		Name:        "keycloak_user",
		Description: "Keycloak users and relevant information",
		List: &plugin.ListConfig{
			Hydrate: listUsers,
		},
		//Get: &plugin.GetConfig{
		//    KeyColumns: plugin.AnyColumn([]string{"username", "email"}),
		//    Hydrate: getUser,
		//},
		Columns: userColumns(),
	}
}

func userColumns() []*plugin.Column {
	return []*plugin.Column{
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
		{
			Name:        "email_verified",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the user has verified their email address",
		},
		{
			Name:        "created_timestamp",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "Timestamp of user creation",
			Transform:   transform.FromField("CreatedTimestamp").Transform(convertTimestamp),
		},
	}
}

// Hydrate Functions
func listUsers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	kc, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	perPage := 1000
	offset := 0
	doneCount := 0
	criteria := gocloak.GetUsersParams{BriefRepresentation: BoolAddr(true)}

	userCount, err := kc.api.GetUserCount(ctx, kc.token.AccessToken, kc.realm, criteria)
	if err != nil {
		return nil, fmt.Errorf("error obtaining user count: %v", err)
	}

	for {
		criteria.Max = &perPage
		criteria.First = &offset

		users, err := kc.api.GetUsers(ctx, kc.token.AccessToken, kc.realm, criteria)
		if err != nil {
			return nil, fmt.Errorf("error obtaining users: %v", err)
		}

		if len(users) == 0 {
			break
		}

		for _, user := range users {
			d.StreamListItem(ctx, user)
		}

		doneCount += len(users)
		if doneCount >= userCount {
			break
		}
		offset += perPage
	}
	return nil, nil
}
