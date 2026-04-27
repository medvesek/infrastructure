package temp

import (
	"os"
)

func CreateDir(pattern string) (string, func(), error) {
	dir, err := os.MkdirTemp("", pattern)
	if err != nil {
		return dir, func() {}, err
	}

	cleanup := func() {
		os.RemoveAll(dir)
	}

	return dir, cleanup, err
}
