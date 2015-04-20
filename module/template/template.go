package template

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var (
	Template *template.Template
)

func LoadTemplates(templatePath, format string) error {
	var templates []string

	fn := func(path string, f os.FileInfo, err error) error {
		if f.IsDir() != true && strings.HasSuffix(f.Name(), ".html") {
			templates = append(templates, path)
		}
		return nil
	}

	err := filepath.Walk(templatePath, fn)

	if err != nil {
		return err
	}

	Template = template.Must(template.ParseFiles(templates...))
	return nil
}

func Render(w http.ResponseWriter, filename string, obj interface{}) error {
	err := Template.ExecuteTemplate(w, filename, obj)
	return err
}
