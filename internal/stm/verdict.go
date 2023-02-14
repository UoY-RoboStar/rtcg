package stm

import (
	"encoding/json"
	"fmt"

	"github.com/UoY-RoboStar/rtcg/internal/structure"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

type VerdictSet map[testlang.Status]*Verdict

// IsInc gets whether there are any inconclusive tests in this verdict set.
func (v *VerdictSet) IsInc() bool {
	return v.Is(testlang.StatusInc)
}

// Inc is shorthand for getting the inconclusive verdicts from this verdict set.
func (v *VerdictSet) Inc() *Verdict {
	return (*v)[testlang.StatusInc]
}

// IsPass gets whether there are any passing tests in this verdict set.
func (v *VerdictSet) IsPass() bool {
	return v.Is(testlang.StatusPass)
}

// Pass is shorthand for getting the passing verdicts from this verdict set.
func (v *VerdictSet) Pass() *Verdict {
	return (*v)[testlang.StatusPass]
}

// IsFail gets whether there are any passing tests in this verdict set.
func (v *VerdictSet) IsFail() bool {
	return v.Is(testlang.StatusFail)
}

// Is gets whether there are any tests in this verdict set with the given status.
func (v *VerdictSet) Is(status testlang.Status) bool {
	_, ok := (*v)[status]

	return ok
}

// Fail is shorthand for getting the failing verdicts from this verdict set.
func (v *VerdictSet) Fail() *Verdict {
	return (*v)[testlang.StatusFail]
}

// Add adds each test in tests into the affected list for Verdict v, and sets it to active.
func (v *VerdictSet) Add(status testlang.Status, tests ...string) {
	if s, ok := (*v)[status]; ok {
		s.Tests.Add(tests...)
	} else {
		(*v)[status] = NewVerdict(tests...)
	}
}

// MarshalJSON marshals a VerdictSet by first changing each status to its string.
func (v *VerdictSet) MarshalJSON() ([]byte, error) {
	strMap := make(map[string]*Verdict, len(*v))

	for k, v := range *v {
		strMap[k.String()] = v
	}

	bs, err := json.Marshal(strMap)
	if err != nil {
		return nil, fmt.Errorf("couldn't marshal verdict set: %w", err)
	}

	return bs, nil
}

// UnmarshalJSON marshals a VerdictSet by first parsing each status from its string.
func (v *VerdictSet) UnmarshalJSON(bs []byte) error {
	var strMap map[string]*Verdict

	if err := json.Unmarshal(bs, &strMap); err != nil {
		return fmt.Errorf("couldn't unmarshal verdict set: %w", err)
	}

	*v = make(VerdictSet, len(strMap))

	for k, ver := range strMap {
		var status testlang.Status

		if err := status.UnmarshalText([]byte(k)); err != nil {
			return fmt.Errorf("couldn't unmarshal status %q: %w", k, err)
		}

		(*v)[status] = ver
	}

	return nil
}

// Verdict contains information about a potential test verdict.
type Verdict struct {
	// Tests is the set of names of tests that are known to be affected by this verdict.
	Tests structure.Set[string] `json:"tests,omitempty"`
}

// NewVerdict constructs a new Verdict with the given tests.
func NewVerdict(tests ...string) *Verdict {
	return &Verdict{Tests: structure.NewSet[string](tests...)}
}

func (v *Verdict) String() string {
	return v.Tests.String()
}
