// Package cpp contains the C++ code generator.
package cpp

import (
	"fmt"
	"path/filepath"

	"github.com/UoY-RoboStar/rtcg/internal/gen/templating"

	cfg "github.com/UoY-RoboStar/rtcg/internal/gen/config/cpp"
	"github.com/UoY-RoboStar/rtcg/internal/gen/gencommon"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

// Generator is a C++ code generator.
type Generator struct {
	config     cfg.Config       // config is the configuration for this Generator.
	dirSet     gencommon.DirSet // dirSet is the input and output directory set for this Generator.
	srcBaseDir string           // srcBaseDir is the output source directory for this Generator.

	templating.Generator
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

	if err := gencommon.GenerateTests(suite, g); err != nil {
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
	gen.dirSet = dirs
	gen.srcBaseDir = dirs.SrcDir()

	if gen.Generator, err = NewTemplatedGenerator(config); err != nil {
		return nil, err
	}

	return &gen, nil
}
