select
  name,
  namespace,
  data.key,
  data.value,
  age(current_timestamp, creation_timestamp)
from
  rancher.rancher_config_map,
  jsonb_each(data) as data
where
  namespace = '' and name = '';

