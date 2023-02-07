// Package gen contains the test code generation logic.
package gen

import (
	"errors"
	"fmt"
	"rtcg/internal/cli"
	"rtcg/internal/testlang"
)

// Generator is the main engine for generating test code.
type Generator struct {
	TemplateDir string // TemplateDir is the path to the template directory.
	InputFile   string // InputFile is the path to the input file.
}

func (g *Generator) Run() error {
	f, err := cli.OpenFileOrStdin(g.InputFile)
	if err != nil {
		return fmt.Errorf("couldn't open %s: %w", g.inputFileName(), err)
	}

	_, err = testlang.ReadSuite(f)
	return errors.Join(err, f.Close())
}

func (g *Generator) isStdin() bool {
	// This may become something more sophisticated eventually.
	return g.InputFile == "-"
}

func (g *Generator) inputFileName() any {
	if g.isStdin() {
		return "standard input"
	}
	return g.InputFile
}
