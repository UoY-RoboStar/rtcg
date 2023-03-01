// Package makefile contains the Makefile generator.
package makefile

import (
	"github.com/UoY-RoboStar/rtcg/internal/gen/gencommon"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
	"path/filepath"
	"text/template"
)

// Generator is a generator for Makefiles.
type Generator struct {
	template  *template.Template // template is the loaded Makefile template.
	outputDir string             // outputDir is the output directory for the Makefile.
}

func (g *Generator) Generate(suite stm.Suite) error {
	outPath := filepath.Join(g.outputDir, "Makefile")

	return gencommon.ExecuteTemplateOnFile(outPath, "Makefile.tmpl", g.template, suite)
}

// New constructs a new Makefile generator.
func New(_ Config, outputDir string) (*Generator, error) {
	tmpl, err := parseTemplate()
	if err != nil {
		return nil, err
	}

	return &Generator{outputDir: outputDir, template: tmpl}, nil
}
