package rancher

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/rancher"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"github.com/mitchellh/go-homedir"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

// GetNewClientset :: gets client for querying k8s apis for the provided context
func GetNewClientset(ctx context.Context, d *plugin.QueryData) (*rancher.Clientset, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("GetNewClientset")

	// have we already created and cached the session?
	serviceCacheKey := "k8sClient" //should probably per connection/context keys...

	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		// logger.Warn("!!!! Clientset Found in Cache !!!!")
		return cachedData.(*rancher.Clientset), nil
	}

	kubeconfig, err := getK8Config(ctx, d)
	if err != nil {
		return nil, err
	}

	// Get a rest.Config from the kubeconfig file.
	restconfig, err := kubeconfig.ClientConfig()
	if err != nil {
		return nil, err
	}

	clientset, err := rancher.NewForConfig(restconfig)
	if err != nil {
		return nil, err
	}

	// save clientset in cache
	d.ConnectionManager.Cache.Set(serviceCacheKey, clientset)

	// logger.Warn("@@@@@@@  GetNewClientset SET cache status ", "success", success)
	// time.Sleep(5000 * time.Millisecond)
	// if value, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
	// 	logger.Warn("!!!! Clientset added to cache !!!!")
	// } else {
	// 	logger.Warn("!!!! Clientset NOT Found in Cache after adding !!!!", "serviceCacheKey", serviceCacheKey, "Value", value)
	// }

	return clientset, err
}

// Get rancher config based on environment variable and plugin config
func getK8Config(ctx context.Context, d *plugin.QueryData) (clientcmd.ClientConfig, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("getK8Config")

	// have we already created and cached the session?
	cacheKey := "getK8Config" //should probably per connection/context keys...

	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(clientcmd.ClientConfig), nil
	}

	// get rancher config info
	rancherConfig := GetConfig(d.Connection)

	// Set default loader and overriding rules
	loader := &clientcmd.ClientConfigLoadingRules{}
	overrides := &clientcmd.ConfigOverrides{}

	// variable to store paths for rancher config
	// default kube config path
	var configPaths = []string{"~/.kube/config"}
	// Error: invalid configuration: no configuration has been provided, try setting KUBERNETES_MASTER environment variable

	if rancherConfig.ConfigPath != nil {
		configPaths = []string{*rancherConfig.ConfigPath}
	} else if rancherConfig.ConfigPaths != nil && len(rancherConfig.ConfigPaths) > 0 {
		configPaths = rancherConfig.ConfigPaths
	} else if v := os.Getenv("KUBE_CONFIG_PATHS"); v != "" {
		configPaths = filepath.SplitList(v)
	} else if v := os.Getenv("KUBERNETES_MASTER"); v != "" {
		configPaths = []string{v}
	}

	if len(configPaths) > 0 {
		expandedPaths := []string{}
		for _, p := range configPaths {
			path, err := homedir.Expand(p)
			if err != nil {
				return nil, err
			}

			logger.Debug("GetNewClientset", "Using kubeconfig: %s", path)
			expandedPaths = append(expandedPaths, path)
		}

		if len(expandedPaths) == 1 {
			loader.ExplicitPath = expandedPaths[0]
		} else {
			loader.Precedence = expandedPaths
		}

		// TODO -- Add other config options
		// if rancherConfig.ConfigContext != nil {
		// 	kubectx = *rancherConfig.ConfigContext
		// }
		// kubectx, ctxOk := d.GetOk("config_context")
		// authInfo, authInfoOk := d.GetOk("config_context_auth_info")
		// cluster, clusterOk := d.GetOk("config_context_cluster")
		// if ctxOk || authInfoOk || clusterOk {
		if rancherConfig.ConfigContext != nil {
			// ctxSuffix := "; overridden context"
			// if ctxOk {
			overrides.CurrentContext = *rancherConfig.ConfigContext
			// ctxSuffix += fmt.Sprintf("; overridden context ; config ctx: %s", overrides.CurrentContext)
			logger.Debug("GetNewClientset", "Using custom current context: %q", overrides.CurrentContext)
			// }

			overrides.Context = clientcmdapi.Context{}

			// TODO -- Add other config options
			// if authInfoOk {
			// 	overrides.Context.AuthInfo = authInfo.(string)
			// 	ctxSuffix += fmt.Sprintf("; auth_info: %s", overrides.Context.AuthInfo)
			// }
			// if clusterOk {
			// 	overrides.Context.Cluster = cluster.(string)
			// 	ctxSuffix += fmt.Sprintf("; cluster: %s", overrides.Context.Cluster)
			// }
			logger.Debug("GetNewClientset", "Using overidden context: %#v", overrides.Context)
		}
	}

	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loader, overrides)

	// save the config in cache
	d.ConnectionManager.Cache.Set(cacheKey, kubeconfig)

	return kubeconfig, nil
}

