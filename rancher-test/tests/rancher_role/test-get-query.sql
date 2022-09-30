select
  name,
  namespace,
  rules
from
  rancher.rancher_role
where
  name = 'jenkins'
  and namespace = 'default'
order by
  namespace,
  name;

