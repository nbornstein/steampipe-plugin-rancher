select
  name,
  namespace
from
  rancher.rancher_limit_range
where
  name = 'cpu-limit-range'
  and namespace = 'default'
order by
  namespace,
  name;
