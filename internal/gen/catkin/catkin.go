// Package catkin contains generators for Catkin build files.
package catkin

import (
	"path/filepath"

	cfg "github.com/UoY-RoboStar/rtcg/internal/gen/config/catkin"
	"github.com/UoY-RoboStar/rtcg/internal/gen/gencommon"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

// Generator is a Catkin auxiliary file generator.
type Generator struct {
	config       cfg.Config // config is the user-supplied configuration.
	workspaceDir string     // workspaceDir is the Catkin workspace directory.

	gencommon.TemplatedGenerator
}

// New creates a new Catkin generator with the given configuration.
func New(config *cfg.Config, dirs gencommon.DirSet) (*Generator, error) {
	tg, err := NewTemplatedGenerator()
	if err != nil {
		return nil, err
	}

	gen := Generator{config: *config, workspaceDir: filepath.Clean(dirs.Output), TemplatedGenerator: tg}
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

func (g *Generator) Generate(_ stm.Suite) error {
	return nil
}
