package trace_test

import (
	"reflect"
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"github.com/UoY-RoboStar/rtcg/internal/testlang/value"
	"github.com/UoY-RoboStar/rtcg/internal/trace"
)

// TestForbidden_String tests the String method of Forbidden.
func TestForbidden_String(t *testing.T) {
	t.Parallel()

	barOut := testlang.Output("bar", value.None())
	fooIn := testlang.Input("foo", value.None())
	bazOut := testlang.Output("baz", value.Int(2))

	tests := map[string]struct {
		input trace.Forbidden
		want  string
	}{
		"empty": {
			input: trace.New().Forbid(barOut),
			want:  "<>!bar.out",
		},
		"small": {
			input: trace.New(fooIn).Forbid(barOut),
			want:  "<foo.in>!bar.out",
		},
		"named": {
			input: trace.New(fooIn, bazOut, fooIn).ForbidWithName(barOut, "named"),
			want:  "named:<foo.in, baz.out.int!2, foo.in>!bar.out",
		},
	}
	for name, test := range tests {
		name := name
		test := test

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := test.input.String(); got != test.want {
				t.Errorf("%s: got %v, want %v", name, got, test.want)
			}
		})
	}
}

// TestName tests the Name function.
func TestName(t *testing.T) {
	t.Parallel()

	// TODO: table driven test

	barOut := testlang.Output("bar", value.None())
	fooIn := testlang.Input("foo", value.None())
	bazOut := testlang.Output("baz", value.Int(2))

	input := []trace.Forbidden{
		trace.New().Forbid(barOut),                                      // test0
		trace.New(fooIn).Forbid(barOut),                                 // test1
		trace.New(fooIn, bazOut, fooIn).ForbidWithName(barOut, "named"), // named
		trace.New(bazOut).Forbid(barOut),                                // test2
		trace.New(fooIn, bazOut).ForbidWithName(barOut, "named"),        // named0
		trace.New(bazOut).ForbidWithName(bazOut, "test"),                // test3
	}

	want := []string{"test0", "test1", "named", "test2", "named0", "test3"}

	got := trace.Name(input)

	if len(got) != len(want) {
		t.Fatalf("Name() changed number of traces: got %d, want %d", len(got), len(want))
	}

	for i, n := range want {
		wantT := input[i]
		gotT := got[n]

		if !reflect.DeepEqual(gotT, wantT) {
			t.Errorf("%s: got %v, want %v", n, gotT, wantT)
		}
	}
}
