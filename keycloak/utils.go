package keycloak

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v9"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"os"
	"time"
)

type Keycloak struct {
	api   gocloak.GoCloak
	token *gocloak.JWT
	realm string
}

func connect(ctx context.Context, d *plugin.QueryData) (*Keycloak, error) {
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

	if baseUrl == "" || user == "" || password == "" || realm == "" {
		errorMsg := ""

		if baseUrl == "" {
			errorMsg += missingConfigOptionError("base_url", "KEYCLOAK_ADDR")
		}

		if user == "" {
			errorMsg += missingConfigOptionError("user", "KEYCLOAK_USER")
		}

		if password == "" {
			errorMsg += missingConfigOptionError("password", "KEYCLOAK_PASSWORD")
		}

		if realm == "" {
			errorMsg += missingConfigOptionError("realm", "KEYCLOAK_REALM")
		}

		errorMsg += "please set the required values and restart Steampipe"
		return new(Keycloak), fmt.Errorf(errorMsg)
	}

	client := new(Keycloak)
	client.api = gocloak.NewClient(baseUrl)
	client.realm = realm

	// Acquire token
	bg := context.Background()
	t, err := client.api.LoginAdmin(bg, user, password, realm)
	if err != nil {
		return new(Keycloak), fmt.Errorf("error authenticating to %s in realm %s as %s - please check credentials\n%v", baseUrl, realm, user, err)
	}

	client.token = t

	return client, nil
}

// missingConfigOptionError is a utility function for returning parts of error string
func missingConfigOptionError(f string, ev string) string {
	return fmt.Sprintf("configuration option '%s' or Environment Variable '%s' must be set.\n", f, ev)
}

// BoolAddr returns a pointer address for a bool value
func BoolAddr(b bool) *bool {
	boolVar := b
	return &boolVar
}

// Transforms
func convertTimestamp(_ context.Context, input *transform.TransformData) (interface{}, error) {
	return time.Unix(*input.Value.(*int64)/1000, 0), nil
}
