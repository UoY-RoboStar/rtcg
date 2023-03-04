// Package gen concerns the generation (template expansion) part of rtcg.
package gen

import (
	"fmt"
	"os"

	"github.com/UoY-RoboStar/rtcg/internal/gen/cpp"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

const (
	outputDirPerms = 0o754 // outputDirPerms is the output directory permission mask.
)

// Generator is a test code generator.
type Generator struct {
	subgenerators []Subgenerator // subgenerators is the set of configured sub-generators.
	outputDir     string         // outputDir is the output directory of the generator.
}

func (g *Generator) initCpp(cfgs []cpp.Config) error {
	for _, cfg := range cfgs {
		gen, err := cpp.New(cfg, g.outputDir)
		if err != nil {
			return fmt.Errorf("couldn't init c++ generator: %w", err)
		}

		g.subgenerators = append(g.subgenerators, gen)
	}

	return nil
}

// Subgenerator captures the idea of a test code sub-generator.
type Subgenerator interface {
	// Name gets the name of this item.
	Name() string

	// Dirs gets the list of directories to make for the given test suite.
	Dirs(suite stm.Suite) []string

	// Generate generates code for suite.
	Generate(suite stm.Suite) error
}

// New creates a new Generator from config, targeting outputDir.
func New(cfg Config, outputDir string) (*Generator, error) {
	gen := Generator{
		subgenerators: make([]Subgenerator, 0, len(cfg.Cpps)),
		outputDir:     outputDir,
	}

	if err := gen.initCpp(cfg.Cpps); err != nil {
		return nil, err
	}

	return &gen, nil
}

func (g *Generator) Generate(suite stm.Suite) error {
	if err := g.mkdirs(suite); err != nil {
		return err
	}

	if err := g.runSubgens(suite); err != nil {
		return err
	}

	return nil
}

// mkdirs makes the various directories used by the generator.
func (g *Generator) mkdirs(suite stm.Suite) error {
	for _, subg := range g.subgenerators {
		name := subg.Name()

		for _, dir := range subg.Dirs(suite) {
			if err := os.MkdirAll(dir, outputDirPerms); err != nil {
				return fmt.Errorf("couldn't make %s directory at %q: %w", name, dir, err)
			}
		}
	}

	return nil
}

func (g *Generator) runSubgens(suite stm.Suite) error {
	// TODO: possibly move the sub-generator logic into here, eg by having the main generator
	// receive templates to instantiate, files to copy, etc.
	for _, subg := range g.subgenerators {
		if err := subg.Generate(suite); err != nil {
			return fmt.Errorf("generator %s failed: %w", subg.Name(), err)
		}
	}

	return nil
}
