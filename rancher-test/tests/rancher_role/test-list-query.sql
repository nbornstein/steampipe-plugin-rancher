select
  name,
  namespace,
  rules
from
  rancher.rancher_role
where
  name = 'jenkins'
order by
  namespace,
  name;

