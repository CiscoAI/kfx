package skaffold

import "log"

// SkaffoldConfig gives a generic config for people to edit and use
const SkaffoldConfig = `# Generic config for Kubeflow ML App
apiVersion: skaffold/v1alpha2
kind: Config
build:
  artifacts:
  - imageName: data-download
    workspace: components/data-download/
  - imageName: train
    workspace: components/train/
  - imageName: serving
    workspace: components/serving/
  - imageName: inference
    workspace: components/inference/
#deploy:
#  kubectl:
#    manifests:
#	- ./manifests/*`

// InitApp initializes a skaffold directory for the ML app
func InitApp() bool {
	log.Printf("InitApp called...")
	return true
}
