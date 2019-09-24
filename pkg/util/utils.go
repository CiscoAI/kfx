package util

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// RunShellCommands is a utility function to run Kubectl Apply
func RunShellCommands(filepath string) error {
	command := "bash"
	args := []string{}
	args = append(args, "kubectl apply -f")
	args = append(args, filepath)
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("Shell command %s %v failed with %s\n", command, args, err)
		return err
	}
	return nil
}

// KubectlPortForward shelled out kubectl port-forward
func KubectlPortForward(serviceName string, namespace string, localPort string, remotePort string) error {
	command := "bash"
	args := []string{}
	args = append(args, "-c", fmt.Sprintf("kubectl port-forward svc/%s -n %s %s:%s --pod-running-timeout=1s", serviceName, namespace, localPort, remotePort))
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("Shell command %s %v failed with %s\n", command, args, err)
		return err
	}
	return nil
}

// RunShellScript is a utility fucntion used to run a shell script file
func RunShellScript(filepath string) {
	cmd := exec.Command("sh", filepath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("Shell command failed with %s\n", err)
	}
}
