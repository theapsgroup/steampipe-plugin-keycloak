# Table: keycloak_event

Obtain event audit from the Keycloak realm.

## Examples

### List all events

```sql
select
  time,
  type,
  client_id,
  user_id,
  session_id,
  ip_address,
  details
from
  keycloak_event;
```

### List events from a specific user id

```sql
select
  time,
  type,
  client_id,
  user_id,
  session_id,
  ip_address,
  details
from
  keycloak_event
where
  user_id = '457b2129-dbcb-4b0b-bd21-7c61fb797f14';
```
