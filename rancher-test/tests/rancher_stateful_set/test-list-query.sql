select
  name,
  namespace,
  service_name,
  replicas
from
  rancher.rancher_stateful_set
where
  name = 'web'
order by
  namespace,
  name;
