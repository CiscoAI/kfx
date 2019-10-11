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
	"fmt"

	"github.com/CiscoAI/kfx/pkg/bootstrap"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for version
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "mla",
		Short: "Launches the MLAnywhere UI",
		Long:  "To launch the UI, 'kfx ui mla'",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runE(cmd, args)
		},
	}
	return cmd
}

func runE(cmd *cobra.Command, args []string) error {
	log.Infof("Launching MLAnywhere UI..")
	err := bootstrap.MLAPortForwardShell()
	if err != nil {
		return fmt.Errorf("error connecting to UI: %v", err)
	}
	return nil
}
