# https://github.com/bridgecrewio/checkov/blob/master/checkov/rancher/checks/HostPort.py
# https://github.com/bridgecrewio/checkov/tree/master/tests/rancher/checks/example_HostPort
# https://github.com/bridgecrewio/checkov/blob/master/tests/rancher/checks/test_HostPort.py

# https://rancher.io/docs/concepts/configuration/overview/
# Don’t specify a hostPort for a Pod unless it is absolutely necessary.
# When you bind a Pod to a hostPort, it limits the number of places the
# Pod can be scheduled, because each <hostIP, hostPort, protocol> combination
# must be unique.


resource "null_resource" "namespace_monitoring" {
  provisioner "local-exec" {
    command = "kubectl create namespace monitoring"
  }
}

# deploy DS-node-exporter-FAILED
resource "null_resource" "create_daemonset" {
  depends_on = [
    null_resource.namespace_monitoring
  ]
  provisioner "local-exec" {
    command = "kubectl apply -f ${path.cwd}/daemonset.yaml"
  }
}

resource "null_resource" "delay" {
  provisioner "local-exec" {
    command = "sleep 60"
  }
}


# Delay in order to get te resource creation complete
resource "null_resource" "get_daemonsets" {
  depends_on = [
    null_resource.delay
  ]
  provisioner "local-exec" {
    command = "kubectl get daemonsets -n monitoring"
  }
}


