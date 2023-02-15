package stm

import (
	"fmt"
	"strings"

	"github.com/UoY-RoboStar/rtcg/internal/testlang/comm"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

// TransitionSet is a set of transitions for a given channel.
type TransitionSet struct {
	Channel     comm.Channel `json:"channel"`               // Channel is the channel at the head of this set.
	Transitions []Transition `json:"transitions,omitempty"` // Transitions is the list of transitions.
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
	Value testlang.Value  `json:"value,omitempty"` // Value is the value to observe before this transition may occur.
	Next  testlang.NodeID `json:"next,omitempty"`  // Next is the next node ID to jump to in the state machine.
	// TODO: record which test this transition is from
}

func (t Transition) String() string {
	return fmt.Sprintf("-[%s]-> %s", &t.Value, t.Next)
}
