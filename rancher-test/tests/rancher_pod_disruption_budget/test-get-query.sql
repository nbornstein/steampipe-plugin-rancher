select
  name,
  max_unavailable
from
  rancher_pod_disruption_budget
where
  name = 'zk-pdb'
  and namespace = 'default';

