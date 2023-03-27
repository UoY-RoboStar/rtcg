// Package gen concerns the generation (template expansion) part of rtcg.
package gen

import (
	"fmt"
	"os"

	"github.com/UoY-RoboStar/rtcg/internal/gen/catkin"
	"github.com/UoY-RoboStar/rtcg/internal/gen/config"
	cfg "github.com/UoY-RoboStar/rtcg/internal/gen/config/cpp"
	"github.com/UoY-RoboStar/rtcg/internal/gen/cpp"
	"github.com/UoY-RoboStar/rtcg/internal/gen/gencommon"
	"github.com/UoY-RoboStar/rtcg/internal/gen/makefile"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

const (
	outputDirPerms = 0o754 // outputDirPerms is the output directory permission mask.
)

// Generator is a test code generator.
type Generator struct {
	subgenerators []Subgenerator   // subgenerators is the set of configured sub-generators.
	dirs          gencommon.DirSet // dirs is the directory set of the generator.
}

func initCpp(cfgs []cfg.Config, dirs gencommon.DirSet) ([]Subgenerator, error) {
	var (
		subs, varSubs []Subgenerator
		err           error
	)

	for i := range cfgs {
		if varSubs, err = initCppVariant(&cfgs[i], dirs.Subdir(cfgs[i].Variant.String())); err != nil {
			return nil, err
		}

		subs = append(subs, varSubs...)
	}

	return subs, nil
}

type genFunc func(*cfg.Config, gencommon.DirSet) ([]Subgenerator, error)

func initCppVariant(config *cfg.Config, dirs gencommon.DirSet) ([]Subgenerator, error) {
	var (
		varSubs, stepSubs []Subgenerator
		err               error
	)

	for _, fun := range []genFunc{
		initCppMain,
		initCppCatkin,
		initCppMakefile,
	} {
		if stepSubs, err = fun(config, dirs); err != nil {
			return nil, err
		}

		varSubs = append(varSubs, stepSubs...)
	}

	return varSubs, nil
}

func initCppMain(config *cfg.Config, dirs gencommon.DirSet) ([]Subgenerator, error) {
	gen, err := cpp.New(config, dirs)
	if err != nil {
		return nil, fmt.Errorf("couldn't init c++ generator: %w", err)
	}

	return []Subgenerator{gen}, nil
}

func initCppCatkin(config *cfg.Config, dirs gencommon.DirSet) ([]Subgenerator, error) {
	if config.Catkin == nil {
		return nil, nil
	}

	gen, err := catkin.New(config.Catkin, dirs)
	if err != nil {
		return nil, fmt.Errorf("couldn't init Catkin generator: %w", err)
	}

	return []Subgenerator{gen}, nil
}

func initCppMakefile(config *cfg.Config, dirs gencommon.DirSet) ([]Subgenerator, error) {
	if config.Makefile == nil {
		return nil, nil
	}

	gen, err := makefile.New(config, dirs)
	if err != nil {
		return nil, fmt.Errorf("couldn't init Makefile generator: %w", err)
	}

	return []Subgenerator{gen}, nil
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
func New(cfg config.Config, outputDir string) (*Generator, error) {
	dirs := gencommon.DirSet{
		Input:  cfg.Directory,
		Output: outputDir,
	}

	subs, err := initCpp(cfg.Cpps, dirs)
	if err != nil {
		return nil, err
	}

	return &Generator{subgenerators: subs, dirs: dirs}, nil
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
