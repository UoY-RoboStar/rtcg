// Package transition contains state transitions, sets, and aggregated sets.
package transition

import (
	"fmt"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"github.com/UoY-RoboStar/rtcg/internal/testlang/channel"
	"github.com/UoY-RoboStar/rtcg/internal/testlang/value"
)

// Transition is a transition from one state to another.
type Transition struct {
	Value value.Value     `json:"value,omitempty"` // Value is the value to observe before this transition may occur.
	Next  testlang.NodeID `json:"next,omitempty"`  // Next is the next node ID to jump to in the state machine.
	// TODO: record which test this transition is from
}

func (t Transition) String() string {
	return fmt.Sprintf("-[%s]-> %s", &t.Value, t.Next)
}

// Flat is a Transition with its channel attached.
type Flat struct {
	Transition

	Channel channel.Channel // Channel is the transition's channel.
}
