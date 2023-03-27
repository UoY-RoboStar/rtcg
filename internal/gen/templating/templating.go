// Package templating assists in template based code generation.
package templating

import (
	"errors"
	"fmt"
	"os"
	"text/template"
)

// CreateFile creates path, executes tmpl with dot ctx, and handles closing.
func CreateFile(path, tmplName string, tmpl *template.Template, ctx any) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("couldn't create output file: %w", err)
	}

	err = tmpl.ExecuteTemplate(file, tmplName, ctx)

	return errors.Join(err, file.Close())
}
