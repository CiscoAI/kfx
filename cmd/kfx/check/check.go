package check

import (
	"fmt"
	"os"

	"github.com/CiscoAI/kfx/pkg/healthcheck"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type checkFlags struct {
	pre bool
}

// NewCommand returns a new cobra.Command for version
func NewCommand() *cobra.Command {
	flags := &checkFlags{}
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "check",
		Short: "Checks to ensure a k8s cluster is KF ready.",
		Long:  "Checks to ensure a k8s cluster is KF ready.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if flags.pre {
				return preflightChecks(cmd, args)
			}
			return postInstallChecks(cmd, args)
		},
	}
	cmd.Flags().BoolVar(&flags.pre, "pre", false, "Run pre-installation checks, to determine if the control plane can be installed.")
	return cmd
}

func preflightChecks(cmd *cobra.Command, args []string) error {
	log.Info("Running preflight checks")
	kubeconfig := os.Getenv("KUBECONFIG")
	versionCheck, err := healthcheck.CheckK8sVersion(kubeconfig)
	if err != nil {
		return err
	}
	if !versionCheck {
		return fmt.Errorf("Kubernetes Version mis-match, please install the right k8s version")
	}
	return nil
}

func postInstallChecks(cmd *cobra.Command, args []string) error {
	log.Info("Running post install checks")
	return nil
}
