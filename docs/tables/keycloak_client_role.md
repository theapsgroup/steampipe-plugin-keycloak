# Table: keycloak_client_role

Obtaining client roles for clients in the Keycloak realm, however **you must specify a `client_id` (this is `id` in the `keycloak_client` table) in there where or join clause**

## Examples

### List all client roles for every client
```sql
select
  c.client_id client,
  c.name,
  cr.name role_name
from keycloak_client c
left join
  keycloak_client_role cr
on
  c.id = cr.client_id;
```

### List client roles for a single client
```sql
select 
  *
from
    keycloak_client_role
where
  client_id = '53a020bf-8ddf-4f24-9f66-c7310e116f97'
```
