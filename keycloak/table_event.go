package keycloak

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v12"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableEvent() *plugin.Table {
	return &plugin.Table{
		Name:        "keycloak_event",
		Description: "Keycloak events for the connected realm.",
		List: &plugin.ListConfig{
			Hydrate: listEvents,
			KeyColumns: []*plugin.KeyColumn{
				// {
				// 	Name:    "client",
				// 	Require: plugin.Optional,
				// },
				{
					Name:    "user_id",
					Require: plugin.Optional,
				},
			},
		},
		Columns: eventColumns(),
	}
}

func listEvents(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Debug("listEvents", "started")
	kc, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listEvents", fmt.Sprintf("unable to connect to Keycloak instance: %v", err))
		return nil, fmt.Errorf("unable to connect to Keycloak instance: %v", err)
	}

	// Set page size to `limit` if limit is less than page size.
	limit := d.QueryContext.Limit
	perPage := int32(1000)
	offset := int32(0)
	if limit != nil {
		if *limit < int64(perPage) {
			perPage = int32(*limit)
			plugin.Logger(ctx).Debug("listEvents", fmt.Sprintf("limit %d is less than a page, adjusting perPage size as appropriate", limit))
		}
	}

	params := gocloak.GetEventsParams{}
	eq := d.EqualsQuals

	// if eq["client"] != nil {
	// 	c := eq["client"].GetStringValue()
	// 	params.Client = &c
	// 	plugin.Logger(ctx).Debug("listEvents", "filtering for client", c)
	// }

	if eq["user_id"] != nil {
		uid := eq["user_id"].GetStringValue()
		params.UserID = &uid
		plugin.Logger(ctx).Debug("listEvents", "filtering for user_id", uid)
	}

	for {
		params.First = &offset
		params.Max = &perPage

		events, err := kc.api.GetEvents(ctx, kc.token.AccessToken, kc.realm, params)
		if err != nil {
			plugin.Logger(ctx).Error("listEvents", fmt.Sprintf("error obtaining events: %v", err))
			return nil, fmt.Errorf("error obtaining events: %v", err)
		}

		if len(events) == 0 {
			plugin.Logger(ctx).Debug("listEvents", "current page returned 0 events, completing")
			break
		}

		for _, event := range events {
			d.StreamListItem(ctx, event)

			// Context cancellation can be manual or limit hit
			if plugin.IsCancelled(ctx) {
				plugin.Logger(ctx).Debug("listEvents", "completed via context cancellation")
				return nil, nil
			}
		}

		offset += perPage
	}

	plugin.Logger(ctx).Debug("listEvents", "completed successfully")
	return nil, nil
}

func eventColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "time",
			Description: "Timestamp at which the event occurred",
			Type:        proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromField("Time").Transform(convertTimestamp),
		},
		{
			Name:        "type",
			Description: "The type of event",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "client_id",
			Description: "Identifier of the client",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("ClientID"),
		},
		{
			Name:        "user_id",
			Description: "Identifier of the user who raised the event",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("UserID"),
		},
		{
			Name:        "session_id",
			Description: "Identifier of the session in which the event was raised",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("SessionID"),
		},
		{
			Name:        "ip_address",
			Description: "The IP address from which the event was raised",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("IPAddress"),
		},
		{
			Name:        "details",
			Description: "The details of the event",
			Type:        proto.ColumnType_JSON,
		},
	}
}
