package assets

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

//go:embed static/templates/docker-compose.yaml.template
var DockerComposeTemplate string

type DockerComposeData struct {
	Name   string
	Domain string
}

//go:embed static/templates/default.conf.template
var NginxDefaultConf string

type NginxDefaultConfData struct {
	Type string
}

//go:embed static/files
var StaticFiles embed.FS

func GetAllFiles(efs *embed.FS) ([]string, error) {
	var files []string
	err := fs.WalkDir(efs, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		files = append(files, path)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func WriteFiles(efs *embed.FS, destination string) error {
	files, err := GetAllFiles(efs)

	if err != nil {
		return err
	}

	for _, file := range files {
		content, err := efs.ReadFile(file)
		if err != nil {
			return err
		}
		parts := strings.Split(file, "/")
		fileName := parts[len(parts)-1]
		err = os.WriteFile(fmt.Sprintf("%s/%s", destination, fileName), content, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
