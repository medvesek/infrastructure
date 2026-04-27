package cmd

import (
	"os"
	"os/exec"
)

func Run(name string, args ...string) error {
	cmd := exec.Command(name, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Rsync(source string, destination string) error {
	return Run("rsync", "-avP", source, destination)
}

func RsyncD(source string, destination string) error {
	return Run("rsync", "-avP", source, destination, "--delete")
}
