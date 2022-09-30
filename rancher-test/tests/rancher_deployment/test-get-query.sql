select
  name,
  namespace,
  status_replicas,
  ready_replicas,
  updated_replicas,
  available_replicas,
  unavailable_replicas
from
  rancher.rancher_deployment
where
  name = 'nginx-deployment-test'
  and namespace = 'default';

