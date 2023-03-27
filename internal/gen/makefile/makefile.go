// Package makefile contains the Makefile generator.
package makefile

import (
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/UoY-RoboStar/rtcg/internal/gen/config/cpp"
	"github.com/UoY-RoboStar/rtcg/internal/gen/gencommon"
	"github.com/UoY-RoboStar/rtcg/internal/gen/subgen"
	"github.com/UoY-RoboStar/rtcg/internal/gen/templating"
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

func (g *Generator) OnSuite(suite *stm.Suite) subgen.OnSuite {
	return &OnSuite{
		parent: g,
		suite:  suite,
		cctx:   g.config.Process(suite.Types),
	}
}

type OnSuite struct {
	parent *Generator
	cctx   cpp.Context
	suite  *stm.Suite
}

func (o *OnSuite) Parent() subgen.Subgenerator {
	return o.parent
}

func (o *OnSuite) Dirs() []string {
	// Assume the C++ generator makes the appropriate directories.
	return nil
}

// Generate generates a Makefile for tests.
func (o *OnSuite) Generate() error {
	outPath := filepath.Join(o.parent.outputDir, "Makefile")

	ctx := NewContext(o.suite, o.cctx)

	err := templating.CreateFile(outPath, "Makefile.tmpl", o.parent.template, ctx)
	if err != nil {
		return fmt.Errorf("couldn't generate Makefile: %w", err)
	}

	return nil
}

// New constructs a new C++ Makefile generator rooted at outputDir.
// The Makefile will be generated in the variant subdirectory inside outputDir.
func New(cfg *cpp.Config, dirs gencommon.DirSet) (*Generator, error) {
	tmpl, err := parseTemplate()
	if err != nil {
		return nil, err
	}

	return &Generator{config: *cfg, outputDir: dirs.Output, template: tmpl}, nil
}
