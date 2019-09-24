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

package delete

import (
	"github.com/CiscoAI/create-kf-app/kfx/cmd/delete/cluster"
	"github.com/CiscoAI/create-kf-app/kfx/cmd/delete/kf"
	"github.com/spf13/cobra"
)

// ClusterName denotes the name of the KinD cluster to be deleted
var ClusterName string

// NewCommand returns a new cobra.Command for version
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "delete",
		Short: "Deletes the cluster / kf application created",
		Long:  "Deletes the cluster / KF application created",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	cmd.AddCommand(cluster.NewCommand())
	cmd.AddCommand(kf.NewCommand())
	return cmd
}
