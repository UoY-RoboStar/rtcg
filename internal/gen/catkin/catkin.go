// Package catkin contains generators for Catkin build files.
package catkin

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/UoY-RoboStar/rtcg/internal/gen/subgen"

	"github.com/UoY-RoboStar/rtcg/internal/gen/templating"

	cfg "github.com/UoY-RoboStar/rtcg/internal/gen/config/catkin"
	"github.com/UoY-RoboStar/rtcg/internal/gen/gencommon"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

// Generator is a Catkin auxiliary file generator.
type Generator struct {
	config    cfg.Config       // config is the user-supplied configuration.
	srcDirSet gencommon.DirSet // srcBaseDir is the source directory set for this Catkin workspace.

	cmakeGen *templating.Generator // cmakeGen is the generator for CMakeLists.txt.
}

// New creates a new Catkin generator with the given configuration.
func New(config *cfg.Config, dirs gencommon.DirSet) (*Generator, error) {
	tg, err := NewTemplatedGenerator()
	if err != nil {
		return nil, err
	}

	gen := Generator{config: *config, srcDirSet: dirs.Subdir("src"), cmakeGen: tg}

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

func (g *Generator) OnSuite(suite *stm.Suite) subgen.OnSuite {
	return &OnSuite{
		parent: g,
		suite:  suite,
	}
}

// OnSuite is a generator specialised to generating one particular test suite.
type OnSuite struct {
	parent *Generator
	suite  *stm.Suite
}

func (o *OnSuite) Parent() subgen.Subgenerator {
	return o.parent
}

func (o *OnSuite) Dirs() []string {
	// Assume the C++ generator makes the appropriate directories.
	return nil
}

func (o *OnSuite) Generate() error {
	for name := range o.suite.Tests {
		dirs := o.parent.srcDirSet.Subdir(name)

		if err := o.generateTest(dirs, name); err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}
	}

	return nil
}

func (o *OnSuite) generateTest(dirs gencommon.DirSet, name string) error {
	// Take a copy here to avoid accidentally expanding the name in every test.
	pkg := *o.parent.config.Package
	pkg.Expand(name)

	if err := generatePackageXML(dirs.Output, pkg); err != nil {
		return err
	}

	return o.parent.cmakeGen.Generate(dirs.Output, pkg)
}

func generatePackageXML(dir string, pkg cfg.Package) error {
	w, err := os.Create(filepath.Join(dir, "package.xml"))
	if err != nil {
		return fmt.Errorf("couldn't create package.xml: %w", err)
	}

	err = pkg.Write(w)

	return errors.Join(err, w.Close())
}
