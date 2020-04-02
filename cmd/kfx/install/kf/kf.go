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
	"github.com/CiscoAI/kfx/pkg/bootstrap"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type flagpole struct {
	Name    string
	Version string
}

// NewCommand returns a new cobra.Command for version
func NewCommand() *cobra.Command {
	flags := &flagpole{}
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "kf",
		Short: "Installs Kubeflow on the cluster",
		Long:  "To install to the cluster, 'kfx install kf'",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runE(flags, cmd, args)
		},
	}
	cmd.Flags().StringVar(&flags.Name, "name", "kf-app", "App Directory Name")
	cmd.Flags().StringVarP(&flags.Version, "version", "v", "v1.0.0", "Version of Kubeflow to be installed.")
	return cmd
}

func runE(flags *flagpole, cmd *cobra.Command, args []string) error {
	appName := flags.Name
	version := flags.Version
	err := bootstrap.InstallKubeflow(appName, version)
	if err != nil {
		log.Errorf("%v", err)
		return err
	}
	log.Info("Kubeflow Installed!")
	return nil
}
