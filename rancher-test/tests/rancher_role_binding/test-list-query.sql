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
order by
  namespace,
  name;

