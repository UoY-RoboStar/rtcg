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
	//go:embed embed/templates/base/*.cpp.tmpl embed/templates/base/convert/*.h.tmpl
	baseTemplates embed.FS
	//go:embed embed/templates/animate/*.cpp.tmpl embed/templates/animate/convert/*.h.tmpl
	animateTemplates embed.FS
	//go:embed embed/templates/ros/*.cpp.tmpl embed/templates/ros/convert/*.h.tmpl
	rosTemplates embed.FS
)

// TemplateSet maps filenames to templates to use to generate them.
type TemplateSet map[string]*template.Template

// NewTemplateSet constructs a template set for a C++ variant, given the file list files.
func NewTemplateSet(variant Variant, files []TestFile) (TemplateSet, error) {
	tset := make(TemplateSet, len(files))

	varFS := variantFS(variant)
	varStr := variant.String()

	for _, file := range files {
		tmpl, err := parseTemplate(file, varFS, varStr)
		if err != nil {
			return nil, err
		}

		tset[file.Name] = tmpl
	}

	return tset, nil
}

// TODO: migrate this to a more generic package

func parseTemplate(file TestFile, varFS fs.FS, varStr string) (*template.Template, error) {
	builder := templateBuilder{
		tmpl:   template.New(file.Name),
		glob:   file.Glob,
		varFS:  varFS,
		varStr: varStr,
	}

	if err := builder.parse(); err != nil {
		return nil, err
	}

	return builder.tmpl, nil
}

type templateBuilder struct {
	tmpl   *template.Template
	glob   string
	varFS  fs.FS
	varStr string
}

func (t *templateBuilder) parse() error {
	t.tmpl = Funcs(t.tmpl)
	t.tmpl = gencommon.Funcs(t.tmpl)

	if err := t.parseFS(baseTemplates, "base"); err != nil {
		return err
	}

	if err := t.parseFS(t.varFS, t.varStr); err != nil {
		return err
	}

	return nil
}

func (t *templateBuilder) parseFS(fsys fs.FS, dir string) error {
	var err error

	t.tmpl, err = t.tmpl.ParseFS(fsys, path.Join("embed", "templates", dir, t.glob))
	if err != nil {
		return fmt.Errorf("couldn't open %s templates: %w", dir, err)
	}

	return nil
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
