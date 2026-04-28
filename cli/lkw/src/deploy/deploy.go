package deploy

import (
	"embed"
	"fmt"

	"github.com/medvesek/infrastructure/lkw/lib/efs"
	"github.com/medvesek/infrastructure/lkw/lib/temp"
	"github.com/medvesek/infrastructure/lkw/lib/template"
	"github.com/medvesek/infrastructure/lkw/src/constants"
	"github.com/spf13/viper"
)

type DeployCmd[T any] interface {
	Run(config T) error
	Remove(config T)
}

type TemplateItem struct {
	String   string
	Data     any
	FileName string
}

func Run[T any](deployCmd DeployCmd[T]) error {
	var config T
	viper.Unmarshal(&config)
	return deployCmd.Run(config)
}

func Remove[T any](deployCmd DeployCmd[T]) {
	var config T
	viper.Unmarshal(&config)
	deployCmd.Remove(config)
}

func PrepareSupportFiles(templates []TemplateItem, embedFs embed.FS) (string, func(), error) {
	tempDir, cleanup, err := temp.CreateDir(fmt.Sprintf("%s-*", constants.CliName))
	if err != nil {
		return "", cleanup, err
	}

	for _, templateItem := range templates {
		err = template.RenderStringToFile(templateItem.String, fmt.Sprintf("%s/%s", tempDir, templateItem.FileName), templateItem.Data)
		if err != nil {
			return tempDir, cleanup, err
		}
	}

	err = efs.WriteFiles(&embedFs, tempDir)

	if err != nil {
		return tempDir, cleanup, err
	}

	return tempDir, cleanup, nil
}
