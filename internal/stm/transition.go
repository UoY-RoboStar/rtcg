package stm

import (
	"fmt"
	"strings"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

// TransitionSet is a set of transitions for a given channel.
type TransitionSet struct {
	Channel     testlang.Channel // Channel is the channel at the head of this transition set.
	Transitions []Transition     // Transitions is the list of transitions.
}

func (t TransitionSet) String() string {
	tstrs := make([]string, len(t.Transitions))
	for i, v := range t.Transitions {
		tstrs[i] = v.String()
	}

	return fmt.Sprintf("%s:{%s}", &t.Channel, strings.Join(tstrs, ", "))
}

// Transition is a transition from one state to another.
type Transition struct {
	Value testlang.Value  // Value is the value that must be observed or sent for this transition to occur.
	Next  testlang.NodeID // Next is the next node ID to jump to in the state machine.
	// TODO: record which test this transition is from
}

func (t Transition) String() string {
	return fmt.Sprintf("-[%s]-> %s", &t.Value, t.Next)
}
