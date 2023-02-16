package stm_test

import (
	"reflect"
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/stm"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"github.com/UoY-RoboStar/rtcg/internal/trace"
)

// Tests the building of a state machine relating to a tree with only one (failing) event.
func TestBuilder_Build_EmptyPrefixTrace(t *testing.T) {
	t.Parallel()

	event := testlang.Output("foo", testlang.Int(42))

	tree := trace.New(event).Expand("test")

	var builder stm.Builder
	mach := builder.Build("tree", tree)

	gotTests := mach.Tests.Values()
	wantTests := []string{"test"}

	if !reflect.DeepEqual(gotTests, wantTests) {
		t.Errorf("incorrect tests: got %v, wanted %v", gotTests, wantTests)
	}

	emptyPrefixTraceStates(t, mach)
}

func emptyPrefixTraceStates(t *testing.T, mach stm.Stm) {
	t.Helper()

	if len(mach.States) != 2 {
		t.Errorf("incorrect number of states: got %d (%v), wanted 2", len(mach.States), mach.States)

		return
	}

	emptyPrefixTraceNode1(t, mach.States[0])

	emptyPrefixTraceNode2(t, mach.States[1])
}

func emptyPrefixTraceNode1(t *testing.T, state *stm.State) {
	t.Helper()

	if state.ID != "tree" {
		t.Errorf("expected first state to have name of test-tree, got %q", state.ID)
	}

	assertVerdict(t, state.Verdicts, testlang.OutcomeInc, false)
	assertVerdict(t, state.Verdicts, testlang.OutcomeFail, false)
	assertVerdict(t, state.Verdicts, testlang.OutcomePass, true)
}

func emptyPrefixTraceNode2(t *testing.T, state *stm.State) {
	t.Helper()

	if state.ID != "node0" {
		t.Errorf("expected second state to have generated name, got %q", state.ID)
	}

	nsets := len(state.TransitionSets)
	if nsets != 0 {
		t.Errorf("second state should have no transition sets, got %d", nsets)
	}

	assertVerdict(t, state.Verdicts, testlang.OutcomeInc, false)
	assertVerdict(t, state.Verdicts, testlang.OutcomeFail, true)
	assertVerdict(t, state.Verdicts, testlang.OutcomePass, false)
}

func assertVerdict(t *testing.T, verdicts *stm.VerdictSet, status testlang.Outcome, want bool) {
	t.Helper()

	if verdicts.Is(status) != want {
		t.Errorf("state should%s be %s", verdictFailSuffix(want), &status)
	}
}

func verdictFailSuffix(exist bool) string {
	if exist {
		return "n't"
	}

	return ""
}
