package testlang_test

import (
	"reflect"
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"github.com/UoY-RoboStar/rtcg/internal/testlang/value"
)

// TestPass tests that the 'pass' constructor works.
func TestPass(t *testing.T) {
	t.Parallel()

	event := testlang.Output("batteryStatus", value.Enum("ok"))

	got := testlang.Pass(event)
	want := testlang.Node{
		ID:      "",
		Tests:   nil,
		Outcome: testlang.OutcomePass,
		Event:   &event,
		Next:    nil,
	}

	checkNode(t, got, want)
}

// TestInc tests that the 'inconclusive' constructor works.
func TestInc(t *testing.T) {
	t.Parallel()

	event := testlang.Output("batteryStatus", value.Enum("ok"))

	got := testlang.Inc(event)
	want := testlang.Node{
		ID:      "",
		Tests:   nil,
		Outcome: testlang.OutcomeInc,
		Event:   &event,
		Next:    nil,
	}

	checkNode(t, got, want)
}

// TestFail tests that the 'fail' constructor works.
func TestFail(t *testing.T) {
	t.Parallel()

	got := testlang.Fail()
	want := testlang.Node{
		ID:      "",
		Tests:   nil,
		Outcome: testlang.OutcomeFail,
		Event:   nil,
		Next:    nil,
	}

	checkNode(t, got, want)
}

func checkNode(t *testing.T, got, want testlang.Node) {
	t.Helper()

	// TODO: check other fields

	if got.Outcome != want.Outcome {
		t.Errorf("want node status %q, got %q", &want.Outcome, &got.Outcome)
	}

	if !reflect.DeepEqual(got.Event, want.Event) {
		t.Errorf("want node event %q, got %q", want.Event, got.Event)
	}

	if !reflect.DeepEqual(got.Next, want.Next) {
		t.Errorf("want node successors %v, got %v", want.Next, got.Next)
	}
}
