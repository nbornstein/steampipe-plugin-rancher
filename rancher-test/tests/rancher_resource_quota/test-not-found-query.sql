select
  name,
  namespace
from
  rancher.rancher_resource_quota
where
  name = ''
  and namespace = ''
order by
  namespace,
  name;