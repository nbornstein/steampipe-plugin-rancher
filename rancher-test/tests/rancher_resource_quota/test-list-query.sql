select
  name,
  namespace
from
  rancher.rancher_resource_quota
where
  name = 'pods-medium'
order by
  namespace,
  name;