package bootstrap

import (
	"github.com/CiscoAI/create-kf-app/pkg/util"
)

// Inspired by Tilt and kubefwd's port forwarding implementation
// Huge shoutout to their awesome work
// Tilt: https://github.com/windmilleng/tilt/
// kubefwd: https://github.com/txn2/kubefwd/
// Skaffold

// MLAPortForwardShell shells out to kubectl to do the port-forward into the MLA app
func MLAPortForwardShell() error {
	err := util.KubectlPortForward("mlanywhere", "kubeflow", "5000", "5000")
	if err != nil {
		return err
	}
	return nil
}

// MLAPortForward uses the client-go tools to port forward to the app service
// func MLAPortForward() error {
// 	pf, err := portforward.NewPortForwarder("default", metav1.LabelSelector{
// 		MatchLabels: map[string]string{
// 			"app": "nginx",
// 		},
// 	}, 80)
// 	if err != nil {
// 		log.Fatal("Error setting up port forwarder: ", err)
// 	}

// 	err = pf.Start()
// 	if err != nil {
// 		log.Fatal("Error starting port forward: ", err)
// 	}

// 	log.Printf("Started tunnel on %d\n", pf.ListenPort)

// 	return nil
// }
