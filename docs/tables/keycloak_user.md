# Table: keycloak_user

Obtaining basic user information from the Keycloak realm.

## Examples

### List all users

```sql
select
  *
from
  keycloak_user;
```

### Get a specific user

```sql
select
  *
from
  keycloak_user
where
  id = 'a1399321-d1a5-4e00-9034-eb8046d6a9dc';
  
-- OR

select 
  *
from
  keycloak_user
where
  username = 'testuser';

-- OR

select
  *
from
  keycloak_user
where
  email = 'testuser@example.com';
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
