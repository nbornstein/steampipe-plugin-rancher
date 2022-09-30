select
  name,
  namespace,
  rules,
  ingress_class_name as class
from
  rancher.rancher_ingress
where
  name = 'ingress-wildcard-host'
  and namespace = 'default';

