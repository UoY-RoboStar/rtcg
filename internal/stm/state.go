package stm

import (
	"fmt"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"strings"
)

// State is a state machine state (with attached transitions).
type State struct {
	ID             testlang.NodeID // ID is the identifier of the State.
	TransitionSets []TransitionSet // TransitionSets is the list of transition sets out of this state.

	Verdicts [testlang.NumStatus]*Verdict // Verdict gives information about which test verdicts this state reports.
}

// NewState creates a new State with the given id.
func NewState(id testlang.NodeID) *State {
	s := State{ID: id}
	for _, st := range testlang.AllStatuses {
		s.Verdicts[st] = NewVerdict()
	}
	return &s
}

// Inc is shorthand for getting the inconclusive verdict information.
func (s *State) Inc() *Verdict {
	return s.Verdicts[testlang.StatInc]
}

// Pass is shorthand for getting the passing verdict information.
func (s *State) Pass() *Verdict {
	return s.Verdicts[testlang.StatPass]
}

// Fail is shorthand for getting the failing verdict information.
func (s *State) Fail() *Verdict {
	return s.Verdicts[testlang.StatFail]
}

// AddTransitionToNode adds a transition from this node to the test-tree node n.
//
// We assume the node has already been assigned an ID.
func (s *State) AddTransitionToNode(n *testlang.Node) {
	tr := Transition{Value: n.Event.Value, Next: n.ID}
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

// AddVerdictsFromNode adds the test verdicts from n into s.
func (s *State) AddVerdictsFromNode(n testlang.Node) {
	for _, st := range testlang.AllStatuses {
		if n.Status == st {
			s.Verdicts[st].Add(n.Tests...)
			return
		}
	}
}

func (s *State) String() string {
	vsets := make([]string, testlang.NumStatus)
	for _, st := range testlang.AllStatuses {
		vsets[st] = fmt.Sprintf("%s %s", &st, s.Verdicts[st])
	}
	vsetStr := strings.Join(vsets, ", ")

	tsets := make([]string, len(s.TransitionSets))
	for i, t := range s.TransitionSets {
		tsets[i] = t.String()
	}
	tsetStr := strings.Join(tsets, ", ")

	return fmt.Sprintf("%s:{%s}(%s)", s.ID, vsetStr, tsetStr)
}
