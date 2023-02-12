// Package stm contains the testing state machine.
//
// Compared to the testing tree representation, a state machine is slightly better organised for emission as code.
// For instance:
//
// - all testing states are linearised into one slice for easy body emission;
// - testing states have names, with transitions occurring on a 'jump to state with this name' basis;
// - information about which tests have been failed, or are about to pass, is centralised in each state.
package stm

import "github.com/UoY-RoboStar/rtcg/internal/structure"

// Suite is a test suite, in state machine format.
type Suite map[string]*Stm

// Stm is a state machine.
type Stm struct {
	// States is the list of states in this state machine.
	//
	// Conventionally, the first state in the machine is the initial state.
	States []*State

	// Tests is the set of names of tests being captured by this state machine.
	Tests structure.Set[string]
}
