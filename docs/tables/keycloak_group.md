# Table: keycloak_group

Obtaining groups from the Keycloak realm.

## Examples

### List all groups

```sql
select
  *
from
  keycloak_group;
```

### Get a specific group by name

```sql
select
  *
from
  keycloak_group
where
  name = 'my-group';
```

### Get a specific group by id

```sql
select
  *
from
  keycloak_group
where
  id = 'eb5342ee-c1bd-4ed2-b84f-fa7880539ccf';
```

### List a subset of groups which have a name containing '7'

```sql
select
  *
from
  keycloak_group
where
  name like '%7%';
```
