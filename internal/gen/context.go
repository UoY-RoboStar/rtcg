package gen

import (
	"time"

	"github.com/UoY-RoboStar/rtcg/internal/stm/transition"

	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

// Context is the context passed into the template.
type Context struct {
	Name string    // Name is the name of the test case being generated.
	Date time.Time // Date is the time of generation.

	Stm *stm.Stm // Stm is the state machine being generated.

	Transitions []transition.AggregateSet // Transitions is a pre-calculated transition set.
	Inputs      []transition.AggregateSet // Inputs is the list of all Transitions that are inputs.
	Outputs     []transition.AggregateSet // Outputs is the list of all Transitions that are outputs.
}

// NewContext creates a new template context from a named state machine.
func NewContext(name string, stm *stm.Stm) Context {
	ctx := Context{
		Name: name,
		Date: time.Now(),
		Stm:  stm,
	}

	ctx.populateTransitions()

	return ctx
}

func (c *Context) populateTransitions() {
	c.Transitions = c.Stm.TransitionSets()
	nTransitions := len(c.Transitions)

	c.Inputs = make([]transition.AggregateSet, 0, nTransitions)
	c.Outputs = make([]transition.AggregateSet, 0, nTransitions)

	for _, t := range c.Transitions {
		if t.Channel.IsIn() {
			c.Inputs = append(c.Inputs, t)
		} else {
			c.Outputs = append(c.Outputs, t)
		}
	}
}
