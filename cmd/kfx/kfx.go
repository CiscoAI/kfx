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

package main

import (
	"os"

	"github.com/CiscoAI/kfx/cmd/kfx/build"
	"github.com/CiscoAI/kfx/cmd/kfx/get"
	"github.com/CiscoAI/kfx/cmd/kfx/create"
	"github.com/CiscoAI/kfx/cmd/kfx/delete"
	"github.com/CiscoAI/kfx/cmd/kfx/install"
	"github.com/CiscoAI/kfx/cmd/kfx/run"
	"github.com/CiscoAI/kfx/cmd/kfx/ui"
	"github.com/CiscoAI/kfx/cmd/kfx/version"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	logutil "sigs.k8s.io/kind/pkg/log"
)

const defaultLevel = log.WarnLevel

// Flags for the kind command
type Flags struct {
	LogLevel string
}

// NewCommand creates the root cobra command
func NewCommand() *cobra.Command {
	flags := &Flags{LogLevel: "info"}
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "kfx",
		Short: "kfx is a productivity tool Kubeflow on-premise",
		Long: `
	Installs Kubeflow to a Kubernetes cluster.
	Creates a KinD cluster.
	`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return runE(flags, cmd, args)
		},
		SilenceUsage: true,
		Version:      version.Version,
	}
	cmd.AddCommand(create.NewCommand())
	cmd.AddCommand(get.NewCommand())
	cmd.AddCommand(delete.NewCommand())
	cmd.AddCommand(install.NewCommand())
	cmd.AddCommand(build.NewCommand())
	cmd.AddCommand(run.NewCommand())
	cmd.AddCommand(ui.NewCommand())
	cmd.AddCommand(version.NewCommand())
	return cmd
}

func runE(flags *Flags, cmd *cobra.Command, args []string) error {
	level := defaultLevel
	parsed, err := log.ParseLevel(flags.LogLevel)
	if err != nil {
		log.Warnf("Invalid log level '%s', defaulting to '%s'", flags.LogLevel, level)
	} else {
		level = parsed
	}
	log.SetLevel(level)
	return nil
}

// Run runs the `kfx` root command
func Run() error {
	return NewCommand().Execute()
}

func main() {
	log.SetOutput(os.Stdout)
	// this formatter is the default, but the timestamps output aren't
	// particularly useful, they're relative to the command start
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
		// we force colors because this only forces over the isTerminal check
		// and this will not be accurately checkable later on when we wrap
		// the logger output with our logutil.StatusFriendlyWriter
		ForceColors: logutil.IsTerminal(log.StandardLogger().Out),
	})
	if err := Run(); err != nil {
		os.Exit(1)
	}
}
