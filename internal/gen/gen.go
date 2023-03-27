// Package gen concerns the generation (template expansion) part of rtcg.
package gen

import (
	"fmt"
	"os"

	"github.com/UoY-RoboStar/rtcg/internal/gen/subgen"

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
	subgens []subgen.Subgenerator // subgens is the set of configured sub-generators.
	dirs    gencommon.DirSet      // dirs is the directory set of the generator.
}

func initCpp(cfgs []cfg.Config, dirs gencommon.DirSet) ([]subgen.Subgenerator, error) {
	var (
		subs, varSubs []subgen.Subgenerator
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

type genFunc func(*cfg.Config, gencommon.DirSet) ([]subgen.Subgenerator, error)

func initCppVariant(config *cfg.Config, dirs gencommon.DirSet) ([]subgen.Subgenerator, error) {
	var (
		varSubs, stepSubs []subgen.Subgenerator
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

func initCppMain(config *cfg.Config, dirs gencommon.DirSet) ([]subgen.Subgenerator, error) {
	gen, err := cpp.New(config, dirs)
	if err != nil {
		return nil, fmt.Errorf("couldn't init c++ generator: %w", err)
	}

	return []subgen.Subgenerator{gen}, nil
}

func initCppCatkin(config *cfg.Config, dirs gencommon.DirSet) ([]subgen.Subgenerator, error) {
	if config.Catkin == nil {
		return nil, nil
	}

	gen, err := catkin.New(config.Catkin, dirs)
	if err != nil {
		return nil, fmt.Errorf("couldn't init Catkin generator: %w", err)
	}

	return []subgen.Subgenerator{gen}, nil
}

func initCppMakefile(config *cfg.Config, dirs gencommon.DirSet) ([]subgen.Subgenerator, error) {
	if config.Makefile == nil {
		return nil, nil
	}

	gen, err := makefile.New(config, dirs)
	if err != nil {
		return nil, fmt.Errorf("couldn't init Makefile generator: %w", err)
	}

	return []subgen.Subgenerator{gen}, nil
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

	return &Generator{subgens: subs, dirs: dirs}, nil
}

func (g *Generator) OnSuite(suite *stm.Suite) *OnSuite {
	// TODO(@MattWindsor91): possibly make this return a subgen.OnSuite

	subgens := make([]subgen.OnSuite, len(g.subgens))

	for i, subg := range g.subgens {
		subgens[i] = subg.OnSuite(suite)
	}

	return &OnSuite{parent: g, subgens: subgens}
}

type OnSuite struct {
	parent  *Generator
	subgens []subgen.OnSuite
}

func (o *OnSuite) Generate() error {
	if err := o.mkdirs(); err != nil {
		return err
	}

	if err := o.runSubgens(); err != nil {
		return err
	}

	return nil
}

// mkdirs makes the various directories used by the generator.
func (o *OnSuite) mkdirs() error {
	for _, subg := range o.subgens {
		name := subg.Parent().Name()

		for _, dir := range subg.Dirs() {
			if err := os.MkdirAll(dir, outputDirPerms); err != nil {
				return fmt.Errorf("couldn't make %s directory at %q: %w", name, dir, err)
			}
		}
	}

	return nil
}

func (o *OnSuite) runSubgens() error {
	for _, subg := range o.subgens {
		name := subg.Parent().Name()

		if err := subg.Generate(); err != nil {
			return fmt.Errorf("generator %s failed: %w", name, err)
		}
	}

	return nil
}
