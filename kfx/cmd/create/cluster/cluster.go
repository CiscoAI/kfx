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

package cluster

import (
	"os"
	
	"github.com/CiscoAI/create-kf-app/pkg/kind"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)


type flagpole struct {
	Name        string
	Size        string
	PipelineURI string
}

// NewCommand returns a new cobra.Command for version
func NewCommand() *cobra.Command {
	flags := &flagpole{}
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "cluster",
		Short: "Creates a KinD cluster",
		Long:  "To create the cluster, 'kfx create cluster'",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runE(flags, cmd, args)
		},
	}
	cmd.Flags().StringVar(&flags.Name, "name", "kf-app", "App Directory Name")
	// size denotes the number of components to be installed part of Kubeflow.
	// Follows t-shirt sizes - small, large (not implemented: medium)
	// small is a very minimal kubeflow - pipelines, a notebook
	// large is the full fledged deployment with istio et all.
	cmd.Flags().StringVar(&flags.Size, "size", "large", "Number of components to be installed.")
	cmd.Flags().StringVar(&flags.PipelineURI, "pipeline", "", "URI for fetching the Pipeline")
	return cmd
}

func runE(flags *flagpole, cmd *cobra.Command, args []string) error {
	// CreateKinDCluster with a default config
	err := kind.CreateKindCluster("kf-kind")
	if err != nil {
		log.Error("Error creating cluster")
		return err
	}
	kindKubeconfig, err := kind.CheckClusterStatus("kf-kind")
	if err != nil {
		return err
	}
	log.Printf("Cluster kubeconfig: %s", kindKubeconfig)

	// Set Kubeconfig to created cluster
	os.Setenv("KUBECONFIG", kindKubeconfig)
	
	log.Printf("KinD cluster Created!")
	return nil
}