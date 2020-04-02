package app

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for version
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "app",
		Short: "scaffold a new app for Cisco KF pack",
		Long:  "scaffold a new app for Cisco KF pack",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Info("`create app` invoked")
			return nil
		},
	}
	return cmd
}
