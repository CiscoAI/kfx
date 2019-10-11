package bundle

import (
	"errors"

	"github.com/CiscoAI/kfx/pkg/bootstrap"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var pipelineURI string
var appName string

// NewCommand returns a new cobra.Command for version
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "bundle",
		Short: "Bootstraps a ML app bundle",
		Long:  "Bootstraps a ML app bundle",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("`kfx create bundle` called...")
			// Bootstrap ML project repo
			if pipelineURI == "" {
				isDone := bootstrap.CreateProjectStructure(appName)
				if !isDone {
					err := errors.New("bootstrap failed")
					log.Errorf("Error bootstrapping project: %v", err)
					return err
				}
			}
			log.Infof("Created a shell project structure for ML app under %v/", appName)
			return nil
		},
	}
	cmd.Flags().StringVar(&appName, "name", "", "Name for the ML sapplication")
	cmd.Flags().StringVar(&pipelineURI, "pipeline", "", "URI for fetching the Pipeline")
	return cmd
}
