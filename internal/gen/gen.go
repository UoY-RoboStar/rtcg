// Package gen concerns the generation (template expansion) part of rtcg.
package gen

import (
	"embed"
	"errors"
	"fmt"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"
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
	g := Generator{OutputDir: outDir}

	base, err := template.ParseFS(baseTemplates, "embed/templates/*.c.tmpl")
	if err != nil {
		return nil, fmt.Errorf("couldn't open base templates: %w", err)
	}

	g.Template, err = base.ParseFS(inFS, "*.c.tmpl")

	return &g, err
}

func (g *Generator) Generate(stms stm.Suite) error {
	if err := os.MkdirAll(g.OutputDir, 0754); err != nil {
		return fmt.Errorf("couldn't make test directory: %w", err)
	}

	for k, v := range stms {
		if err := g.generateStm(k, v); err != nil {
			return fmt.Errorf("couldn't generate test %s: %w", k, err)
		}
	}
	return nil
}

func (g *Generator) generateStm(name string, s *stm.Stm) error {
	path := filepath.Join(g.OutputDir, name+".c")
	w, err := os.Create(path)
	if err != nil {
		return err
	}
	err = g.Template.ExecuteTemplate(w, "top.c.tmpl", s)
	return errors.Join(err, w.Close())
}
