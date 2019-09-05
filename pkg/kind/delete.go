package kind

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/kind/pkg/cluster"
)

// DeleteKindCluster deletes any KinD cluster with the cluster name
func DeleteKindCluster(clusterName string) error {
	// Check if the cluster name already exists
	known, err := cluster.IsKnown(clusterName)
	if err != nil {
		return err
	}
	if !known {
		log.Printf("Cluster %v not found!", clusterName)
		return errors.New("cluster could not be deleted")
	}
	clusterContext := cluster.NewContext(clusterName)
	err = clusterContext.Delete()
	if err != nil {
		return err
	}
	log.Printf("Cluster %v deleted successfully.", clusterName)
	return nil
}
