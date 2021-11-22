# Table: keycloak_user_group

Obtain group memberships from explicitly defined user, **you must specify a `user_id` in the where or join clause.**

> Note: It may be worth using the `keycloak_group_member` table if you wish to obtain group memberships for all/many users.

## Examples

### List group memberships for a specific user

```sql
select
  *
from
  keycloak_user_group
where
  user_id = '1083691f-b628-4ba9-a1cb-5b058481b882';
```
