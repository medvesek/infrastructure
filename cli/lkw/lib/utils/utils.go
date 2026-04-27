package utils

import (
	"strings"
)

func DockerName(name string) string {
	return strings.ReplaceAll(name, ".", "_")
}

func EnsureTrailingSlash(path string) string {
	if !strings.HasSuffix(path, "/") {
		return path + "/"
	}
	return path
}
