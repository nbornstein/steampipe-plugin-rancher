# Table: rancher_service

In Kubernetes, service are used to provide an abstract way to expose an application running on a set of Pods as a network service. Kubernetes gives Pods their own IP addresses and a single DNS name for a set of Pods, and can load-balance across them.

## Examples

### Basic Info - `kubectl describe service --all-namespaces` columns

```sql
select
  name,
  namespace,
  type,
  cluster_ip,
  age(current_timestamp, creation_timestamp)
from
  rancher_service
order by
  namespace,
  name;
```
