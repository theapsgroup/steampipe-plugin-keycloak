package keycloak

import (
    "context"
    "fmt"
    "github.com/turbot/steampipe-plugin-sdk/plugin"
    "os"
)

func connect(ctx context.Context, d *plugin.QueryData) (interface{}, error) {
    baseUrl := os.Getenv("KEYCLOAK_ADDR")
    user := os.Getenv("KEYCLOAK_USER")
    password := os.Getenv("KEYCLOAK_PASSWORD")
    realm := os.Getenv("KEYCLOAK_REALM")

    keycloakConfig := GetConfig(d.Connection)
    if &keycloakConfig != nil {
        if keycloakConfig.BaseUrl != nil {
            baseUrl = *keycloakConfig.BaseUrl
        }
        if keycloakConfig.User != nil {
            user = *keycloakConfig.User
        }
        if keycloakConfig.Password != nil {
            password = *keycloakConfig.Password
        }
        if keycloakConfig.Realm != nil {
            realm = *keycloakConfig.Realm
        }
    }

    if baseUrl == "" {
        return nil, missingConfigOptionError("baseurl", "KEYCLOAK_ADDR")
    }

    if user == "" {
        return nil, missingConfigOptionError("user", "KEYCLOAK_USER")
    }

    if password == "" {
        return nil, missingConfigOptionError("password", "KEYCLOAK_PASSWORD")
    }

    if realm == "" {
        return nil, missingConfigOptionError("realm", "KEYCLOAK_REALM")
    }

    return nil, nil
}

func missingConfigOptionError(f string, ev string) error {
    return fmt.Errorf("configuration option '%s' or Environment Variable '%s' must be set then restart Steampipe")
}