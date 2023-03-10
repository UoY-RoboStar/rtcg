// Package stm contains the testing state machine.
//
// Compared to the testing tree representation, a state machine is slightly better organised for emission as code.
// For instance:
//
// - all testing states are linearised into one slice for easy body emission;
// - testing states have names, with transitions occurring on a 'jump to state with this name' basis;
// - information about which tests have been failed, or are about to pass, is centralised in each state.
package stm

import (
	"fmt"
	"io"

	"github.com/UoY-RoboStar/rtcg/internal/serial"
	"github.com/UoY-RoboStar/rtcg/internal/stm/transition"
	"github.com/UoY-RoboStar/rtcg/internal/structure"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"github.com/UoY-RoboStar/rtcg/internal/testlang/rstype"
)

// Suite is a test suite, in state machine format.
type Suite map[string]*Stm

// ReadSuite reads a state machine suite from JSON in reader r.
func ReadSuite(r io.Reader) (Suite, error) {
	var suite Suite

	if err := serial.ReadJSON(r, &suite); err != nil {
		return nil, fmt.Errorf("couldn't read state machine suite: %w", err)
	}

	return suite, nil
}

// Write pretty-prints a state machine suite, as JSON, into writer w.
func (s *Suite) Write(w io.Writer) error {
	if err := serial.WriteJSON(w, s); err != nil {
		return fmt.Errorf("couldn't write state machine suite: %w", err)
	}

	return nil
}

// UnifiedTypes returns a map of unified inferred types for each channel across each test in s.
func (s *Suite) UnifiedTypes() (TypeMap, error) {
	tmap := make(TypeMap)

	for _, test := range *s {
		for cName, cType := range test.Types {
			uType, err := rstype.Unify(cType, tmap[cName])
			if err != nil {
				return nil, fmt.Errorf("couldn't reconcile types for channel %s: %w", cName, err)
			}

			tmap[cName] = uType
		}
	}

	return tmap, nil
}

// Stm is a testing state machine.
//
// A state machine arranges the nodes of a test tree into a form that is easy to emit as test code:
// a list of states with lists of transitions to other states, and metadata about the test that is
// gleaned from the test tree during construction.
type Stm struct {
	// States is the list of states in this state machine.
	//
	// Conventionally, the first state in the machine is the initial state.
	States []*State `json:"states"`

	// Types maps each channel to its inferred type.
	Types TypeMap `json:"types"`

	// Tests is the set of names of tests being captured by this state machine.
	Tests structure.Set[string] `json:"tests"`
}

type TypeMap map[string]*rstype.RsType

// InitialState is the node ID of the initial state.
func (s *Stm) InitialState() testlang.NodeID {
	// TODO: do we need to guard against an empty state machine?
	return s.States[0].ID
}

// TransitionSets calculates all aggregate transition sets across the whole state machine.
func (s *Stm) TransitionSets() []transition.AggregateSet {
	var result []transition.AggregateSet

	for _, st := range s.States {
		for _, ts := range st.TransitionSets {
			result = transition.AddToAggregateSets(result, st.ID, ts)
		}
	}

	return result
}
