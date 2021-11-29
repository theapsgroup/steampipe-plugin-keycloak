![image](https://github.com/theapsgroup/steampipe-plugin-keycloak/raw/main/docs/keycloak-social-graphic.png)

# Keycloak Plugin for Steampipe

Use SQL to query information including Users, Groups, Clients, Roles and more from Keycloak.

- **[Get started →](https://hub.steampipe.io/plugins/theapsgroup/keycloak)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/theapsgroup/keycloak/tables)
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
- Get involved: [Issues](https://github.com/theapsgroup/steampipe-plugin-keycloak/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install theapsgroup/keycloak
```

Set up the configuration:

```shell
vi ~/.steampipe/config/keycloak.spc
```

or set the following Environment Variables

- `KEYCLOAK_ADDR` : The Endpoint at which to contact your Keycloak instance (example: https://auth.example.com/ )
- `KEYCLOAK_USER` : The Username for a user with Admin privileges
- `KEYCLOAK_PASSWORD` : The password for a user with Admin privileges
- `KEYCLOAK_REALM` : The realm in the Keycloak instance you wish to query.

Run a query:

```sql
select * from keycloak_user;
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)
- [Keycloak](https://www.keycloak.org/) - with Admin credentials.

Clone:

```sh
git clone https://github.com/theapsgroup/steampipe-plugin-keycloak.git
cd steampipe-plugin-keycloak
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make install
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/keycloak.spc
```

Try it!

```
steampipe query
> .inspect keycloak
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Credits

Keycloak API Wrapper [Nerzal/gocloak](https://github.com/Nerzal/gocloak) licensed separately using this [Apache License](https://github.com/Nerzal/gocloak/blob/main/LICENSE).
