package stm

import (
	"fmt"
	"strings"
)

// Verdict contains information about a potential test verdict.
type Verdict struct {
	// IsObserved is true if this test verdict is active on the parent state for at least one test.
	//
	// The test itself may not be in Tests (for example, if the test tree isn't tracking which node came from which
	// test).
	IsObserved bool

	// Tests is the set of names of tests that are known to be affected by this verdict.
	Tests map[string]bool
}

// NewVerdict constructs a new Verdict.
func NewVerdict() *Verdict {
	return &Verdict{Tests: map[string]bool{}}
}

// Add adds each test in tests into the affected list for Verdict v, and sets it to active.
func (v *Verdict) Add(tests ...string) {
	v.IsObserved = true
	for _, t := range tests {
		v.Tests[t] = true
	}
}

// TestList get the list of affected tests from Verdict v.
func (v *Verdict) TestList() []string {
	tl := make([]string, 0, len(v.Tests))
	for t, b := range v.Tests {
		if b {
			tl = append(tl, t)
		}
	}
	return tl
}

func (v *Verdict) String() string {
	if !v.IsObserved {
		return "none"
	}
	return fmt.Sprintf("[%s]", strings.Join(v.TestList(), ", "))
}
