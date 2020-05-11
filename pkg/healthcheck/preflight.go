package healthcheck

import (
	"context"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// Follows a design pattern similar to the linkerd health checks
// Source: https://github.com/linkerd/linkerd2/blob/master/pkg/healthcheck/healthcheck.go#L42

// ComponentID is an identifier for the type of health check
type ComponentID string

const (
	APIServerChecks             ComponentID = "api-server"
	KubernetesVersion           ComponentID = "kubernetes-version"
	KubernetesStorageClassCheck ComponentID = "kubernetes-storage-class"
	KubeflowPreInstallChecks    ComponentID = "pre-kubeflow-setup"
	KubeflowConfigChecks        ComponentID = "kubeflow-config"
)

const (
	MaximumKubernetesMajorVersion int = 1
	MaximumKubernetesMinorVersion int = 15
)

type checker struct {
	description string
	fatal       bool
	warning     bool
	check       func(context.Context) error
}

type CheckResult struct {
	Component   ComponentID
	Description string
	Warning     bool
	Err         error
}

// CheckK8sVersion validates kubernetes target version
func CheckK8sVersion() (bool, error) {

	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	config, err := kubeConfig.ClientConfig()
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		log.Errorf("Error building clientset from client-config: %v", err)
	}

	serverVersion, err := clientset.Discovery().ServerVersion()
	if err != nil {
		log.Errorf("Error fetching server version: %v", err)
	}

	serverMajorVersion, err := strconv.Atoi(serverVersion.Major)
	if err != nil {
		log.Errorf("Error converting server version: %v", err)
		return false, err
	}

	serverMinorVersion, err := strconv.Atoi(strings.Trim(serverVersion.Minor, "+"))
	if err != nil {
		log.Errorf("Error converting server version: %v", err)
		return false, err
	}

	if serverMajorVersion > MaximumKubernetesMajorVersion {
		log.Errorf("Server version: v.%s.%s.x but expected server version: v.%d.%d.x", serverVersion.Major, serverVersion.Minor, MaximumKubernetesMajorVersion, MaximumKubernetesMinorVersion)
		return false, nil
	} else if serverMajorVersion == MaximumKubernetesMajorVersion && serverMinorVersion > MaximumKubernetesMinorVersion {
		log.Errorf("Server version: v.%s.%s.x but expected server version: v.%d.%d.x", serverVersion.Major, serverVersion.Minor, MaximumKubernetesMajorVersion, MaximumKubernetesMinorVersion)
		return false, nil
	} else {
		log.Infof("Kubernetes server version %v is compatible.", serverVersion)
	}
	return true, nil
}
