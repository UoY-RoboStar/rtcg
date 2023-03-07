// Package makefile contains the Makefile generator.
package makefile

import (
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/UoY-RoboStar/rtcg/internal/gen/cpp"
	"github.com/UoY-RoboStar/rtcg/internal/gen/gencommon"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

// Generator is a generator for Makefiles.
type Generator struct {
	config    cpp.Config         // config is the C++ configuration.
	template  *template.Template // template is the loaded Makefile template.
	outputDir string             // outputDir is the output directory for the Makefile.
}

func (g *Generator) Name() string {
	return "Makefile"
}

func (g *Generator) Dirs(stm.Suite) []string {
	// Assume the C++ generator makes the appropriate directories.
	return nil
}

// Generate generates a Makefile for tests.
func (g *Generator) Generate(tests stm.Suite) error {
	outPath := filepath.Join(g.outputDir, "Makefile")

	ctx, err := NewContext(tests, g.config)
	if err != nil {
		return fmt.Errorf("couldn't create Makefile context: %w", err)
	}

	err = gencommon.ExecuteTemplateOnFile(outPath, "Makefile.tmpl", g.template, ctx)
	if err != nil {
		return fmt.Errorf("couldn't generate Makefile: %w", err)
	}

	return nil
}

// New constructs a new C++ Makefile generator rooted at outputDir.
// The Makefile will be generated in the variant subdirectory inside outputDir.
func New(cfg cpp.Config, outputDir string) (*Generator, error) {
	tmpl, err := parseTemplate()
	if err != nil {
		return nil, err
	}

	return &Generator{config: cfg, outputDir: cfg.Variant.Dir(outputDir), template: tmpl}, nil
}