//// HYDRATE FUNCTIONS

func getKubectlContext(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	cacheKey := "getKubectlContext"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		// plugin.Logger(ctx).Warn("getKubectlContext", "######## CACHED CURRENT CONTEXT", cachedData.(string))
		return cachedData.(string), nil
	}

	kubeconfig, err := getK8Config(ctx, d)
	if err != nil {
		return nil, err
	}

	rawConfig, _ := kubeconfig.RawConfig()
	currentContext := rawConfig.CurrentContext

	// get rancher config info
	rancherConfig := GetConfig(d.Connection)

	// If set in plugin's (~/.steampipe/config/rancher.spc) connection profile
	if rancherConfig.ConfigContext != nil {
		currentContext = *rancherConfig.ConfigContext
	}

	// plugin.Logger(ctx).Warn("getKubectlContext", "######## CURRENT CONTEXT", currentContext)

	// save current context in cache
	d.ConnectionManager.Cache.Set(cacheKey, currentContext)

	return currentContext, nil
}

//// COMMON TRANSFORM FUNCTIONS

func v1TimeToRFC3339(_ context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}

	switch v := d.Value.(type) {
	case v1.Time:
		return v.ToUnstructured(), nil
	case *v1.Time:
		if v == nil {
			return nil, nil
		}
		return v.ToUnstructured(), nil
	default:
		return nil, fmt.Errorf("invalid time format %T! ", v)
	}
}

func labelSelectorToString(_ context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}

	selector := d.Value.(*v1.LabelSelector)

	ss, err := v1.LabelSelectorAsSelector(selector)
	if err != nil {
		return nil, err
	}

	return ss.String(), nil
}

func selectorMapToString(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("selectorMapToString")

	selector_map := d.Value.(map[string]string)

	if len(selector_map) == 0 {
		return nil, nil
	}

	selector_string := labels.SelectorFromSet(selector_map).String()

	return selector_string, nil
}

//// Other Utility functions

func isNotFoundError(err error) bool {
	return strings.HasSuffix(err.Error(), "not found")
}

func getCommonOptionalKeyQuals() []*plugin.KeyColumn {
	return []*plugin.KeyColumn{
		{Name: "name", Require: plugin.Optional},
		{Name: "namespace", Require: plugin.Optional},
	}
}

func getOptionalKeyQualWithCommonKeyQuals(otherOptionalQuals []*plugin.KeyColumn) []*plugin.KeyColumn {
	return append(otherOptionalQuals, getCommonOptionalKeyQuals()...)
}

func getCommonOptionalKeyQualsValueForFieldSelector(d *plugin.QueryData) []string {
	fieldSelectors := []string{}

	if d.KeyColumnQualString("name") != "" {
		fieldSelectors = append(fieldSelectors, fmt.Sprintf("metadata.name=%v", d.KeyColumnQualString("name")))
	}

	if d.KeyColumnQualString("namespace") != "" {
		fieldSelectors = append(fieldSelectors, fmt.Sprintf("metadata.namespace=%v", d.KeyColumnQualString("namespace")))
	}

	return fieldSelectors
}

func mergeTags(labels map[string]string, annotations map[string]string) map[string]string {
	tags := make(map[string]string)
	for k, v := range annotations {
		tags[k] = v
	}
	for k, v := range labels {
		tags[k] = v
	}
	return tags
}
