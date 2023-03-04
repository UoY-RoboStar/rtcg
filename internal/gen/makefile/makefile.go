// Package makefile contains the Makefile generator.
package makefile

import (
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/UoY-RoboStar/rtcg/internal/gen/gencommon"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

// Generator is a generator for Makefiles.
type Generator struct {
	template  *template.Template // template is the loaded Makefile template.
	outputDir string             // outputDir is the output directory for the Makefile.
}

func (g *Generator) Generate(suite stm.Suite) error {
	outPath := filepath.Join(g.outputDir, "Makefile")

	err := gencommon.ExecuteTemplateOnFile(outPath, "Makefile.tmpl", g.template, suite)
	if err != nil {
		return fmt.Errorf("couldn't generate Makefile: %w", err)
	}

	return nil
}

// New constructs a new Makefile generator.
func New(_ Config, outputDir string) (*Generator, error) {
	tmpl, err := parseTemplate()
	if err != nil {
		return nil, err
	}

	return &Generator{outputDir: outputDir, template: tmpl}, nil
}
