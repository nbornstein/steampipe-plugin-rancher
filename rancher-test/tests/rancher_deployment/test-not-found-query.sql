select
  name,
  namespace,
  status_replicas,
  ready_replicas,
  updated_replicas,
  available_replicas,
  unavailable_replicas,
  age(current_timestamp, creation_timestamp)
from
  rancher.rancher_deployment
where
  name = '' and namespace = '';

