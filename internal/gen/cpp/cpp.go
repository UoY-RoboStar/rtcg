// Package cpp contains the C++ code generator.
package cpp

import (
	"fmt"
	"github.com/UoY-RoboStar/rtcg/internal/gen/gencommon"
	"github.com/UoY-RoboStar/rtcg/internal/gen/makefile"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
	"path/filepath"
	"text/template"
)

const (
	srcDir = "src" // srcDir is the subdirectory of the output directory for source code.
)

// Generator is a C++ code generator.
type Generator struct {
	config    Config              // config is the configuration for this Generator.
	template  *template.Template  // template is the template to use for C++ files.
	makefile  *makefile.Generator // makefile is the optional generator for Makefiles.
	outputDir string              // outputDir is the output directory for this Generator.
}

func (g *Generator) Name() string {
	return fmt.Sprintf("C++ (%s)", g.config.Variant)
}

func (g *Generator) Dirs(suite stm.Suite) []string {
	baseDir := filepath.Join(g.outputDir, srcDir)

	dirs := make([]string, 1, len(suite)+1)
	dirs[0] = filepath.Join(baseDir, preludeDir)

	for name := range suite {
		dirs = append(dirs, filepath.Join(baseDir, name))
	}

	return dirs
}

func (g *Generator) Generate(suite stm.Suite) error {
	if err := g.copyPrelude(); err != nil {
		return err
	}

	if err := g.generateMakefile(suite); err != nil {
		return err
	}

	if err := g.generateSuite(suite); err != nil {
		return err
	}

	return nil
}

func (g *Generator) generateMakefile(suite stm.Suite) error {
	if g.makefile == nil {
		return nil
	}

	return g.makefile.Generate(suite)
}

func (g *Generator) generateSuite(suite stm.Suite) error {
	for k, v := range suite {
		if err := g.generateStm(k, v); err != nil {
			return fmt.Errorf("couldn't generate test %s: %w", k, err)
		}
	}

	return nil
}

func (g *Generator) generateStm(name string, body *stm.Stm) error {
	outPath := filepath.Join(g.outputDir, srcDir, name, name+".cpp")

	ctx := NewContext(name, body, g.config)

	return gencommon.ExecuteTemplateOnFile(outPath, "top.cpp.tmpl", g.template, ctx)
}

// New constructs a new C++ code generator from config, rooted at outputDir.
func New(config Config, outputDir string) (*Generator, error) {
	var (
		gen Generator
		err error
	)

	gen.config = config
	gen.outputDir = filepath.Join(outputDir, config.Variant.String())

	if gen.template, err = parseTemplate(config.Variant); err != nil {
		return nil, err
	}

	if gen.config.Makefile != nil {
		if gen.makefile, err = makefile.New(*gen.config.Makefile, gen.outputDir); err != nil {
			return nil, fmt.Errorf("couldn't construct Makefile generator: %w", err)
		}
	}

	return &gen, nil
}
