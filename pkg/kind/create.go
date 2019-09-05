package kind

import (
	"io/ioutil"
	"time"

	"github.com/CiscoAI/create-kf-app/pkg/manifests"
	kfutil "github.com/CiscoAI/create-kf-app/pkg/util"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/kind/pkg/cluster"
	"sigs.k8s.io/kind/pkg/cluster/config/encoding"
	"sigs.k8s.io/kind/pkg/cluster/create"
	"sigs.k8s.io/kind/pkg/util"
)

// StorageClasses implements StorageClassInterface
type StorageClasses struct {
	client rest.Interface
}

// CreateKindCluster bootstraps a kubernetes in docker cluster
func CreateKindCluster(ClusterName string) error {
	// Check if the cluster name already exists
	known, err := cluster.IsKnown(ClusterName)
	if err != nil {
		return err
	}
	if known {
		log.Printf("Skipping creating a cluster with the name %q, because it already exists", ClusterName)
		return nil
	}

	// Create a context for the cluster and validate it
	ctx := cluster.NewContext(ClusterName)
	KubeconfigPath := ctx.KubeConfigPath()
	log.Printf("Creating KinD cluster with the kubeconfig placed at: %s", KubeconfigPath)
	// Set Config for cluster
	// TODO: re-write to use a static config file or generate the config on the fly.
	// MUSTDO: really, re-write this part.
	configFileBytes, err := manifests.Asset("manifests/kind-config.yaml")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("/tmp/kind-config.yaml", configFileBytes, 0700)
	if err != nil {
		return err
	}
	configFilePath := "/tmp/kind-config.yaml"
	log.Printf("Config File Path is %v", configFilePath)
	if err != nil {
		return err
	}
	cfg, err := encoding.Load(configFilePath)
	if err != nil {
		return errors.Wrap(err, "error loading config")
	}
	err = cfg.Validate()
	if err != nil {
		log.Error("Invalid configuration")
		configErrors := err.(util.Errors)
		for _, problem := range configErrors.Errors() {
			log.Error(problem)
		}
		return errors.New("aborting due to invalid configuration")
	}
	imageName := "kindest/node:v1.15.0@sha256:b4d092fd2b507843dd096fe6c85d06a27a0cbd740a0b32a880fe61aba24bb478"

	for node := range cfg.Nodes {
		cfg.Nodes[node].Image = imageName
	}
	err = cfg.Validate()
	if err != nil {
		log.Errorf("Invalid configuration, invalid image name : %v", err)
		return errors.New("aborting due to invalid configuration")
	}

	// Create Cluster takes a bunch of config options this is to create them
	// and pass them along
	if err = ctx.Create(
		cfg,
		create.Retain(true),
		create.WaitForReady(time.Duration(60)*time.Second),
	); err != nil {
		if utilErrors, ok := err.(util.Errors); ok {
			for _, problem := range utilErrors.Errors() {
				log.Error(problem)
			}
			return errors.New("aborting due to invalid configuration")
		}
		return errors.Wrap(err, "failed to create cluster")
	}
	return nil
}

// CheckClusterStatus checks and returns the kubeconfig if KF cluster exists
func CheckClusterStatus(ClusterName string) (string, error) {

	// Check if the cluster name already exists
	known, err := cluster.IsKnown(ClusterName)
	if err != nil {
		return "", err
	}
	if known {
		log.Println("Cluster already exists")
	}
	clusterContext := cluster.NewContext(ClusterName)
	return clusterContext.KubeConfigPath(), nil
}

// IsClusterReady iterates through the nodes and checks if the cluster is ready
func IsClusterReady(ClusterName string) bool {
	time.Sleep(time.Duration(15) * time.Second)

	var nodes []*v1.Node
	clusterKubeConfig, err := CheckClusterStatus(ClusterName)
	if err != nil {
		log.Errorf("Error getting cluster status: %v", err)
	}
	clusterConfig, err := clientcmd.BuildConfigFromFlags("", clusterKubeConfig)
	if err != nil {
		log.Printf("Cluster config can't be created: %v", err)
	}
	clientSet, err := kubernetes.NewForConfig(clusterConfig)
	if err != nil {
		log.Printf("Cluster clientset creation error: %v", err)
	}
	nodeItems, err := clientSet.Core().Nodes().List(metav1.ListOptions{})
	if err != nil {
		log.Errorf("Error listing nodes: %v", err)
	}
	for i := range nodeItems.Items {
		node := nodeItems.Items[i]
		nodes = append(nodes, &node)
	}
	for _, node := range nodes {
		if !IsNodeReady(node) {
			return false
		}
	}
	// Replace Storage Class with local-path-provisioner
	err = ReplaceStorageClass(clientSet)
	if err != nil {
		log.Printf("Error replacing StorageClass: %v", err)
	}
	return true
}

// IsNodeReady checks the status for each node in the cluster
func IsNodeReady(node *v1.Node) bool {
	for i := range node.Status.Conditions {
		nodeCondition := &node.Status.Conditions[i]
		if nodeCondition.Type == v1.NodeReady && nodeCondition.Status != v1.ConditionTrue {
			return false
		}
	}
	return true
}

// ReplaceStorageClass removes the default storage class that comes with KinD and
// installs the local-path-provisioner from Rancher. This is done because the default
// StorageClass does not work for minio which in turn makes the pipelines deployment fail
func ReplaceStorageClass(ClusterClient *kubernetes.Clientset) error {
	StorageClass, err := ClusterClient.StorageV1().StorageClasses().Get("standard", metav1.GetOptions{})
	if err != nil {
		log.Printf("Error fetching StorageClass 'standard': %v", err)
		return err
	}
	className := StorageClass.Name
	log.Printf("StorageClass %v found!", className)
	err = ClusterClient.StorageV1().StorageClasses().Delete("standard", &metav1.DeleteOptions{})
	if err != nil {
		log.Printf("Error deleting Storage Class: %v, error: %v", StorageClass.Name, err)
	}
	log.Printf("StorageClass %v deleted!", className)

	scScript, err := manifests.Asset("manifests/apply-storageclass.sh")
	if err != nil {
		log.Errorln("error fetching storageclass apply script")
		return err
	}
	err = ioutil.WriteFile("/tmp/apply-storageclass.sh", scScript, 0700)

	scManifest, err := manifests.Asset("manifests/local-path-storage.yaml")
	if err != nil {
		log.Errorln("error fetching storageclass manifest")
		return err
	}
	err = ioutil.WriteFile("/tmp/local-path-storage.yaml", scManifest, 0700)

	kfutil.RunShellScript("/tmp/apply-storageclass.sh")
	return nil
}
