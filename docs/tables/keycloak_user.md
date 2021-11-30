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

### Get a specific user by id

```sql
select
  *
from
  keycloak_user
where
  id = 'a1399321-d1a5-4e00-9034-eb8046d6a9dc';
```

### Get a specific user by username

```sql
select 
  *
from
  keycloak_user
where
  username = 'testuser';
```

### Get a specific user by email address

```sql
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
