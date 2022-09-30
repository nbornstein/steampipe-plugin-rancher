select
  name,
  role_name,
  role_kind,
  subjects,
  age(current_timestamp, creation_timestamp)
from
  rancher.rancher_cluster_role_binding
where
  name = '';

