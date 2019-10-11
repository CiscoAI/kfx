package bootstrap

import (
	"os"

	kftypes "github.com/kubeflow/kubeflow/bootstrap/v3/pkg/apis/apps"
	"github.com/kubeflow/kubeflow/bootstrap/v3/pkg/kfapp/coordinator"
	log "github.com/sirupsen/logrus"
)

// the config file for v0.6.
// Needs to be changed out manually for updating each release.

const (
	masterConfigFile = "https://raw.githubusercontent.com/kubeflow/manifests/master/kfdef/kfctl_k8s_istio.yaml"
	v06ConfigFile    = "https://raw.githubusercontent.com/kubeflow/kubeflow/v0.6-branch/bootstrap/config/kfctl_k8s_istio.0.6.2.yaml"
)

const gauntletFile = "kubeflow-context.yaml"
const gitIgnoreFileContents = "secrets/"

// InstallKubeflow connects to the Kubeflow coordinator to bootstrap and install Kubeflow on KinD
func InstallKubeflow(clusterName string, version string) error {
	// Initialize a Kubeflow application
	err := KfApply(clusterName, version)
	if err != nil {
		log.Printf("Error creating a kubeflow app: %v", err)
		return err
	}
	return nil
}

// KfApply borrows code from github.com/kubeflow/bootstrap to start the install Kubeflow
func KfApply(appName string, version string) error {
	log.Println("Kubeflow init...")
	configFilePath := ""
	if version == "latest" {
		configFilePath = masterConfigFile
	} else {
		configFilePath = v06ConfigFile
	}
	// Create a kf-app config with the app name from CLI and internal config
	kfApp, err := coordinator.BuildKfAppFromURI(configFilePath)
	if err != nil {
		log.Errorf("unable to build KfApp: %v", err)
	}
	err = kfApp.Apply(kftypes.ALL)
	if err != nil {
		log.Errorf("Unable to apply resources for KfApp", err)
		return err
	}
	return nil
}

// CreateAppDir creates a project directory for installing components
func CreateAppDir(appName string) string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting current working directory: %v", err)
	}
	destinationPath := cwd + "/" + appName + "/"
	appdirErr := CreateDirFromURI(destinationPath)
	if appdirErr != nil {
		return ""
	}
	return destinationPath
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
	err = CreateDirFromURI(appDir + "/secrets")
	err = CreateFile(appDir + "/.gitignore")
	if err != nil {
		log.Errorf("Error creating .gitignore: %v")
		return false
	}
	err = WriteToFile(appDir+"/.gitignore", gitIgnoreFileContents)
	if err != nil {
		log.Errorf("Error writing to .gitignore")
		return false
	}
	componentPath := "/app/components"
	err = CreateDirFromURI(appDir + "/notebooks")
	err = CreateDirFromURI(appDir + componentPath + "/data-download")
	err = CreateDirFromURI(appDir + componentPath + "/train")
	err = CreateDirFromURI(appDir + componentPath + "/tensorboard")
	err = CreateDirFromURI(appDir + componentPath + "/serving")
	err = CreateDirFromURI(appDir + componentPath + "/inference")
	if err != nil {
		log.Errorf("Error creating project directory: %v", err)
		return false
	}
	err = CreateFile(appDir + "/" + gauntletFile)
	if err != nil {
		log.Errorf("Creating MLGauntlet file: %v", err)
		return false
	}
	return true
}

// WriteToFile writes a string to a file.
// Used to write to gitignore
func WriteToFile(filePath string, fileContent string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fileContent)
	if err != nil {
		return err
	}

	err = file.Sync()
	if err != nil {
		return err
	}

	return nil
}

// CreateFile bootstraps the ML app directory with a config file.
func CreateFile(filePath string) error {
	_, err := os.Stat(filePath)
	// create file if not exists
	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

// CreateSymlink for connecting local appDir to the Jupyter notebook
func CreateSymlink(appDir string) error {

	return nil
}

// CreateDefaultProfile uses the profile-controller to create a user-profile
func CreateDefaultProfile() error {

	return nil
}

// CreateDefaultNotebook uses the Kubeflow Notebook Controller to create an user notebook
func CreateDefaultNotebook(notebookName string) error {

	// objectName := v1meta1.ObjectMeta{Name: notebookName}
	// dynamicClient, _ := dynamic.NewForConfig(config)

	return nil
}

// apiVersion: kubeflow.org/v1alpha1
// kind: Notebook
// metadata:
//   name: default-workspace
//   labels:
//     app: default-notebook
// spec:
//   template:
//     spec:
//       serviceAccountName: jupyter-notebook
//       containers:
//         - name: default-workspace
//           image: "notebook-image:tag"
//           volumeMounts:
//           - mountPath: /home/jovyan
//             name: default-workspace
//           resources:
//             requests:
//               cpu: "2.0"
//               memory: "4.0Gi"
//       volumes:
//       - name: default-workspace
//         persistentVolumeClaim:
//           claimName: default-workspace
