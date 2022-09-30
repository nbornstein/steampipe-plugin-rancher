select
  name,
  namespace,
  jsonb_array_length(secrets) as secrets
from
  rancher.rancher_service_account
where
  name = 'jenkins'
  and namespace = 'default'
order by
  namespace,
  name;

