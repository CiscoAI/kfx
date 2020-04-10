package healthcheck

import (
	"context"
	"fmt"
	"strconv"

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

func CheckK8sVersion(kubeconfig string) (bool, error) {
	if kubeconfig == "" {
		log.Error("Error in kubeconfig")
		return false, fmt.Errorf("kubeconfig is not valid")
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Errorf("error building client-config from kubeconfig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Errorf("error building clientset from client-config: %v", err)
	}
	serverVersion, err := clientset.Discovery().ServerVersion()
	log.Infof("Kubernetes Server Version: %v", serverVersion)
	if serverMinorVersion, err := strconv.Atoi(serverVersion.Minor); err == nil {
		if serverMinorVersion > MaximumKubernetesMinorVersion {
			log.Errorf("Server version: v.%s.%s.x but expected server version: v.%d.%d.x", serverVersion.Major, serverVersion.Minor, MaximumKubernetesMajorVersion, MaximumKubernetesMinorVersion)
			return false, nil
		}
	}
	return true, nil
}
