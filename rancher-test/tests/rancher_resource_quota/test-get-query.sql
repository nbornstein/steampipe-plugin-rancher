select
  name,
  namespace
from
  rancher.rancher_resource_quota
where
  name = 'pods-medium'
  and namespace = 'default'
order by
  namespace,
  name;