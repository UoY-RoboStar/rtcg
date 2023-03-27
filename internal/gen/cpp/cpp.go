// Package cpp contains the C++ code generator.
package cpp

import (
	"fmt"

	"github.com/UoY-RoboStar/rtcg/internal/gen/templating"

	cfg "github.com/UoY-RoboStar/rtcg/internal/gen/config/cpp"
	"github.com/UoY-RoboStar/rtcg/internal/gen/gencommon"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

// Generator is a C++ code generator.
type Generator struct {
	config    cfg.Config       // config is the configuration for this Generator.
	srcDirSet gencommon.DirSet // srcDirSet is the source directory set for this Generator.

	testGen *templating.Generator // testGen is the template-based generator for test-based files.
}

func (g *Generator) Name() string {
	return fmt.Sprintf("C++ (%s)", g.config.Variant)
}

func (g *Generator) Dirs(suite stm.Suite) []string {
	dirs := make([]string, 2, len(suite)+2)
	dirs[0] = g.srcDirSet.OutputPath(preludeDir)
	dirs[1] = g.srcDirSet.OutputPath(convertDir)

	for name := range suite {
		dirs = append(dirs, g.testDirs(name)...)
	}

	return dirs
}

func (g *Generator) testDirs(name string) []string {
	dirSet := g.srcDirSet.Subdir(name)

	// This directory structure mirrors that of catkin, even if we're not generating ROS.
	return []string{
		dirSet.OutputPath("include"),
		dirSet.OutputPath("src"),
	}
}

func (g *Generator) Generate(suite stm.Suite) error {
	if err := g.copyPrelude(); err != nil {
		return err
	}

	if err := gencommon.GenerateTests(g.srcDirSet, suite, g); err != nil {
		return fmt.Errorf("couldn't generate for tests: %w", err)
	}

	return nil
}

// New constructs a new C++ code generator from config, rooted at the given directories.
func New(config *cfg.Config, dirs gencommon.DirSet) (*Generator, error) {
	var (
		gen Generator
		err error
	)

	gen.config = *config
	gen.srcDirSet = dirs.Subdir("src")

	if gen.testGen, err = NewTemplatedGenerator(config); err != nil {
		return nil, err
	}

	return &gen, nil
}

func (g *Generator) GenerateTest(dirs gencommon.DirSet, name string, test *stm.Stm) error {
	ctx := NewContext(name, test, g.config)

	// TODO(@MattWindsor91): only do this once
	if err := g.copyConvertFile(ctx); err != nil {
		return err
	}

	return g.testGen.Generate(dirs.Output, ctx)
}

// copyConvertFile copies convert.cpp from the input directory, if there is one.
func (g *Generator) copyConvertFile(ctx *Context) error {
	if !ctx.HasConversion {
		return nil
	}

	return copyLocalFile(g.srcDirSet.Subdir(convertDir), "convert.cpp")
}
