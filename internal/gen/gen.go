// Package gen contains the test code generation logic.
package gen

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"test-code-gen/internal/testlang"
)

// Generator is the main engine for generating test code.
type Generator struct {
	TemplateDir string // TemplateDir is the path to the template directory.
	InputFile   string // InputFile is the path to the input file.
}

func (g *Generator) Run() error {
	f, err := os.Open(g.InputFile)
	if err != nil {
		return fmt.Errorf("couldn't open %s: %w", g.InputFile, err)
	}

	j := json.NewDecoder(f)

	var test testlang.Node

	if err = j.Decode(&test); err != nil {
		return errors.Join(err, f.Close())
	}

	return f.Close()
}
