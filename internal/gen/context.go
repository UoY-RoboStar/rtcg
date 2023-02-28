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

	Transitions TransitionContext // Transitions is a pre-calculated transition context.
}

// NewContext creates a new template context from a named state machine.
func NewContext(name string, machine *stm.Stm) Context {
	return Context{
		Name:        name,
		Date:        time.Now(),
		Stm:         machine,
		Transitions: makeTransitions(machine),
	}
}

// TransitionContext provides a pre-computed filtered list of transition sets.
//
// This context, unlike Stm, has a lot of sub-normalised fields.
type TransitionContext struct {
	All []transition.AggregateSet // All is the list of all transitions.
	In  []transition.AggregateSet // In is the list of all transitions that are inputs.
	Out []transition.AggregateSet // Out is the list of all transitions that are outputs.

	InMerged transition.StateMap // InMerged is a merger of all state maps from In.

	HasIn           bool // HasIn is true if In is non-empty.
	FirstStateHasIn bool // HasIn is true if the first state contains an input transition.
}

func (c *TransitionContext) addInput(t transition.AggregateSet) {
	c.HasIn = true
	c.In = append(c.In, t)
	for s, ts := range t.States {
		c.InMerged[s] = append(c.InMerged[s], ts...)
	}
}

func (c *TransitionContext) populateFilters() {
	for _, t := range c.All {
		if t.Channel.IsIn() {
			c.addInput(t)
		} else {
			c.Out = append(c.Out, t)
		}
	}
}

func makeTransitions(machine *stm.Stm) TransitionContext {
	all := machine.TransitionSets()
	count := len(all)

	tctx := TransitionContext{
		All:             all,
		In:              make([]transition.AggregateSet, 0, count),
		Out:             make([]transition.AggregateSet, 0, count),
		InMerged:        make(transition.StateMap),
		HasIn:           false,
		FirstStateHasIn: firstStateHasIn(machine),
	}

	tctx.populateFilters()

	return tctx
}

func firstStateHasIn(machine *stm.Stm) bool {
	if len(machine.States) == 0 {
		return false
	}

	for _, ts := range machine.States[0].TransitionSets {
		if ts.Channel.IsIn() {
			return true
		}
	}

	return false
}
