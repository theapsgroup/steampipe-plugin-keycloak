# Table: keycloak_group_member

Obtaining group members from a group, however **you must specify which group** in the where or join clause using the `group_id`.

## Examples

## List members for a specific group

```sql
select 
  *
from 
    keycloak_group_member
where group_id = '0eb3fe4f-f13c-4931-865c-7d6ba5a1b9fa';
```

## List members for all groups

```sql
select
  g.name,
  g.path,
  gm.username,
  gm.email
from keycloak_group g
left join
  keycloak_group_member gm
on
  g.id = gm.group_id;
```
