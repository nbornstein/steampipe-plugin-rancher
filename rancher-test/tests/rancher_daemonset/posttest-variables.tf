# https://github.com/bridgecrewio/checkov/blob/master/checkov/rancher/checks/HostPort.py
# https://github.com/bridgecrewio/checkov/tree/master/tests/rancher/checks/example_HostPort
# https://github.com/bridgecrewio/checkov/blob/master/tests/rancher/checks/test_HostPort.py

# delete DS-node-exporter-FAILED
resource "null_resource" "delete_daemonset" {
  provisioner "local-exec" {
    command = "kubectl delete -f ${path.cwd}/daemonset.yaml"
  }
}

resource "null_resource" "delete-namespace_monitoring" {
  depends_on = [
    null_resource.delete_daemonset
  ]
  provisioner "local-exec" {
    command = "kubectl delete namespace monitoring"
  }
}




