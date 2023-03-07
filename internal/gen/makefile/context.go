package makefile

import (
	"fmt"

	"github.com/UoY-RoboStar/rtcg/internal/gen/cpp"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

// Context is the context passed into the Makefile code generator.
type Context struct {
	Tests stm.Suite // Tests is the set of tests for which we are generating Makefile rules.

	cpp.ConfigContext
}

// NewContext creates a new template context from a test suite.
func NewContext(tests stm.Suite, cfg cpp.Config) (*Context, error) {
	types, err := tests.UnifiedTypes()
	if err != nil {
		return nil, fmt.Errorf("couldn't get channel types to agree: %w", err)
	}

	ctx := Context{
		Tests:         tests,
		ConfigContext: cfg.Process(types),
	}

	return &ctx, nil
}
