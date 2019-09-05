package bootstrap

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/CiscoAI/create-kf-app/pkg/manifests"
	kftypes "github.com/kubeflow/kubeflow/bootstrap/v3/pkg/apis/apps"
	kfdefsv3 "github.com/kubeflow/kubeflow/bootstrap/v3/pkg/apis/apps/kfdef/v1alpha1"
	"github.com/kubeflow/kubeflow/bootstrap/v3/pkg/kfapp/coordinator"
	log "github.com/sirupsen/logrus"
)

// InstallKubeflow connects to the Kubeflow coordinator to bootstrap and install Kubeflow on KinD
func InstallKubeflow(clusterName string) {
	// // Create a project directory
	//err := CreateAppDir(clusterName)
	// if err != nil {
	// 	log.Printf("Erroe creating project directory: %v", err)
	// }
	// Initialize a Kubeflow application
	err := KindKfApply(clusterName)
	if err != nil {
		log.Printf("Error creating a kubeflow app: %v", err)
	}
}

// KindKfApply borrows code from github.com/kubeflow/bootstrap to start the Kubeflow install process
func KindKfApply(appName string) error {
	log.Println("Kubeflow init...")
	// Get config from static file
	configFile, err := manifests.Asset("manifests/kfctl_k8s_kind.yaml")
	if err != nil {
		log.Errorln("Error loading KfDef for Kubeflow")
		return err
	}
	err = ioutil.WriteFile("/tmp/kind-config.yaml", configFile, 0700)
	if err != nil {
		return err
	}
	configFilePath := "/tmp/kind-config.yaml"

	// Create a kf-app config with the app name from CLI and internal config
	kfDef := &kfdefsv3.KfDef{}
	kfDef, err = kfdefsv3.LoadKFDefFromURI(configFilePath)
	if err != nil {
		log.Printf("Unable to create KfDef from config file: %v", err)
	}
	if kfDef.Name != "" {
		log.Warnf("Overriding KfDef.Spec.Name; old value %v; new value %v", kfDef.Name, appName)
	}
	kfDef.Name = appName
	isValid, msg := kfDef.IsValid()
	if !isValid {
		log.Printf("Invalid kfdef: %v", isValid)
		log.Printf("Error validating generated KfDef, please check config file validity: %v", msg)
	}
	kfDef.Spec.AppDir = CreateAppDir(appName)
	if kfDef.Spec.AppDir == "" {
		return errors.New("kfDef App Dir not set")
	}
	log.Warnf("App directory name: %v", kfDef.Spec.AppDir)
	cfgFilePath, err := coordinator.CreateKfAppCfgFile(kfDef)
	if err != nil {
		return err
	}

	log.Printf("Syncing Cache")
	err = kfDef.SyncCache()
	if err != nil {
		log.Errorf("Failed to synchronize the cache; error: %v", err)
		return err
	}
	// Save app.yaml because we need to preserve information about the cache.
	if err := kfDef.WriteToFile(cfgFilePath); err != nil {
		log.Errorf("Failed to save KfDef to %v; error %v", cfgFilePath, err)
		return err
	}
	log.Warnf("Saved configfile as kfdef in path: %v", cfgFilePath)

	// Load KfApp for Generate and Apply
	KfApp, KfErr := coordinator.LoadKfAppCfgFile(cfgFilePath)
	if KfErr != nil {
		log.Printf("Error loading KfApp from configfilepath: %v", KfErr)
	}
	// Once init is done, we generate and apply subsequently
	kfResource := kftypes.K8S
	log.Println("Kubeflow Generate...")
	generateErr := KfApp.Generate(kfResource)
	if generateErr != nil {
		log.Println("Unable to generate resources for KfApp", generateErr)
		return generateErr
	}
	log.Println("Kubeflow Apply...")
	applyErr := KfApp.Apply(kfResource)
	if applyErr != nil {
		log.Println("Unable to apply resources for KfApp", applyErr)
	}
	return nil
}

// CreateAppDir creates a project directory for installing components
func CreateAppDir(appName string) string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting current working directory: %v", err)
	}
	appdirErr := CreateDirFromURI(cwd + "/" + appName)
	if appdirErr != nil {
		return ""
	}
	return cwd + "/" + appName
}

// CreateDirFromURI is common utility function to create a directory given a path
func CreateDirFromURI(dirPath string) error {
	dirErr := os.MkdirAll(dirPath, os.ModePerm)
	if dirErr != nil {
		log.Errorf("Could not create directory: %v", dirErr)
		return dirErr
	}
	return nil
}

// CreateProjectStructure bootstraps a common workflow structure with
// data download, training, tensorboard, sering, inference components and
// strings them together in a single workflow
func CreateProjectStructure(appName string) bool {
	cwd, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting current working directory: %v", err)
		return false
	}
	appDir := cwd + "/" + appName
	_, err = os.Stat(appDir)
	if os.IsNotExist(err) {
		_ = CreateAppDir(appName)
	}
	err = CreateDirFromURI(appDir + "/components/data-download")
	err = CreateDirFromURI(appDir + "/components/train")
	err = CreateDirFromURI(appDir + "/components/tensorboard")
	err = CreateDirFromURI(appDir + "/components/serving")
	err = CreateDirFromURI(appDir + "/components/inference")
	if err != nil {
		log.Errorf("Error creating project directory: %v", err)
		return false
	}
	return true
}

// CreateDefaultNotebook uses the Kubeflow Notebook Controller to create an user notebook
func CreateDefaultNotebook(appName string) bool {

	return true
}
