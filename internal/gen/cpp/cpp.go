// Package cpp contains the C++ code generator.
package cpp

import (
	"fmt"

	cfg "github.com/UoY-RoboStar/rtcg/internal/gen/config/cpp"
	"github.com/UoY-RoboStar/rtcg/internal/gen/gencommon"
	"github.com/UoY-RoboStar/rtcg/internal/gen/subgen"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

// Generator is a C++ code generator.
type Generator struct {
	config    cfg.Config       // config is the configuration for this Generator.
	srcDirSet gencommon.DirSet // srcDirSet is the source directory set for this Generator.

	templated templatedGenerators // templated holds the template-based generators for C++ code.
}

// New constructs a new C++ code generator from config, rooted at the given directories.
func New(config *cfg.Config, dirs gencommon.DirSet) (*Generator, error) {
	var (
		gen Generator
		err error
	)

	gen.config = *config
	gen.srcDirSet = dirs.Subdir("src")

	if gen.templated, err = makeTemplatedGenerators(config); err != nil {
		return nil, err
	}

	return &gen, nil
}

func (g *Generator) Name() string {
	return fmt.Sprintf("C++ (%s)", g.config.Variant)
}

func (g *Generator) OnSuite(suite *stm.Suite) subgen.OnSuite {
	return &OnSuite{
		parent: g,
		suite:  suite,
		cctx:   g.config.Process(suite.Types),
	}
}

type OnSuite struct {
	parent *Generator
	suite  *stm.Suite
	cctx   cfg.Context
}

func (o *OnSuite) Parent() subgen.Subgenerator {
	return o.parent
}

func (o *OnSuite) Dirs() []string {
	dirs := make([]string, 1, len(o.suite.Tests)+1)
	dirs[0] = o.parent.srcDirSet.OutputPath(preludeDir)

	if o.cctx.HasConversion {
		dirs = append(dirs, o.parent.srcDirSet.OutputPath(convertDir))
	}

	for name := range o.suite.Tests {
		dirs = append(dirs, o.testDirs(name)...)
	}

	return dirs
}

func (o *OnSuite) testDirs(name string) []string {
	dirSet := o.parent.srcDirSet.Subdir(name)

	// This directory structure mirrors that of catkin, even if we're not generating ROS.
	return []string{
		dirSet.OutputPath("include"),
		dirSet.OutputPath("src"),
	}
}

func (o *OnSuite) Generate() error {
	if err := o.parent.copyPrelude(); err != nil {
		return err
	}

	if err := o.generateCommon(); err != nil {
		return err
	}

	return o.generateTestSpecific()
}

func (o *OnSuite) generateCommon() error {
	if err := o.copyConvertFile(); err != nil {
		return err
	}

	if err := o.parent.templated.common.Generate(o.parent.srcDirSet.Output, o.cctx); err != nil {
		return fmt.Errorf("generating common files: %w", err)
	}

	return nil
}

// copyConvertFile copies convert.cpp from the input directory, if there is one.
func (o *OnSuite) copyConvertFile() error {
	if !o.cctx.HasConversion {
		return nil
	}

	return copyLocalFile(o.parent.srcDirSet.Subdir(convertDir), "convert.cpp")
}

func (o *OnSuite) generateTestSpecific() error {
	for name, test := range o.suite.Tests {
		dirs := o.parent.srcDirSet.Subdir(name)

		if err := o.generateTest(dirs, name, test); err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}
	}

	return nil
}

func (o *OnSuite) generateTest(dirs gencommon.DirSet, name string, test *stm.Stm) error {
	ctx := NewContext(name, test, o.cctx)

	if err := o.parent.templated.testSpecific.Generate(dirs.Output, ctx); err != nil {
		return fmt.Errorf("generating test-specific files: %w", err)
	}

	return nil
}
