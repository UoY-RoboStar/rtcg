package cpp

import (
	"fmt"
	"path/filepath"

	"github.com/UoY-RoboStar/rtcg/internal/gen/gencommon"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

func (g *Generator) generateTests(suite stm.Suite) error {
	for name, test := range suite {
		if err := g.generateTest(name, test); err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) generateTest(name string, test *stm.Stm) error {
	ctx := NewContext(name, test, g.config)
	gen := TestGenerator{ctx: ctx, parent: g}

	return gen.generate()
}

type TestGenerator struct {
	ctx    *Context
	parent *Generator
}

func (t *TestGenerator) generate() error {
	if err := t.copyConvertFile(); err != nil {
		return err
	}

	return t.generateTestTemplatedFiles()
}

// copyConvertFile copies convert.cpp from the input directory, if there is one.
func (t *TestGenerator) copyConvertFile() error {
	if !t.ctx.HasConversion {
		return nil
	}

	return t.parent.copyLocalFile("convert.cpp")
}

func (t *TestGenerator) generateTestTemplatedFiles() error {
	for _, file := range t.parent.testFiles {
		if err := t.generateTestFile(file); err != nil {
			return err
		}
	}

	return nil
}

func (t *TestGenerator) generateTestFile(file TestFile) error {
	outPath := t.testSourcePath(file)

	err := gencommon.ExecuteTemplateOnFile(outPath, file.Name+".tmpl", t.parent.templates[file.Name], t.ctx)
	if err != nil {
		return fmt.Errorf("couldn't generate %s for %s: %w", file.Desc, t.ctx.Name, err)
	}

	return nil
}

func (t *TestGenerator) testSourcePath(file TestFile) string {
	return filepath.Join(t.parent.outputDir, srcDir, t.ctx.Name, file.Dir, file.Name)
}
