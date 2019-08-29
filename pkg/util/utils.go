package util

import (
	"log"
	"os"
	"os/exec"
)

// RunShellCommands is a utility function to run Kubectl Apply
func RunShellCommands(filepath string) {
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
	}
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
