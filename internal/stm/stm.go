// Package stm contains the testing state machine.
package stm

import (
	"fmt"
	"rtcg/internal/testlang"
	"strings"
)

// Stm is a state machine.
type Stm struct {
	// Nodes is the list of nodes in this state machine.
	//
	// Conventionally, the first node in the machine is the initial node.
	Nodes []Node
}

// Node is a state machine node.
type Node struct {
	ID             testlang.NodeID // ID is the identifier of the Node.
	TransitionSets []TransitionSet // TransitionSets is the list of transition sets out of this node.

	IsFail bool            // IsFail is set if any test fails on this node.
	Fails  map[string]bool // Fails is a set of tests the node is known to fail.  (IsFail can be true even if empty).
}

// AddTransition adds tr onto the transition set for channel ch.
//
// If no such transition set exists, one is created.
func (n *Node) AddTransition(ch testlang.Channel, tr Transition) {
	// Try merging onto an existing channel set.
	for i := range n.TransitionSets {
		ts := &n.TransitionSets[i]
		if ts.Channel.Equals(ch) {
			ts.Transitions = append(ts.Transitions, tr)
			return
		}
	}
	// No transition set with this channel exists yet.
	n.TransitionSets = append(n.TransitionSets, TransitionSet{
		Channel: ch, Transitions: []Transition{tr},
	})
}

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
	Value  testlang.Value  // Value is the value that must be observed or sent for this transition to occur.
	Status testlang.Status // Status is the status before taking this transition.
	Next   testlang.NodeID // Next is the next node ID to jump to in the state machine.
	// TODO: record which test this transition is from
}

func (t Transition) String() string {
	return fmt.Sprintf("%s -[%s]-> %s", &t.Status, &t.Value, t.Next)
}
