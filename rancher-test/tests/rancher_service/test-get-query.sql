select
  name,
  namespace,
  cluster_ip,
  type,
  cluster_ips,
  ports,
  selector
from
  rancher.rancher_service
where
  name = 'jenkins'
  and namespace = 'default'
order by
  namespace,
  name;
