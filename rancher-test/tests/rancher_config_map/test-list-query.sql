select
  name,
  namespace,
  data.key,
  data.value
from
  rancher.rancher_config_map,
  jsonb_each(data) as data
where
  name = 'game-demo'
