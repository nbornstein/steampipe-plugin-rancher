select
  name,
  namespace
from
  rancher.rancher_limit_range
where
  name = ''
  and namespace = ''
order by
  namespace,
  name;
