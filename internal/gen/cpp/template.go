package cpp

import (
	"embed"
	"fmt"
	"io/fs"

	cfg "github.com/UoY-RoboStar/rtcg/internal/gen/config/cpp"
	"github.com/UoY-RoboStar/rtcg/internal/gen/gencommon"
)

var (
	//go:embed embed/templates/base/*.cpp.tmpl embed/templates/base/convert/*.h.tmpl
	baseTemplates embed.FS
	//go:embed embed/templates/animate/*.cpp.tmpl embed/templates/animate/convert/*.h.tmpl
	animateTemplates embed.FS
	//go:embed embed/templates/ros/*.cpp.tmpl embed/templates/ros/convert/*.h.tmpl
	rosTemplates embed.FS
)

// NewTemplatedGenerator sets up a templated generator for C++.
func NewTemplatedGenerator(config *cfg.Config) (gencommon.TemplatedGenerator, error) {
	testFiles := []gencommon.TestFile{
		{Dir: "src", Name: "main.cpp", Desc: "main C++ file", Glob: "*.cpp.tmpl"},
		{Dir: "include", Name: "convert.h", Desc: "type convert header", Glob: "convert/*.h.tmpl"},
	}

	builder := gencommon.TemplateBuilder{
		Srcs: []gencommon.TemplateSource{
			{Name: "base", Src: baseTemplates},
			variantSource(config.Variant),
		},
		Funcs: Funcs(),
	}

	gen, err := gencommon.NewTemplatedGenerator(testFiles, builder)
	if err != nil {
		return gen, fmt.Errorf("couldn't create C++ template-based generator: %w", err)
	}

	return gen, nil
}

func variantSource(variant cfg.Variant) gencommon.TemplateSource {
	return gencommon.TemplateSource{
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
