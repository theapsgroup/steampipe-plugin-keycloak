---
organization: The APS Group
category: ["security"]
icon_url: "/images/plugins/theapsgroup/keycloak.svg"
brand_color: "#33C6E9"
display_name: "Keycloak"
short_name: "keycloak"
description: "Steampipe plugin for querying Keycloak users, groups and other resources."
og_description: Query Keycloak with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/theapsgroup/keycloak-social-graphic.png"
---

# Keycloak + Steampipe

[Keycloak](https://www.keycloak.org/) is an open source Identity and Access Management (IAM) solution.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

## Documentation

- [Table definitions / examples](https://hub.steampipe.io/plugins/theapsgroup/keycloak/tables)

## Getting Started

### Installation

```shell
steampipe plugin install theapsgroup/keycloak
```

### Prerequisites

- Keycloak 
- Admin Credentials for the Keycloak instance/realm.

### Configuration

> Note: Configuration file will take precedence over Env Vars.

Configuration can be done via Environment Variables or via the Configuration file `~./steampipe/config/keycloak.spc`.

Environment Variables:
- `KEYCLOAK_ADDR` : The Endpoint at which to contact your Keycloak instance (example: https://auth.example.com/ )
- `KEYCLOAK_USER` : The Username for a user with Admin privileges
- `KEYCLOAK_PASSWORD` : The password for a user with Admin privileges
- `KEYCLOAK_REALM` : The realm in the Keycloak instance you wish to query.

Configuration File:

```hcl
connection "keycloak" {
  plugin   = "theapsgroup/keycloak"
  base_url  = "https://sso.mycompany.com/"
  realm    = "example"
  user     = "my-admin-account"
  password = "eXamPl3P@$$w0rD"
}
```

## Get involved

- Open source: https://github.com/theapsgroup/steampipe-plugin-keycloak
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)