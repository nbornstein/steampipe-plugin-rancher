select
  name,
  rules,
  age(current_timestamp, creation_timestamp)
from
  rancher.rancher_cluster_role
where
  name = ''
order by
  name;

