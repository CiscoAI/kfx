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

	"github.com/CiscoAI/kfx/pkg/kind"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for version
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "cluster",
		Short: "get the KinD cluster information",
		Long:  "To get the cluster, 'kfx get cluster'",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runE(cmd, args)
		},
	}
	return cmd
}

func runE(cmd *cobra.Command, args []string) error {
	kindKubeconfig, err := kind.CheckClusterStatus("kf-kind")
	if err != nil {
		log.Errorf("Could not fetch cluster: %v", err)
		return err
	}
	// Set Kubeconfig to created cluster
	os.Setenv("KUBECONFIG", kindKubeconfig)
	log.Infof("Cluster `kf-kind` found.")
	log.Infof("Cluster kubeconfig: %s", kindKubeconfig)
	return nil
}
