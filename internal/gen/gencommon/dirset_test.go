package gencommon_test

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/gen/gencommon"
)

func TestDirSet_OutputPath(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input gencommon.DirSet
		want  string
	}{
		"empty": {
			input: gencommon.DirSet{Input: "", Output: ""},
			want:  "src",
		},
		"out-empty": {
			input: gencommon.DirSet{Input: "foo", Output: ""},
			want:  "src",
		},
		"non-empty": {
			input: gencommon.DirSet{Input: "foo", Output: "bar"},
			want:  filepath.Join("bar", "src"),
		},
	}
	for name, test := range tests {
		name := name
		test := test

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := test.input.OutputPath("src"); got != test.want {
				t.Errorf("OutputPath(\"src\") of %s: got %v, want %v", name, got, test.want)
			}
		})
	}
}

func TestDirSet_Subdir(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input gencommon.DirSet
		dir   string
		want  gencommon.DirSet
	}{
		"both-empty": {
			input: gencommon.DirSet{Input: "", Output: ""},
			dir:   "",
			want:  gencommon.DirSet{Input: "", Output: ""},
		},
		"inputs-empty": {
			input: gencommon.DirSet{Input: "", Output: ""},
			dir:   "foo",
			want:  gencommon.DirSet{Input: "foo", Output: "foo"},
		},
		"dir-empty": {
			input: gencommon.DirSet{Input: "in", Output: "out"},
			dir:   "",
			want:  gencommon.DirSet{Input: "in", Output: "out"},
		},
		"both-present": {
			input: gencommon.DirSet{Input: "in", Output: "out"},
			dir:   "foo",
			want:  gencommon.DirSet{Input: filepath.Join("in", "foo"), Output: filepath.Join("out", "foo")},
		},
	}

	for name, test := range tests {
		name := name
		test := test

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := test.input.Subdir(test.dir); !reflect.DeepEqual(got, test.want) {
				t.Errorf("%s: got %v, want %v", name, got, test.want)
			}
		})
	}
}
