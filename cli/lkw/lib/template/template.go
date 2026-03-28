package template

import (
	"bytes"
	"os"
	"text/template"
)

func RenderStringToString(templateString string, data any) (string, error) {
	template, err := prepare(templateString)

	if err != nil {
		return "", err
	}

	var buff bytes.Buffer

	err = template.Execute(&buff, data)

	if err != nil {
		return "", err
	}

	return buff.String(), nil
}

func RenderStringToFile(templateString string, destinationPath string, data any) error {
	template, err := prepare(templateString)

	if err != nil {
		return err
	}

	f, err := os.Create(destinationPath)

	if err != nil {
		return err
	}

	defer f.Close()

	return template.Execute(f, data)
}

func RenderFileToString(filePath string, data any) (string, error) {
	fileContents, err := os.ReadFile(filePath)

	if err != nil {
		return "", err
	}

	return RenderStringToString(string(fileContents), data)
}

func RenderFileToFile(sourcePath string, destinationPath string, data any) error {
	fileContents, err := os.ReadFile(sourcePath)

	if err != nil {
		return err
	}

	return RenderStringToFile(string(fileContents), destinationPath, data)
}

func prepare(templateString string) (*template.Template, error) {
	tmpl := template.New("t")
	return tmpl.Parse(templateString)
}
