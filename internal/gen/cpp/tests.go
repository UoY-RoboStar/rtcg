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

	if err := g.copyConvertFile(ctx); err != nil {
		return err
	}

	return g.generateTestTemplatedFiles(ctx)
}

// copyConvertFile copies convert.cpp from the input directory, if there is one.
func (g *Generator) copyConvertFile(ctx *Context) error {
	if !ctx.HasConversion {
		return nil
	}

	return g.copyLocalFile("convert.cpp")
}

func (g *Generator) generateTestTemplatedFiles(ctx *Context) error {
	for _, file := range g.testFiles {
		if err := g.generateTestFile(ctx, file); err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) generateTestFile(ctx *Context, file TestFile) error {
	outPath := g.testSourcePath(ctx.Name, file.Name)
	err := gencommon.ExecuteTemplateOnFile(outPath, file.Name+".tmpl", g.templates[file.Name], ctx)
	if err != nil {
		return fmt.Errorf("couldn't generate %s for %s: %w", file.Desc, ctx.Name, err)
	}
	return nil
}

func (g *Generator) testSourcePath(name, fileName string) string {
	return filepath.Join(g.outputDir, srcDir, name, fileName)
}
