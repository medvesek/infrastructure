package deploy

import (
	"embed"
	"fmt"

	"github.com/medvesek/infrastructure/lkw/assets"
	"github.com/medvesek/infrastructure/lkw/constants"
	"github.com/medvesek/infrastructure/lkw/lib/template"
	"github.com/medvesek/infrastructure/lkw/lib/utils"
)

type TemplateItem struct {
	String   string
	Data     any
	FileName string
}

func PrepareSupportFiles(templates []TemplateItem, embedFs embed.FS) (string, func(), error) {
	tempDir, cleanup, err := utils.CreateTempDir(fmt.Sprintf("%s-*", constants.AppName))
	if err != nil {
		return "", cleanup, err
	}

	for _, templateItem := range templates {
		err = template.RenderStringToFile(templateItem.String, fmt.Sprintf("%s/%s", tempDir, templateItem.FileName), templateItem.Data)
		if err != nil {
			return tempDir, cleanup, err
		}
	}

	err = assets.WriteFiles(&embedFs, tempDir)

	if err != nil {
		return tempDir, cleanup, err
	}

	return tempDir, cleanup, nil
}
