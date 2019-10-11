package build

import (
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for version
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "build",
		Short: "builds the kf-app",
		Long:  "builds the kf-app",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return cmd
}
