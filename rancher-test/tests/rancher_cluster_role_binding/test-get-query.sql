select
  name,
  role_name,
  role_kind,
  subjects
from
  rancher.rancher_cluster_role_binding
where
  name = 'jenkins'
order by
  name;

