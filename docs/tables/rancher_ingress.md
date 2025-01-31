# Table: rancher_ingress

Ingress exposes HTTP and HTTPS routes from outside the cluster to services within the cluster. Traffic routing is controlled by rules defined on the Ingress resource.

## Examples

### Basic Info

```sql
select
  name,
  namespace,
  ingress_class_name as class,
  age(current_timestamp, creation_timestamp)
from
  rancher_ingress
order by
  namespace,
  name;
```

### View rules for the ingress

```sql
select
  name,
  namespace,
  jsonb_pretty(rules) as rules
from
  rancher_ingress;
```
