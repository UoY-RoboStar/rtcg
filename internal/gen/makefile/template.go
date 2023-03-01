package makefile

import (
	"embed"
	"fmt"
	"text/template"
)

//go:embed embed/templates/Makefile.tmpl
var templates embed.FS

// parseTemplate parses the Makefile template.
func parseTemplate() (*template.Template, error) {
	// TODO: possibly have different Makefiles per variant.
	tmpl, err := template.ParseFS(templates, "embed/templates/Makefile.tmpl")
	if err != nil {
		return nil, fmt.Errorf("couldn't open base Makefile template: %w", err)
	}

	return tmpl, nil
}
