package cpp

import (
	"embed"
	"fmt"
	"io/fs"
	"path"
	"text/template"

	"github.com/UoY-RoboStar/rtcg/internal/gen/gencommon"
)

var (
	//go:embed embed/templates/base/*.cpp.tmpl
	baseTemplates embed.FS
	//go:embed embed/templates/animate/*.cpp.tmpl
	animateTemplates embed.FS
	//go:embed embed/templates/ros/*.cpp.tmpl
	rosTemplates embed.FS
)

// parseTemplate parses the C++ templates for the given variant.
func parseTemplate(variant Variant) (*template.Template, error) {
	var err error

	tmpl := Funcs(template.New(""))
	tmpl = gencommon.Funcs(tmpl)

	if tmpl, err = parseTemplateFS(tmpl, baseTemplates, "base"); err != nil {
		return nil, err
	}

	if tmpl, err = parseTemplateFS(tmpl, variantFS(variant), variant.String()); err != nil {
		return nil, err
	}

	return tmpl, nil
}

func parseTemplateFS(tmpl *template.Template, fsys fs.FS, dir string) (*template.Template, error) {
	base, err := tmpl.ParseFS(fsys, path.Join("embed", "templates", dir, "*.cpp.tmpl"))
	if err != nil {
		return nil, fmt.Errorf("couldn't open %s templates: %w", dir, err)
	}

	return base, nil
}

func variantFS(variant Variant) fs.FS {
	switch variant {
	case VariantAnimate:
		return animateTemplates
	case VariantRos:
		return rosTemplates
	default:
		// TODO: throw error if unknown variant?
		return nil
	}
}
