package assets

import (
	"embed"
)

//go:embed static/templates/docker-compose.yaml.template
var StaticDockerComposeTemplate string

type StaticDockerComposeData struct {
	Name   string
	Domain string
}

//go:embed static/templates/default.conf.template
var StaticNginxDefaultConfTemplate string

type StaticNginxDefaultConfData struct {
	Spa bool
}

//go:embed static/files
var StaticFiles embed.FS

var NoFiles embed.FS

//go:embed app/templates/override.yml.template
var AppDockerComposeOverrideTemplate string

type AppDockerComposeOverrideData struct {
	Name   string
	Expose []struct {
		Service string
		Expose  []struct {
			Domain     string
			Port       int
			Entrypoint string
		}
	}
}
