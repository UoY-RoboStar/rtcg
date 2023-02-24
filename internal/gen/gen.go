// Package gen concerns the generation (template expansion) part of rtcg.
package gen

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"github.com/UoY-RoboStar/rtcg/internal/gen/cppfunc"
	"github.com/UoY-RoboStar/rtcg/internal/gen/stdfunc"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

var (
	//go:embed embed/templates/*.cpp.tmpl embed/templates/base.Makefile.tmpl
	baseTemplates embed.FS
	//go:embed embed/prelude/*
	prelude embed.FS
)

const (
	outputDirPerms = 0o754           // outputDirPerms is the output directory permission mask.
	srcDir         = "src"           // srcDir is the subdirectory of the output directory for source code.
	preludeDir     = "rtcg"          // preludeDir is the subdirectory of srcDir where we store the prelude.
	preludeMount   = "embed/prelude" // preludeMount is the directory in prelude where the prelude is.
)

// Generator is a test code generator.
type Generator struct {
	template  *template.Template // template is the template to use for C++ files.
	makefile  *template.Template // makefile is the template to use for Makefiles.
	outputDir string             // outputDir is the output directory for the tests.
}

// New creates a new Generator by reading all templates from inFS, and outputting to outDir.
func New(inFS fs.FS, outDir string) (*Generator, error) {
	generator := Generator{outputDir: outDir, template: nil}

	var err error

	if generator.template, err = parseCppTemplate(inFS); err != nil {
		return nil, err
	}

	if generator.makefile, err = parseMakefileTemplate(inFS); err != nil {
		return nil, err
	}

	return &generator, nil
}

func parseCppTemplate(inFS fs.FS) (*template.Template, error) {
	base := cppfunc.Funcs(template.New(""))
	base = stdfunc.Funcs(base)

	base, err := base.ParseFS(baseTemplates, "embed/templates/*.cpp.tmpl")
	if err != nil {
		return nil, fmt.Errorf("couldn't open base templates: %w", err)
	}

	tmpl, err := base.ParseFS(inFS, "*.cpp.tmpl")
	if err != nil {
		return nil, fmt.Errorf("couldn't open user templates: %w", err)
	}

	return tmpl, nil
}

func parseMakefileTemplate(inFS fs.FS) (*template.Template, error) {
	base, err := template.ParseFS(baseTemplates, "embed/templates/base.Makefile.tmpl")
	if err != nil {
		return nil, fmt.Errorf("couldn't open base Makefile template: %w", err)
	}

	tmpl, err := base.ParseFS(inFS, "Makefile.tmpl")
	if err != nil {
		return nil, fmt.Errorf("couldn't open user Makefile template: %w", err)
	}

	return tmpl, nil
}

func (g *Generator) Generate(suite stm.Suite) error {
	if err := g.mkdirs(); err != nil {
		return err
	}

	if err := g.copyPrelude(); err != nil {
		return err
	}

	if err := g.generateMakefile(suite); err != nil {
		return err
	}

	if err := g.generateSuite(suite); err != nil {
		return err
	}

	return nil
}

// mkdirs makes the various directories used by the
func (g *Generator) mkdirs() error {
	if err := g.mkdir("test"); err != nil {
		return err
	}

	if err := g.mkdir("test", srcDir); err != nil {
		return err
	}

	if err := g.mkdir("test", srcDir, preludeDir); err != nil {
		return err
	}

	return nil
}

// mkdir makes an output directory with the given name and path fragments.
func (g *Generator) mkdir(name string, fragments ...string) error {
	dirPath := filepath.Join(append([]string{g.outputDir}, fragments...)...)

	if err := os.MkdirAll(dirPath, outputDirPerms); err != nil {
		return fmt.Errorf("couldn't make %s directory at %q: %w", name, g.outputDir, err)
	}

	return nil
}

func (g *Generator) copyPrelude() error {
	ents, err := fs.ReadDir(prelude, preludeMount)
	if err != nil {
		return fmt.Errorf("couldn't find prelude in embedded files: %w", err)
	}

	for _, e := range ents {
		if err := g.copyPreludeFile(e.Name()); err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) copyPreludeFile(name string) error {
	src, err := prelude.Open(path.Join(preludeMount, name))
	if err != nil {
		return fmt.Errorf("couldn't open prelude file %q: %w", name, err)
	}

	dst, err := os.Create(filepath.Join(g.outputDir, srcDir, preludeDir, name))
	if err != nil {
		err = fmt.Errorf("couldn't create prelude file %q: %w", name, err)
		return errors.Join(err, src.Close())
	}

	_, err = io.Copy(dst, src)
	if err != nil {
		err = fmt.Errorf("couldn't copy prelude file %q: %w", name, err)
	}

	return errors.Join(err, dst.Close(), src.Close())
}

func (g *Generator) generateMakefile(suite stm.Suite) error {
	outPath := filepath.Join(g.outputDir, "Makefile")
	return executeTemplateOnFile("Makefile", outPath, "Makefile.tmpl", g.makefile, suite)
}

func (g *Generator) generateSuite(suite stm.Suite) error {
	for k, v := range suite {
		if err := g.generateStm(k, v); err != nil {
			return fmt.Errorf("couldn't generate test %s: %w", k, err)
		}
	}
	return nil
}

func (g *Generator) generateStm(name string, body *stm.Stm) error {
	outPath := filepath.Join(g.outputDir, srcDir, name+".cpp")

	ctx := NewContext(name, body)

	return executeTemplateOnFile("test-case "+name, outPath, "top.cpp.tmpl", g.template, ctx)
}

func executeTemplateOnFile(name, path, tmplName string, tmpl *template.Template, ctx any) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("couldn't create output file for %s: %w", name, err)
	}

	err = tmpl.ExecuteTemplate(file, tmplName, ctx)

	return errors.Join(err, file.Close())
}
