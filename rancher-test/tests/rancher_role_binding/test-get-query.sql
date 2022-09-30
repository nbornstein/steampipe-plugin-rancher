select
  name,
  namespace,
  role_name,
  role_kind,
  subjects
from
  rancher.rancher_role_binding
where
  name = 'jenkins'
  and namespace = 'default'
order by
  namespace,
  name;

