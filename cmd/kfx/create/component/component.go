package component

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for version
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "components",
		Short: "scaffold a new component for Cisco KF app",
		Long:  "scaffold a new component for Cisco KF app",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Info("`create component` invoked")
			return nil
		},
	}
	return cmd
}
