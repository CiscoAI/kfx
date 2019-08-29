// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package init

import (
	"errors"
	"os"

	"github.com/CiscoAI/create-kf-app/pkg/bootstrap"
	"github.com/CiscoAI/create-kf-app/pkg/fetch"
	"github.com/CiscoAI/create-kf-app/pkg/kind"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type flagpole struct {
	Name        string
	PipelineURI string
}

// NewCommand is for initializing a new instance with `create-kf-app init`
func NewCommand() *cobra.Command {
	flags := &flagpole{}
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "init --name kf-app --pipeline github.com/CiscoAI/KFLab//pipelines/tf-mnist",
		Short: "Initializes a ML application for development",
		Long:  "Initialize an ML application with common steps for easier development",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runE(flags, cmd, args)
		},
	}
	cmd.Flags().StringVar(&flags.Name, "name", "kf-app", "App Directory Name")
	cmd.Flags().StringVar(&flags.PipelineURI, "pipeline", "", "URI for fetching the Pipeline")
	return cmd
}

func runE(flags *flagpole, cmd *cobra.Command, args []string) error {
	log.Println("create-kf-app init...")
	clusterName := flags.Name
	pipelineURI := flags.PipelineURI

	// CreateKinDCluster with a default config
	err := kind.CreateKindCluster(clusterName)
	if err != nil {
		log.Println("Error creating KinD cluster from config", err)
	}
	kindKubeconfig, err := kind.CheckClusterStatus(clusterName)
	if err != nil {
		log.Printf("Error fetching kubeconfig for local cluster: %v", err)
	}
	log.Printf("Cluster kubeconfig: %s", kindKubeconfig)
	// Set Kubeconfig to created cluster
	os.Setenv("KUBECONFIG", kindKubeconfig)
	isReady := kind.IsClusterReady(clusterName)
	if isReady {
		bootstrap.InstallKubeflow(clusterName)
		log.Printf("Kubeflow Installed!")
	}
	// Bootstrap ML project repo
	if pipelineURI == "" {
		isDone := bootstrap.CreateProjectStructure(clusterName)
		if !isDone {
			err := errors.New("bootstrap failed")
			log.Errorf("Error bootstrapping project: %v", err)
			return err
		}
	} else {
		fetch.GetPipeline(pipelineURI, clusterName)
	}
	return nil
}
