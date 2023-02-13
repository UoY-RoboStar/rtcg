package stm

import (
	"fmt"
	"strings"

	"github.com/UoY-RoboStar/rtcg/internal/structure"
)

// Verdict contains information about a potential test verdict.
type Verdict struct {
	// IsObserved is true if this test verdict is active on the parent state for at least one test.
	//
	// The test itself may not be in Tests (for example, if the test tree isn't tracking which node came from which
	// test).
	IsObserved bool `json:"is_observed"`

	// Tests is the set of names of tests that are known to be affected by this verdict.
	Tests structure.Set[string] `json:"tests,omitempty"`
}

// NewVerdict constructs a new Verdict.
func NewVerdict() *Verdict {
	return &Verdict{Tests: structure.Set[string]{}, IsObserved: false}
}

// Add adds each test in tests into the affected list for Verdict v, and sets it to active.
func (v *Verdict) Add(tests ...string) {
	v.IsObserved = true
	for _, t := range tests {
		v.Tests.Add(t)
	}
}

func (v *Verdict) String() string {
	if !v.IsObserved {
		return "none"
	}

	return fmt.Sprintf("[%s]", strings.Join(v.Tests.Values(), ", "))
}
