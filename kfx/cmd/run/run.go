package run

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version is the create-kf-app CLI version
const Version = "v0.1-alpha"

// NewCommand returns a new cobra.Command for version
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "version",
		Short: "prints the create-kf-app CLI version",
		Long:  "prints the create-kf-app CLI version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(Version)
			return nil
		},
	}
	return cmd
}
