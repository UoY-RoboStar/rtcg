package gencommon

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

// TemplatedGenerator is a mixin struct for dealing with templated parts of generators.
type TemplatedGenerator struct {
	TestFiles []TestFile  // TestFiles is the list of files this generator will make.
	Templates TemplateSet // Templates is the map of templates to use for files in testFiles.
}

// ExecuteTemplateOnFile creates path, executes tmpl with dot ctx, and handles closing.
func ExecuteTemplateOnFile(path, tmplName string, tmpl *template.Template, ctx any) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("couldn't create output file: %w", err)
	}

	err = tmpl.ExecuteTemplate(file, tmplName, ctx)

	return errors.Join(err, file.Close())
}

// TemplateSet maps filenames to templates to use to generate them.
type TemplateSet map[string]*template.Template

// Generate generates each of the templated files for the test with the given name, rooted at path.
// We pass ctx to each template.
func (s TemplateSet) Generate(files []TestFile, path, name string, ctx any) error {
	for _, file := range files {
		if err := s.generateFile(file, path, name, ctx); err != nil {
			return err
		}
	}

	return nil
}

func (s TemplateSet) generateFile(file TestFile, path, name string, ctx any) error {
	outPath := filepath.Join(path, name, file.Dir, file.Name)

	err := ExecuteTemplateOnFile(outPath, file.Name+".tmpl", s[file.Name], ctx)
	if err != nil {
		return fmt.Errorf("couldn't generate %s for %s: %w", file.Desc, name, err)
	}

	return nil
}

// TestFile holds information about a templated test file.
type TestFile struct {
	Dir  string // Dir is the destination directory of this file, within the test directory.
	Name string // Name is the filename of this file.
	Desc string // Desc is a human-readable description for this file.
	Glob string // Glob is the slash-delimited glob of source templates for this file.
}

// TemplateSource names a source of template files.
type TemplateSource struct {
	Name string // Name is the name of the source.
	Src  fs.FS  // Src is the source filesystem.
}

// NewTemplatedGenerator creates a TemplatedGenerator using files and builder.
func NewTemplatedGenerator(files []TestFile, builder TemplateBuilder) (TemplatedGenerator, error) {
	var tg TemplatedGenerator

	tset, err := builder.BuildFiles(files)
	if err != nil {
		return tg, err
	}

	tg.TestFiles = files
	tg.Templates = tset

	return tg, nil
}

// TemplateBuilder is a builder for template files.
type TemplateBuilder struct {
	Srcs  []TemplateSource // Srcs is an ordered set of template sources.
	Funcs template.FuncMap // Funcs is the custom function map to apply to the templates.
}

// BuildFiles uses the builder's source set to parse templates for files.
func (b *TemplateBuilder) BuildFiles(files []TestFile) (TemplateSet, error) {
	tset := make(TemplateSet, len(files))

	for _, file := range files {
		tmpl, err := b.buildFile(file)
		if err != nil {
			return nil, err
		}

		tset[file.Name] = tmpl
	}

	return tset, nil
}

func (b *TemplateBuilder) buildFile(file TestFile) (*template.Template, error) {
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
