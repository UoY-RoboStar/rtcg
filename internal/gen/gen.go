// Package gen concerns the generation (template expansion) part of rtcg.
package gen

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"

	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

//go:embed embed/templates/*.c.tmpl
var baseTemplates embed.FS

// Generator is a test code generator.
type Generator struct {
	Template  *template.Template // Template is the template to use.
	OutputDir string             // OutputDir is the output directory for the tests.
}

// New creates a new Generator by reading all templates from inFS, and outputting to outDir.
func New(inFS fs.FS, outDir string) (*Generator, error) {
	generator := Generator{OutputDir: outDir, Template: nil}

	base, err := template.ParseFS(baseTemplates, "embed/templates/*.c.tmpl")
	if err != nil {
		return nil, fmt.Errorf("couldn't open base templates: %w", err)
	}

	generator.Template, err = base.ParseFS(inFS, "*.c.tmpl")
	if err != nil {
		return nil, fmt.Errorf("couldn't open user templates: %w", err)
	}

	return &generator, nil
}

// outputDirPerms is the permissions to use when generating the output directory.
const outputDirPerms = 0o754

func (g *Generator) Generate(suite stm.Suite) error {
	if err := os.MkdirAll(g.OutputDir, outputDirPerms); err != nil {
		return fmt.Errorf("couldn't make test directory: %w", err)
	}

	for k, v := range suite {
		if err := g.generateStm(k, v); err != nil {
			return fmt.Errorf("couldn't generate test %s: %w", k, err)
		}
	}

	return nil
}

func (g *Generator) generateStm(name string, body *stm.Stm) error {
	path := filepath.Join(g.OutputDir, name+".c")

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("couldn't create output file for test %q: %w", name, err)
	}

	err = g.Template.ExecuteTemplate(file, "top.c.tmpl", body)

	return errors.Join(err, file.Close())
}
