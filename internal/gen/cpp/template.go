package cpp

import (
	"embed"
	"fmt"
	"io/fs"

	"github.com/UoY-RoboStar/rtcg/internal/gen/templating"

	cfg "github.com/UoY-RoboStar/rtcg/internal/gen/config/cpp"
)

var (
	//go:embed embed/templates/base/*.cpp.tmpl embed/templates/base/convert/*.h.tmpl
	baseTemplates embed.FS
	//go:embed embed/templates/animate/*.cpp.tmpl embed/templates/animate/convert/*.h.tmpl
	animateTemplates embed.FS
	//go:embed embed/templates/ros/*.cpp.tmpl embed/templates/ros/convert/*.h.tmpl
	rosTemplates embed.FS
)

type templatedGenerators struct {
	common       *templating.Generator // testSpecific handles test-agnostic code files.
	testSpecific *templating.Generator // testSpecific handles test-specific code files.
}

// makeTemplatedGenerators sets up the templated generators for C++.
func makeTemplatedGenerators(config *cfg.Config) (templatedGenerators, error) {
	commonFiles := []templating.File{
		{Dir: "convert", Name: "convert.h", Desc: "type convert header", Glob: "convert/*.h.tmpl"},
	}

	testSpecificFiles := []templating.File{
		{Dir: "src", Name: "main.cpp", Desc: "main C++ file", Glob: "*.cpp.tmpl"},
	}

	builder := templating.SetBuilder{
		Srcs: []templating.Source{
			{Name: "base", Src: baseTemplates},
			variantSource(config.Variant),
		},
		Funcs: Funcs(),
	}

	var (
		tg  templatedGenerators
		err error
	)

	if tg.common, err = templating.NewGenerator(commonFiles, builder); err != nil {
		return tg, fmt.Errorf("couldn't create C++ template-based common files generator: %w", err)
	}

	if tg.testSpecific, err = templating.NewGenerator(testSpecificFiles, builder); err != nil {
		return tg, fmt.Errorf("couldn't create C++ template-based test files generator: %w", err)
	}

	return tg, nil
}

func variantSource(variant cfg.Variant) templating.Source {
	return templating.Source{
		Name: variant.String(),
		Src:  variantFS(variant),
	}
}

func variantFS(variant cfg.Variant) fs.FS {
	switch variant {
	case cfg.VariantAnimate:
		return animateTemplates
	case cfg.VariantRos:
		return rosTemplates
	default:
		// TODO: throw error if unknown variant?
		return nil
	}
}
