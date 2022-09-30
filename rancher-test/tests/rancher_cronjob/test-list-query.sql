select
  name,
  namespace,
  suspend
from
  rancher_cronjob
where
  name = 'hello';

