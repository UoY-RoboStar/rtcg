package stm_test

import (
	"github.com/UoY-RoboStar/rtcg/internal/stm"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"github.com/UoY-RoboStar/rtcg/internal/trace"
	"reflect"
	"testing"
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

	if state.Verdicts.IsInc() {
		t.Error("first state should not be inconclusive")
	}
	if state.Verdicts.IsFail() {
		t.Error("first state should not be failing")
	}
	if !state.Verdicts.IsPass() {
		t.Error("first state should be passing")
	}
}

func emptyPrefixTraceNode2(t *testing.T, state *stm.State) {
	t.Helper()

	if state.ID != "node_0" {
		t.Errorf("expected second state to have generated name, got %q", state.ID)
	}

	nsets := len(state.TransitionSets)
	if nsets != 0 {
		t.Errorf("second state should have no transition sets, got %d", nsets)
	}

	if state.Verdicts.IsInc() {
		t.Error("first state should not be inconclusive")
	}
	if !state.Verdicts.IsFail() {
		t.Error("first state should be failing")
	}
	if state.Verdicts.IsPass() {
		t.Error("first state should not be passing")
	}
}
