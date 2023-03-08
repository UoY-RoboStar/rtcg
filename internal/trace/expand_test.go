package trace_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"github.com/UoY-RoboStar/rtcg/internal/testlang/value"
	"github.com/UoY-RoboStar/rtcg/internal/trace"
)

// TestTrace_Expand tests expansion of traces to trees.
func TestTrace_Expand(t *testing.T) {
	t.Parallel()

	event1 := testlang.Input("foo", value.Int(42))
	event2 := testlang.Output("bar", value.Enum("baz"))

	for name, test := range map[string]struct {
		input trace.Forbidden
		want  testlang.Node
	}{
		"no-prefix": {
			input: trace.New().Forbid(event1),
			want:  testlang.Root(testlang.TestPoint(event1)),
		},
		"lone-prefix": {
			input: trace.New(event2).Forbid(event1),
			want: testlang.Root(testlang.Inc(
				event2,
				testlang.TestPoint(event1))),
		},
	} {
		name := name
		input := test.input
		want := test.want
		testlang.MarkAll(&want, name)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := input.Expand(name)

			if got == nil {
				t.Fatal("got nil")
			}

			if !reflect.DeepEqual(*got, want) {
				t.Fatalf("got %v, want %v", *got, want)
			}
		})
	}
}

// FuzzExpandAllCollapseAllRoundTrip tests round-tripping trace.ExpandAll and trace.CollapseAll.
func FuzzExpandAllCollapseAllRoundTrip(f *testing.F) {
	f.Add("")
	f.Add("batteryStatus.out.Ok")
	f.Add("batteryInfo.in.{| percentage=BATTERY_MISSION_THRESHOLD |}\nbatteryStatus.out.Ok")
	f.Add("batteryInfo.in.{| percentage=BATTERY_MISSION_THRESHOLD |}, batteryStatus.out.Ok")

	f.Fuzz(func(t *testing.T, input string) {
		t.Parallel()

		r := strings.NewReader(input)

		traces, err := trace.Read(r)
		if err != nil {
			t.SkipNow()
		}

		want := trace.Name(traces)

		suite := trace.ExpandAll(want)

		got, err := trace.CollapseAll(suite)
		if err != nil {
			t.Fatalf("error collapsing traces: %s", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("round-trip failure: got %v, want %v", got, want)
		}
	})
}
