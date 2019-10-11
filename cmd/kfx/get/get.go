package get

import (
	"github.com/CiscoAI/kfx/cmd/kfx/get/cluster"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for version
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "get",
		Short: "get the cluster / kf-app",
		Long:  "get the cluster / kf-app",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	cmd.AddCommand(cluster.NewCommand())
	return cmd
}
