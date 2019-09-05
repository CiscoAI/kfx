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

package kf

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ConfigFilePath fetches the config that was used to install Kubeflow
var ConfigFilePath string

// NewCommand returns a new cobra.Command for version
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "kf",
		Short: "Deletes Kubeflow app in cluster",
		Long:  "To delete Kubeflow, 'create-kf-app delete kf'",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Printf("Deleting Kubeflow.")
			return nil
		},
	}
	cmd.Flags().StringVar(&ConfigFilePath, "config", "", "Config file used in installation.")
	return cmd
}
