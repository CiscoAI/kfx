package ui

import (
	"github.com/CiscoAI/create-kf-app/kfx/cmd/ui/mla"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for ui
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "ui",
		Short: "Connects to the UI",
		Long:  "Connects to the UI.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	cmd.AddCommand(mla.NewCommand())
	return cmd
}
