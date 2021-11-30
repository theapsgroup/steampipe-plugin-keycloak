# Table: keycloak_client

Obtaining client information from the Keycloak realm.

## Examples

### List all clients

```sql
select
  *
from
  keycloak_client;
```

### List all OIDC clients

```sql
select
  *
from
  keycloak_client
where
  protocol = 'openid-connect';
```

### List all Public clients

```sql
select
  *
from
  keycloak_client
where
  public = true;
```

### Get a specific client by friendly identifier

```sql
select
    *
from
    keycloak_client
where
  client_id = 'my-client';
```
