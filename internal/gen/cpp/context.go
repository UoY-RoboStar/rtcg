package cpp

import (
	"github.com/UoY-RoboStar/rtcg/internal/gen/gencommon"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

// Context is the context passed into the C++ code generator.
type Context struct {
	gencommon.Context

	Includes []Include // Includes contains the user-configured includes.
}

// NewContext creates a new template context from a named state machine.
func NewContext(name string, machine *stm.Stm, config Config) Context {
	return Context{
		Context:  gencommon.NewContext(name, machine),
		Includes: config.Includes,
	}
}
