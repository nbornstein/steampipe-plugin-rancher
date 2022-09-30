select
  name,
  namespace,
  jsonb_array_length(secrets) as secrets,
  age(current_timestamp, creation_timestamp)
from
  rancher.rancher_service_account
where
  name = 'jenkins_123_123'
  and namespace = 'default'
order by
  namespace,
  name;

