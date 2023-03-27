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
	for name, tt := range tests {
		name := name
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := tt.input.OutputPath("src"); got != tt.want {
				t.Errorf("OutputPath(\"src\") of %s: got %v, want %v", name, got, tt.want)
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

	for name, tt := range tests {
		name := name
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := tt.input.Subdir(tt.dir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s: got %v, want %v", name, got, tt.want)
			}
		})
	}
}
