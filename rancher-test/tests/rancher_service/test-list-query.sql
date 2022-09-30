select
  name,
  namespace,
  cluster_ip
from
  rancher.rancher_service
where
  name = 'jenkins'
order by
  namespace,
  name;
