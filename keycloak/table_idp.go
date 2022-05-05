package keycloak

import (
    "context"
    "fmt"
    "github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
    "github.com/turbot/steampipe-plugin-sdk/v3/plugin"
    "github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableIdp() *plugin.Table {
    return &plugin.Table{
        Name: "keycloak_idp",
        Description: "Identity Providers configured against the current Keycloak realm",
        List: &plugin.ListConfig{
            Hydrate: listIdentityProviders,
        },
        Columns: idpColumns(),
    }
}

func idpColumns() []*plugin.Column {
    return []*plugin.Column{
        {
            Name: "id",
            Type: proto.ColumnType_STRING,
            Description: "Unique identifier for the Identity Provider",
            Transform: transform.FromField("InternalID"),
        },
        {
            Name: "provider",
            Type: proto.ColumnType_STRING,
            Description: "Provider identifier/type (saml/oidc/etc).",
            Transform: transform.FromField("ProviderID"),
        },
        {
            Name: "display_name",
            Type: proto.ColumnType_STRING,
            Description: "Friendly display name for the Identity Provider",
        },
        {
            Name: "alias",
            Type: proto.ColumnType_STRING,
            Description: "Alias (human-friendly id) for the Identity Provider",
        },
        {
            Name: "enabled",
            Type: proto.ColumnType_BOOL,
            Description: "Indicates if the Identity Provider is enabled",
        },
        {
            Name: "store_token",
            Type: proto.ColumnType_BOOL,
            Description: "Indicates if the token must be stored after user authentication",
        },
        {
            Name: "trust_email",
            Type: proto.ColumnType_BOOL,
            Description: "Indicates if the emails provided by the Identity Provider are trusted and do not require verification",
        },
        {
            Name: "initial_login_flow",
            Type: proto.ColumnType_STRING,
            Description: "Alias of authentication flow performed after the first login from this Identity Provider where a Keycloak account does not exist yet.",
            Transform: transform.FromField("FirstBrokerLoginFlowAlias"),
        },
        {
            Name: "normal_login_flow",
            Type: proto.ColumnType_STRING,
            Description: "Alias of authentication flow performed after the subsequent login from this Identity Provider where a Keycloak account exists.",
            Transform: transform.FromField("PostBrokerLoginFlowAlias"),
        },
    }
}

// Hydrate Functions
func listIdentityProviders(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    kc, err := connect(ctx, d)
    if err != nil {
        return nil, err
    }

    providers, err := kc.api.GetIdentityProviders(ctx, kc.token.AccessToken, kc.realm)
    if err != nil {
        return nil, fmt.Errorf("error obtaining identity prodivers for realm %s: %v", kc.realm, err)
    }

    for _, provider := range providers {
        d.StreamListItem(ctx, provider)
    }

    return nil, nil
}
