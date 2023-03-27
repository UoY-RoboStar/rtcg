// Package catkin contains generators for Catkin build files.
package catkin

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/UoY-RoboStar/rtcg/internal/gen/templating"

	cfg "github.com/UoY-RoboStar/rtcg/internal/gen/config/catkin"
	"github.com/UoY-RoboStar/rtcg/internal/gen/gencommon"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

// Generator is a Catkin auxiliary file generator.
type Generator struct {
	config     cfg.Config       // config is the user-supplied configuration.
	dirSet     gencommon.DirSet // dirSet is the directory set for this Catkin workspace.
	srcBaseDir string           // srcBaseDir is the output source directory for this Generator.

	cmakeGen *templating.Generator // cmakeGen is the generator for CMakeLists.txt.
}

// New creates a new Catkin generator with the given configuration.
func New(config *cfg.Config, dirs gencommon.DirSet) (*Generator, error) {
	tg, err := NewTemplatedGenerator()
	if err != nil {
		return nil, err
	}

	gen := Generator{config: *config, dirSet: dirs, srcBaseDir: dirs.SrcDir(), cmakeGen: tg}

	if gen.config.Package == nil {
		var pkg cfg.Package
		gen.config.Package = &pkg
	}

	gen.config.Package.Autofill()

	return &gen, nil
}

func (g *Generator) Name() string {
	return "Catkin"
}

func (g *Generator) Dirs(_ stm.Suite) []string {
	// Assume the C++ generator makes the appropriate directories.
	return nil
}

func (g *Generator) Generate(suite stm.Suite) error {
	if err := gencommon.GenerateTests(suite, g); err != nil {
		return fmt.Errorf("couldn't generate for tests: %w", err)
	}

	return nil
}

func (g *Generator) GenerateTest(name string, _ *stm.Stm) error {
	dir := filepath.Join(g.srcBaseDir, name)

	// Take a copy here to avoid accidentally expanding the name in every test.
	pkg := *g.config.Package
	pkg.Expand(name)

	if err := generatePackageXML(dir, pkg); err != nil {
		return err
	}

	return g.cmakeGen.Generate(dir, pkg)
}

func generatePackageXML(dir string, pkg cfg.Package) error {
	w, err := os.Create(filepath.Join(dir, "package.xml"))
	if err != nil {
		return fmt.Errorf("couldn't create package.xml: %w", err)
	}

	err = pkg.Write(w)

	return errors.Join(err, w.Close())
}
