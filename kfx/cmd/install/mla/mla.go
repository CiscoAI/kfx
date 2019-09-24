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

package mla

import (
	"io/ioutil"
	
	manifests "github.com/CiscoAI/create-kf-app/pkg/manifests"
	ckfutil "github.com/CiscoAI/create-kf-app/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)


// NewCommand returns a new cobra.Command for version
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "mla",
		Short: "Installs the MLAnywhere application on the cluster",
		Long:  "To install to the cluster, 'kfx install mla'",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runE(cmd, args)
		},
	}
	return cmd
}

func runE(cmd *cobra.Command, args []string) error {
	mlaConfigFile, err := manifests.Asset("manifests/mlanywhere-all-in-one.yaml")
	if err != nil {
		log.Errorf("Error loading ML Anywhere manifest")
		return err
	}
	err = ioutil.WriteFile("/tmp/mla-deploy.yaml", mlaConfigFile, 0700)
	if err != nil {
		log.Errorf("Unable to write ML Anywhere config to disk")
		return err
	}
	mlaScript, err := manifests.Asset("manifests/create-mla.sh")
	if err != nil {
		log.Errorf("Error loading ML Anywhere manifest")
		return err
	}
	err = ioutil.WriteFile("/tmp/create-mla.sh", mlaScript, 0700)
	if err != nil {
		log.Errorf("Unable to write ML Anywhere config to disk")
		return err
	}
	ckfutil.RunShellScript("/tmp/create-mla.sh")
	if err != nil {
		log.Errorf("Failed to apply ML Anywhere manifest to cluster")
		return err
	}
	return nil
}
