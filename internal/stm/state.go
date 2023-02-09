package stm

import "rtcg/internal/testlang"

// State is a state machine state (with attached transitions).
type State struct {
	ID             testlang.NodeID // ID is the identifier of the State.
	TransitionSets []TransitionSet // TransitionSets is the list of transition sets out of this state.

	IsFail bool            // IsFail is set if any test fails on this state.
	Fails  map[string]bool // Fails is a set of tests the state is known to fail.  (IsFail can be true even if empty).
}

// AddTransitionToNode adds a transition from this node to the test-tree node n.
//
// We assume the node has already been assigned an ID.
func (s *State) AddTransitionToNode(n *testlang.Node) {
	tr := Transition{Status: n.Status, Value: n.Event.Value, Next: n.ID}
	s.AddTransition(n.Event.Channel, tr)
}

// AddTransition adds tr onto the transition set for channel ch.
//
// If no such transition set exists, one is created.
func (s *State) AddTransition(ch testlang.Channel, tr Transition) {
	// Try merging onto an existing channel set.
	for i := range s.TransitionSets {
		ts := &s.TransitionSets[i]
		if ts.Channel.Equals(ch) {
			ts.Transitions = append(ts.Transitions, tr)
			return
		}
	}
	// No transition set with this channel exists yet.
	s.TransitionSets = append(s.TransitionSets, TransitionSet{
		Channel: ch, Transitions: []Transition{tr},
	})
}

// MarkFailuresFromNode marks s as a failing state (if applicable) and adds information about each test it fails from n.
func (s *State) MarkFailuresFromNode(n testlang.Node) {
	if n.Status != testlang.StatFail {
		return
	}
	s.IsFail = true
	for _, t := range n.Tests {
		s.Fails[t] = true
	}
}
