// Package verdict contains types for managing verdicts on tests.
//
// Verdict in rtcg map tests to outcomes.
package verdict

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/UoY-RoboStar/rtcg/internal/structure"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

// Verdict maps test names to their outcomes.
type Verdict map[string]testlang.Outcome

// Tests gets a sorted list of all test names covered by this Verdict.
func (v *Verdict) Tests() []string {
	tests := make([]string, 0, len(*v))

	for k := range *v {
		tests = append(tests, k)
	}

	sort.Strings(tests)

	return tests
}

// IsInc gets whether there are any inconclusive tests in this Verdict.
func (v *Verdict) IsInc() bool {
	return v.Is(testlang.OutcomeInc)
}

// Inc is shorthand for getting the inconclusive verdicts from this Verdict.
func (v *Verdict) Inc() TestSet {
	return v.SetOf(testlang.OutcomeInc)
}

// IsPass gets whether there are any passing tests in this Verdict.
func (v *Verdict) IsPass() bool {
	return v.Is(testlang.OutcomePass)
}

// Pass is shorthand for getting the passing verdicts from this Verdict.
func (v *Verdict) Pass() TestSet {
	return v.SetOf(testlang.OutcomePass)
}

// IsFail gets whether there are any passing tests in this Verdict.
func (v *Verdict) IsFail() bool {
	return v.Is(testlang.OutcomeFail)
}

// Fail is shorthand for getting the failing verdicts from this Verdict.
func (v *Verdict) Fail() TestSet {
	return v.SetOf(testlang.OutcomeFail)
}

// Is gets whether there are any tests in this Verdict with the given status.
func (v *Verdict) Is(status testlang.Outcome) bool {
	for _, s := range *v {
		if s == status {
			return true
		}
	}

	return false
}

// SetOf gets the set of all tests that match status.
func (v *Verdict) SetOf(status testlang.Outcome) TestSet {
	tests := structure.NewSet[string]()

	for k, s := range *v {
		if s == status {
			tests.Add(k)
		}
	}

	return TestSet(tests)
}

// Add adds each test in tests as having outcome.
func (v *Verdict) Add(outcome testlang.Outcome, tests ...string) {
	for _, t := range tests {
		(*v)[t] = outcome
	}
}

func (v *Verdict) String() string {
	tests := v.Tests()

	vstrings := make([]string, len(tests))

	for i, t := range tests {
		vstrings[i] = fmt.Sprintf("%s: %s", t, (*v)[t])
	}

	return strings.Join(vstrings, ", ")
}

// MarshalJSON marshals a Verdict by first changing each status to its string.
func (v *Verdict) MarshalJSON() ([]byte, error) {
	strMap := make(map[string]string, len(*v))

	for k, v := range *v {
		strMap[k] = v.String()
	}

	bs, err := json.Marshal(strMap)
	if err != nil {
		return nil, fmt.Errorf("couldn't marshal verdict set: %w", err)
	}

	return bs, nil
}

// UnmarshalJSON marshals a Verdict by first parsing each status from its string.
func (v *Verdict) UnmarshalJSON(bs []byte) error {
	var strMap map[string]string

	if err := json.Unmarshal(bs, &strMap); err != nil {
		return fmt.Errorf("couldn't unmarshal verdict set: %w", err)
	}

	*v = make(Verdict, len(strMap))

	for test, outStr := range strMap {
		var outcome testlang.Outcome

		if err := outcome.UnmarshalText([]byte(outStr)); err != nil {
			return fmt.Errorf("couldn't unmarshal status %q: %w", outStr, err)
		}

		v.Add(outcome, test)
	}

	return nil
}

// TestSet contains information about a potential test verdict.
type TestSet structure.Set[string]
