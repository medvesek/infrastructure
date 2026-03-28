package utils

import (
	"os"
	"os/exec"
	"strings"
)

func Cmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func CreateTempDir(pattern string) (string, func(), error) {
	dir, err := os.MkdirTemp("", pattern)
	if err != nil {
		return dir, func() {}, err
	}

	cleanup := func() {
		os.RemoveAll(dir)
	}

	return dir, cleanup, err
}

func NameFromDomain(domain string) string {
	return strings.ReplaceAll(domain, ".", "_")
}
