select
  name,
  namespace,
  rules,
  age(current_timestamp, creation_timestamp)
from
  rancher.rancher_role
where
  name = 'jenkins_123_123'
  and namespace = ''
order by
  namespace,
  name;

