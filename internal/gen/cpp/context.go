package cpp

import (
	cfg "github.com/UoY-RoboStar/rtcg/internal/gen/config/cpp"
	"github.com/UoY-RoboStar/rtcg/internal/gen/gencommon"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

// Cpp renames cfg.Context to avoid duplication.
type Cpp cfg.Context

// Context is the context passed into the C++ code generator.
type Context struct {
	gencommon.Context

	Cpp
}

// NewContext creates a new template context from a named state machine.
func NewContext(name string, machine *stm.Stm, cctx cfg.Context) *Context {
	genCtx := gencommon.NewContext(name, machine)

	return &Context{
		Context: genCtx,
		Cpp:     Cpp(cctx),
	}
}
