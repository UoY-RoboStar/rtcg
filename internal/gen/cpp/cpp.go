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
	baseDir := g.srcBaseDir

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
		{Name: "main.cpp", Desc: "main C++ file", SrcGlob: "*.cpp.tmpl"},
		{Name: "convert.h", Desc: "type conversion header file", SrcGlob: "convert/*.h.tmpl"},
	}

	if gen.templates, err = NewTemplateSet(config.Variant, gen.testFiles); err != nil {
		return nil, err
	}

	return &gen, nil
}

// TestFile holds information about a test file.
type TestFile struct {
	Name    string // Name is the filename of this file.
	Desc    string // Desc is a human-readable description for this file.
	SrcGlob string // SrcGlob is the slash-delimited glob of source templates for this file.
}
