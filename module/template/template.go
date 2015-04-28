package template

import (
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	t *template.Template
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

	t = template.Must(template.ParseFiles(templates...))
	return nil
}

func Parse(filename string) (*template.Template, error) {
	return t.Parse(filename)
}

func Render(w io.Writer, filename string, obj interface{}) error {
	err := t.ExecuteTemplate(w, filename, obj)
	return err
}
