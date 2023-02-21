package stm

import (
	"fmt"
	"github.com/UoY-RoboStar/rtcg/internal/stm/verdict"
	"strings"

	"github.com/UoY-RoboStar/rtcg/internal/stm/transition"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

// State is a state machine state (with attached transitions).
type State struct {
	// ID is the identifier of the State.
	//
	// Unlike test trees, each state machine state MUST have an identifier.
	ID testlang.NodeID `json:"id"`

	// TransitionSets is the list of transition sets out of this state.
	//
	// Each transition set maps a particular channel to a list of transitions predicated on that channel's value.
	TransitionSets []transition.Set `json:"transitionSets,omitempty"`

	Verdict *verdict.Verdict `json:"verdicts,omitempty"` // Verdict holds the test verdicts that this state reports.
}

// NewState creates a new State with the given id.
func NewState(id testlang.NodeID) *State {
	return &State{
		ID:             id,
		TransitionSets: nil,
		Verdict:        &verdict.Verdict{},
	}
}

// AddOutgoingNode handles all bookkeeping for logging verdicts and producing transitions from s to the state for node.
func (s *State) AddOutgoingNode(node *testlang.Node) {
	s.addVerdictsFromNode(node)
	s.addTransitionToNode(node)
}

// addTransitionToNode adds a transition from this state to the given test-tree node.
//
// We assume the node has already been assigned an ID.
func (s *State) addTransitionToNode(node *testlang.Node) {
	// We don't add transitions to failing nodes; they are just sentinels with no test content.
	if node.Outcome == testlang.OutcomeFail {
		return
	}

	if node.ID == "" {
		panic("should have assigned an ID to node")
	}

	tr := transition.Transition{Value: node.Event.Value, Next: node.ID}
	s.TransitionSets = transition.AddToSets(s.TransitionSets, node.Event.Channel, tr)
}

// addTransition adds trans onto the transition set for channel.
//
// If no such transition set exists, one is created.

// addVerdictsFromNode adds the test verdicts from n into s.
func (s *State) addVerdictsFromNode(node *testlang.Node) {
	s.Verdict.Add(node.Outcome, node.Tests...)
}

func (s *State) String() string {
	return fmt.Sprintf("%s:{%s}(%s)", s.ID, s.Verdict, s.transitionString())
}

func (s *State) transitionString() string {
	tsets := make([]string, len(s.TransitionSets))
	for i, t := range s.TransitionSets {
		tsets[i] = t.String()
	}

	tsetStr := strings.Join(tsets, ", ")

	return tsetStr
}
