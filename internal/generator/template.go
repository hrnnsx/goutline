package generator

import (
	"bytes"
	"embed"
	"io/fs"
	"text/template"
)

//go:embed templates
var templateFS embed.FS

func ReadTemplate(path string) ([]byte, error) {
	data, err := templateFS.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ListTemplates(dir string) ([]string, error) {
	var templates []string

	err := fs.WalkDir(templateFS, dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		templates = append(templates, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return templates, nil
}

func RenderTemplate(path string, data any) ([]byte, error) {
	source, err := ReadTemplate(path)
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New(path).Parse(string(source))
	if err != nil {
		return nil, err
	}

	var output bytes.Buffer
	err = tmpl.Execute(&output, data)
	if err != nil {
		return nil, err
	}

	return output.Bytes(), nil
}
