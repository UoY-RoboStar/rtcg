// Package cpp contains the C++ code generator.
package cpp

import (
	"fmt"
	"path/filepath"

	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

const (
	srcDir = "src" // srcDir is the subdirectory of the output directory for source code.
)

// Generator is a C++ code generator.
type Generator struct {
	config Config // config is the configuration for this Generator.

	testFiles []TestFile  // testFiles is the list of files this generator will make.
	templates TemplateSet // templates is the map of templates to use for files in testFiles.

	inputDir   string // inputDir is the output directory for this Generator.
	srcBaseDir string // srcBaseDir is the output source directory for this Generator.
	outputDir  string // outputDir is the output directory for this Generator.
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

// New constructs a new C++ code generator from config, rooted at outputDir.
func New(config Config, inputDir, outputDir string) (*Generator, error) {
	var (
		gen Generator
		err error
	)

	gen.config = config
	gen.inputDir = config.Variant.Dir(inputDir)
	gen.outputDir = config.Variant.Dir(outputDir)
	gen.srcBaseDir = filepath.Join(gen.outputDir, srcDir)

	gen.testFiles = []TestFile{
		{Dir: "src", Name: "main.cpp", Desc: "main C++ file", Glob: "*.cpp.tmpl"},
		{Dir: "include", Name: "convert.h", Desc: "type convert header", Glob: "convert/*.h.tmpl"},
	}

	if gen.templates, err = NewTemplateSet(config.Variant, gen.testFiles); err != nil {
		return nil, err
	}

	return &gen, nil
}

// TestFile holds information about a test file.
type TestFile struct {
	Dir  string // Dir is the destination directory of this file.
	Name string // Name is the filename of this file.
	Desc string // Desc is a human-readable description for this file.
	Glob string // Glob is the slash-delimited glob of source templates for this file.
}
