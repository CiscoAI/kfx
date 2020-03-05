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
	"errors"

	"github.com/CiscoAI/kfx/pkg/bootstrap"
	"github.com/CiscoAI/kfx/pkg/fetch"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type flagpole struct {
	Name        string
	Size        string
	Version     string
	PipelineURI string
}

// NewCommand returns a new cobra.Command for version
func NewCommand() *cobra.Command {
	flags := &flagpole{}
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "kf",
		Short: "Installs a Kubeflow application on the cluster",
		Long:  "To install to the cluster, 'kfx install kf'",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runE(flags, cmd, args)
		},
	}
	cmd.Flags().StringVar(&flags.Name, "name", "kf-app", "App Directory Name")
	// size denotes the number of components to be installed part of Kubeflow.
	// Follows t-shirt sizes - small, large (not implemented: medium)
	// small is a very minimal kubeflow - pipelines, a notebook
	// large is the full fledged deployment with istio et all.
	cmd.Flags().StringVarP(&flags.Version, "version", "v", "v1.0.0", "Version of Kubeflow to be installed.")
	cmd.Flags().StringVar(&flags.PipelineURI, "pipeline", "", "URI for fetching the Pipeline")
	return cmd
}

func runE(flags *flagpole, cmd *cobra.Command, args []string) error {
	appName := flags.Name
	pipelineURI := flags.PipelineURI
	version := flags.Version
	// Bootstrap ML project repo
	if pipelineURI == "" {
		isDone := bootstrap.CreateProjectStructure(appName)
		if !isDone {
			err := errors.New("bootstrap failed")
			log.Errorf("Error bootstrapping project: %v", err)
			return err
		}
		log.Infof("Created a shell project structure for ML app under %v/", appName)
	} else {
		err := fetch.GetPipeline(pipelineURI, appName)
		if err != nil {
			return err
		}
		log.Infof("Bootstrapped with project %v with pipeline %v", appName, pipelineURI)
	}
	appDir := appName + "/.gitops/kfctl"
	err := bootstrap.InstallKubeflow(appDir, version)
	if err != nil {
		log.Errorf("%v", err)
		return err
	}

	log.Printf("Kubeflow Installed!")
	return nil
}
