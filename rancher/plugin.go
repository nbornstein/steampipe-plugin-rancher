/*
package rancher implements a steampipe plugin for rancher.

This plugin provides data that Steampipe uses to present foreign
tables that represent rancher resources.
*/
package rancher

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

const pluginName = "steampipe-plugin-rancher"

// Plugin creates this (k8s) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromGo(),
		// DefaultGetConfig: &plugin.GetConfig{
		// 	ShouldIgnoreError: isNotFoundError([]string{"ResourceNotFoundException", "NoSuchEntity"}),
		// },
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"rancher_cluster_role":            tableRancherClusterRole(ctx),
			"rancher_cluster_role_binding":    tableRancherClusterRoleBinding(ctx),
			"rancher_config_map":              tableRancherConfigMap(ctx),
			"rancher_cronjob":                 tableRancherCronJob(ctx),
			"rancher_daemonset":               tableRancherDaemonset(ctx),
			"rancher_deployment":              tableRancherDeployment(ctx),
			"rancher_endpoint":                tableRancherEndpoints(ctx),
			"rancher_endpoint_slice":          tableRancherEndpointSlice(ctx),
			"rancher_ingress":                 tableRancherIngress(ctx),
			"rancher_job":                     tableRancherJob(ctx),
			"rancher_limit_range":             tableRancherLimitRange(ctx),
			"rancher_namespace":               tableRancherNamespace(ctx),
			"rancher_network_policy":          tableRancherNetworkPolicy(ctx),
			"rancher_node":                    tableRancherNode(ctx),
			"rancher_persistent_volume":       tableRancherPersistentVolume(ctx),
			"rancher_persistent_volume_claim": tableRancherPersistentVolumeClaim(ctx),
			"rancher_pod":                     tableRancherPod(ctx),
			"rancher_pod_disruption_budget":   tableRancherPDB(ctx),
			"rancher_pod_security_policy":     tableRancherPodSecurityPolicy(ctx),
			"rancher_replicaset":              tableRancherReplicaSet(ctx),
			"rancher_replication_controller":  tableRancherReplicaController(ctx),
			"rancher_resource_quota":          tableRancherResourceQuota(ctx),
			"rancher_role":                    tableRancherRole(ctx),
			"rancher_role_binding":            tableRancherRoleBinding(ctx),
			"rancher_secret":                  tableRancherSecret(ctx),
			"rancher_service":                 tableRancherService(ctx),
			"rancher_service_account":         tableRancherServiceAccount(ctx),
			"rancher_stateful_set":            tableRancherStatefulSet(ctx),

			// "rancher_pod_template_spec":    tableRancherPodTemplateSpec(ctx),
		},
	}

	return p
}
