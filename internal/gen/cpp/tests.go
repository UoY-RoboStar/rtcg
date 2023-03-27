package cpp

import (
	"path/filepath"

	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

func (g *Generator) GenerateTest(name string, test *stm.Stm) error {
	ctx := NewContext(name, test, g.config)

	// TODO(@MattWindsor91): only do this once
	if err := g.copyConvertFile(ctx); err != nil {
		return err
	}

	return g.Generator.Generate(filepath.Join(g.srcBaseDir, name), ctx)
}

// copyConvertFile copies convert.cpp from the input directory, if there is one.
func (g *Generator) copyConvertFile(ctx *Context) error {
	if !ctx.HasConversion {
		return nil
	}

	return g.copyLocalFile("convert.cpp")
}
