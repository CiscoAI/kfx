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

	"github.com/CiscoAI/kfx/cmd/kfx/check"
	"github.com/CiscoAI/kfx/cmd/kfx/delete"
	"github.com/CiscoAI/kfx/cmd/kfx/install"
	"github.com/CiscoAI/kfx/cmd/kfx/ui"
	"github.com/CiscoAI/kfx/cmd/kfx/version"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
kfx is a productivity tool for Kubeflow on-premise.
	`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return runE(flags, cmd, args)
		},
		SilenceUsage: true,
		Version:      version.Version,
	}
	cmd.AddCommand(delete.NewCommand())
	cmd.AddCommand(install.NewCommand())
	cmd.AddCommand(ui.NewCommand())
	cmd.AddCommand(check.NewCommand())
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
		ForceColors:     true,
	})
	if err := Run(); err != nil {
		os.Exit(1)
	}
}
