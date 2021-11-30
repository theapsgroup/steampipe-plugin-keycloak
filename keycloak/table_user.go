package keycloak

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v9"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableUser() *plugin.Table {
	return &plugin.Table{
		Name:        "keycloak_user",
		Description: "Keycloak users and relevant information",
		List: &plugin.ListConfig{
			Hydrate: listUsers,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name: "first_name",
					Require: plugin.Optional,
				},
				{
					Name: "last_name",
					Require: plugin.Optional,
				},
				{
					Name: "enabled",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
		   KeyColumns: plugin.AnyColumn([]string{"id", "username", "email"}),
		   Hydrate: getUser,
		},
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

	// Set page size to `limit` if limit is less than page size.
	limit := d.QueryContext.Limit
	perPage := 1000
	if limit != nil {
		if *limit < int64(perPage) {
			perPage = int(*limit)
		}
	}

	offset := 0
	doneCount := 0
	criteria := gocloak.GetUsersParams{BriefRepresentation: BoolAddr(true)}

	// Additional Filters
	q := d.KeyColumnQuals

	if q["first_name"] != nil {
		fn := q["first_name"].GetStringValue()
		criteria.FirstName = &fn
	}
	if q["last_name"] != nil {
		ln := q["last_name"].GetStringValue()
		criteria.LastName = &ln
	}
	if q["enabled"] != nil {
		e := q["enabled"].GetBoolValue()
		criteria.Enabled = &e
	}


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
			
			// Context cancellation can be manual or limit hit
			if plugin.IsCancelled(ctx) {
				return nil, nil
			}
		}

		doneCount += len(users)
		if doneCount >= userCount {
			break
		}
		offset += perPage
	}
	return nil, nil
}

func getUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	kc, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	userId := d.KeyColumnQuals["id"].GetStringValue()
	userName := d.KeyColumnQuals["username"].GetStringValue()
	userEmail := d.KeyColumnQuals["email"].GetStringValue()
	maxReturn := 1

	if userId != "" {
		user, err := kc.api.GetUserByID(ctx, kc.token.AccessToken, kc.realm, userId)
		if err != nil {
			return nil, fmt.Errorf("error obtaining user with id: %s - %v", userId, err)
		}
		return user, nil
	} else {

		criteria := gocloak.GetUsersParams {
			BriefRepresentation: gocloak.BoolP(true),
			Email: &userEmail,
			Username: &userName,
			Max: &maxReturn,
		}

		users, err := kc.api.GetUsers(ctx, kc.token.AccessToken, kc.realm, criteria)
		if err != nil {
			return nil, fmt.Errorf("error obtaining users: %v", err)
		}

		if len(users) == 0 {
			return nil, nil
		} else {
			return users[0], nil
		}
	}
}
