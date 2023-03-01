package gencommon

import (
	"errors"
	"fmt"
	"os"
	"text/template"
)

// ExecuteTemplateOnFile creates path, executes tmpl with dot ctx, and handles closure.
func ExecuteTemplateOnFile(path, tmplName string, tmpl *template.Template, ctx any) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("couldn't create output file: %w", err)
	}

	err = tmpl.ExecuteTemplate(file, tmplName, ctx)

	return errors.Join(err, file.Close())
}
