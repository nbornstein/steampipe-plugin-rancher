select
  name,
  namespace,
  replicas as desired,
  ready_replicas as ready,
  available_replicas as available,
  selector,
  fully_labeled_replicas
from
  rancher.rancher_replicaset
where
  name = 'frontend'
  and namespace = 'default';

