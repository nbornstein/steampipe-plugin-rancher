# Table: rancher_pod_disruption_budget

A Pod Disruption Budget (PDB) limits the number of Pods of a replicated application that are down simultaneously from voluntary disruptions.

## Examples

### Basic info

```sql
select
   name,
   namespace,
   min_available,
   max_unavailable,
   selector 
from
   rancher_pod_disruption_budget 
order by
   namespace,
   name;
```

### List deployments and their matching PDB 

```sql
select
  d.namespace,
  d.name,
  min_available,
  replicas 
from
  rancher_pod_disruption_budget pdb 
  inner join
    rancher_deployment d 
    on d.selector = pdb.selector 
    and d.namespace = pdb.namespace 
order by
  d.namespace,
  d.name
```
