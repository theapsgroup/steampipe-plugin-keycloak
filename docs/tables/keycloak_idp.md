# Table: keycloak_idp

Obtain information on Identity Providers configured against the current realm.

## Examples

### List all Identity Providers

```sql
select
  *
from
  keycloak_idp;
```

### List only enabled Identity Providers

```sql
select
    *
from
    keycloak_idp
where
  enabled = true;
```

### List only OIDC Identity Providers

```sql
select
    *
from
    keycloak_idp
where
  provider = 'oidc';
```
