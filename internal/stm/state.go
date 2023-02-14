package stm

import (
	"fmt"
	"strings"

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
	TransitionSets []TransitionSet `json:"transitionSets,omitempty"`

	Verdicts *VerdictSet `json:"verdicts,omitempty"` // Verdicts holds the test verdicts that this state reports.
}

// NewState creates a new State with the given id.
func NewState(id testlang.NodeID) *State {
	return &State{
		ID:             id,
		TransitionSets: nil,
		Verdicts:       &VerdictSet{},
	}
}

// AddOutgoingNode handles all bookkeeping for logging verdicts and producing transitions from s to the state for node.
func (s *State) AddOutgoingNode(node *testlang.Node) {
	s.addVerdictsFromNode(node)
	s.addTransitionToNode(node)
}

// AddTransitionToNode adds a transition from this state to the given test-tree node.
//
// We assume the node has already been assigned an ID.
func (s *State) addTransitionToNode(node *testlang.Node) {
	// We don't add transitions to failing nodes; they are just sentinels with no test content.
	if node.Status == testlang.StatusFail {
		return
	}

	if node.ID == "" {
		panic("should have assigned an ID to node")
	}

	tr := Transition{Value: node.Event.Value, Next: node.ID}
	s.addTransition(node.Event.Channel, tr)
}

// AddTransition adds transition onto the transition set for channel.
//
// If no such transition set exists, one is created.
func (s *State) addTransition(channel testlang.Channel, transition Transition) {
	// Try merging onto an existing channel set.
	for i := range s.TransitionSets {
		ts := &s.TransitionSets[i]
		if ts.Channel.Equals(channel) {
			ts.Transitions = append(ts.Transitions, transition)

			return
		}
	}
	// No transition set with this channel exists yet.
	s.TransitionSets = append(s.TransitionSets, TransitionSet{
		Channel: channel, Transitions: []Transition{transition},
	})
}

// AddVerdictsFromNode adds the test verdicts from n into s.
func (s *State) addVerdictsFromNode(node *testlang.Node) {
	s.Verdicts.Add(node.Status, node.Tests...)
}

func (s *State) String() string {
	return fmt.Sprintf("%s:{%s}(%s)", s.ID, s.verdictString(), s.transitionString())
}

func (s *State) verdictString() string {
	vsets := make([]string, testlang.NumStatus)
	for st := testlang.FirstStatus; st <= testlang.LastStatus; st++ {
		vsets[st] = fmt.Sprintf("%s %s", &st, (*s.Verdicts)[st])
	}

	return strings.Join(vsets, ", ")
}

func (s *State) transitionString() string {
	tsets := make([]string, len(s.TransitionSets))
	for i, t := range s.TransitionSets {
		tsets[i] = t.String()
	}

	tsetStr := strings.Join(tsets, ", ")

	return tsetStr
}
