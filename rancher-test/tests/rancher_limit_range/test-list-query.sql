select
  name,
  namespace
from
  rancher.rancher_limit_range
where
  name = 'cpu-limit-range'
order by
  namespace,
  name;
