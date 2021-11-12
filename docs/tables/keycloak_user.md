# Table: keycloak_user

Obtaining basic user information from the Keycloak instance.

## Examples

### List all users

```sql
select
  *
from
  keycloak_user;
```

### List disabled user accounts

```sql
select
  *
from
  keycloak_user
where
  enabled = false;
```
