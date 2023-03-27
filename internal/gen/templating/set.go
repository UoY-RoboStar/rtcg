package templating

import (
	"fmt"
	"path"
	"path/filepath"
	"text/template"
)

// Set maps filenames to templates to use to generate them.
type Set map[string]*template.Template

// Generate generates each of the templated files in this set into dir.
// We pass ctx to each template.
func (s Set) Generate(files []File, dir string, ctx any) error {
	for _, file := range files {
		if err := s.generateFile(file, dir, ctx); err != nil {
			return err
		}
	}

	return nil
}

func (s Set) generateFile(file File, dir string, ctx any) error {
	outPath := filepath.Join(dir, file.Dir, file.Name)

	err := CreateFile(outPath, file.Name+".tmpl", s[file.Name], ctx)
	if err != nil {
		return fmt.Errorf("couldn't generate %s for %s: %w", file.Desc, dir, err)
	}

	return nil
}

// SetBuilder is a builder for Set.
type SetBuilder struct {
	Srcs  []Source         // Srcs is an ordered set of template sources.
	Funcs template.FuncMap // Funcs is the custom function map to apply to the templates.
}

// BuildFiles uses the builder's source set to parse templates for files.
func (b *SetBuilder) BuildFiles(files []File) (Set, error) {
	tset := make(Set, len(files))

	for _, file := range files {
		tmpl, err := b.buildFile(file)
		if err != nil {
			return nil, err
		}

		tset[file.Name] = tmpl
	}

	return tset, nil
}

func (b *SetBuilder) buildFile(file File) (*template.Template, error) {
	tmpl := template.New(file.Name)
	tmpl = tmpl.Funcs(b.Funcs)
	tmpl = Funcs(tmpl)

	var err error

	for _, src := range b.Srcs {
		tmpl, err = tmpl.ParseFS(src.Src, path.Join("embed", "templates", src.Name, file.Glob))
		if err != nil {
			return nil, fmt.Errorf("couldn't open %s templates: %w", src.Name, err)
		}
	}

	return tmpl, nil
}
