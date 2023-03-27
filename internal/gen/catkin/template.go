package catkin

import (
	"embed"
	"fmt"
	"text/template"

	"github.com/UoY-RoboStar/rtcg/internal/gen/templating"
)

//go:embed embed/templates/CMakeLists.txt.tmpl
var templates embed.FS

// NewTemplatedGenerator sets up a templated generator for Catkin.
func NewTemplatedGenerator() (templating.Generator, error) {
	testFiles := []templating.File{
		{Dir: "", Name: "CMakeLists.txt", Desc: "package cmake file", Glob: "CMakeLists.txt.tmpl"},
	}

	builder := templating.SetBuilder{
		Srcs: []templating.Source{{
			Name: "",
			Src:  templates,
		}},
		Funcs: template.FuncMap{},
	}

	gen, err := templating.NewGenerator(testFiles, builder)
	if err != nil {
		return gen, fmt.Errorf("couldn't create template-based generator: %w", err)
	}

	return gen, nil
}
