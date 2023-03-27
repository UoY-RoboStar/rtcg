package makefile

import (
	"github.com/UoY-RoboStar/rtcg/internal/gen/config/cpp"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

// Context is the context passed into the Makefile code generator.
type Context struct {
	Tests map[string]*stm.Stm // Tests is the set of tests for which we are generating Makefile rules.

	cpp.Context
}

// NewContext creates a new template context from a test suite.
func NewContext(suite *stm.Suite, cfg cpp.Context) *Context {
	return &Context{Tests: suite.Tests, Context: cfg}
}
