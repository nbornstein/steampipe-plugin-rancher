select
  name,
  rules
from
  rancher.rancher_cluster_role
where
  name = 'jenkins'
order by
  name;

