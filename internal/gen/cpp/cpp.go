// Package cpp contains the C++ code generator.
package cpp

import (
	"fmt"
	"path/filepath"

	"github.com/UoY-RoboStar/rtcg/internal/gen/gencommon"

	cfg "github.com/UoY-RoboStar/rtcg/internal/gen/config/cpp"

	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

const (
	srcDir = "src" // srcDir is the subdirectory of the output directory for source code.
)

// Generator is a C++ code generator.
type Generator struct {
	config cfg.Config // config is the configuration for this Generator.

	inputDir   string // inputDir is the output directory for this Generator.
	srcBaseDir string // srcBaseDir is the output source directory for this Generator.
	outputDir  string // outputDir is the output directory for this Generator.

	gencommon.TemplatedGenerator
}

func (g *Generator) Name() string {
	return fmt.Sprintf("C++ (%s)", g.config.Variant)
}

func (g *Generator) Dirs(suite stm.Suite) []string {
	dirs := make([]string, 1, len(suite)+1)
	dirs[0] = filepath.Join(g.srcBaseDir, preludeDir)

	for name := range suite {
		dirs = append(dirs, g.testDirs(name)...)
	}

	return dirs
}

func (g *Generator) testDirs(name string) []string {
	baseDir := filepath.Join(g.srcBaseDir, name)

	// This directory structure mirrors that of catkin, even if we're not generating ROS.
	return []string{
		filepath.Join(baseDir, "src"),
		filepath.Join(baseDir, "include"),
	}
}

func (g *Generator) Generate(suite stm.Suite) error {
	if err := g.copyPrelude(); err != nil {
		return err
	}

	if err := g.generateTests(suite); err != nil {
		return err
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
	gen.inputDir = config.Variant.Dir(dirs.Input)
	gen.outputDir = config.Variant.Dir(dirs.Output)
	gen.srcBaseDir = filepath.Join(gen.outputDir, srcDir)

	if gen.TemplatedGenerator, err = NewTemplatedGenerator(config); err != nil {
		return nil, err
	}

	return &gen, nil
}
