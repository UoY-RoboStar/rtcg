package catkin

import (
	"embed"
	"fmt"
	"text/template"

	"github.com/UoY-RoboStar/rtcg/internal/gen/gencommon"
)

//go:embed embed/templates/CMakeLists.txt.tmpl
var templates embed.FS

// NewTemplatedGenerator sets up a templated generator for Catkin.
func NewTemplatedGenerator() (gencommon.TemplatedGenerator, error) {
	testFiles := []gencommon.TestFile{
		{Dir: "", Name: "CMakeLists.txt", Desc: "package cmake file", Glob: "CMakeLists.txt.tmpl"},
	}

	builder := gencommon.TemplateBuilder{
		Srcs: []gencommon.TemplateSource{{
			Name: "main",
			Src:  templates,
		}},
		Funcs: template.FuncMap{},
	}

	gen, err := gencommon.NewTemplatedGenerator(testFiles, builder)
	if err != nil {
		return gen, fmt.Errorf("couldn't create Catkin++ template-based generator: %w", err)
	}

	return gen, nil
}
