package cpp_test

import (
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	cfg "github.com/UoY-RoboStar/rtcg/internal/gen/config/cpp"
	"github.com/UoY-RoboStar/rtcg/internal/gen/cpp"
	"github.com/UoY-RoboStar/rtcg/internal/gen/gencommon"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

// TestOnSuite_Dirs tests the Dirs method of OnSuite.
func TestOnSuite_Dirs(t *testing.T) {
	t.Parallel()

	rtcgConvertDir := filepath.Join("out", "src", "convert")
	rtcgSrcDir := filepath.Join("out", "src", "rtcg")

	for name, input := range map[string]struct {
		config *cfg.Config
		want   []string
	}{
		"empty-animate": {
			config: cfg.New(cfg.VariantAnimate),
			want:   []string{rtcgSrcDir},
		},
		"empty-animate-convert": {
			config: cfg.New(cfg.VariantAnimate, cfg.WithChannel("foo", "bar")),
			want:   []string{rtcgConvertDir, rtcgSrcDir},
		},
		// TODO: add tests for non-empty cases
	} {
		config := input.config
		want := input.want

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			dirs := gencommon.DirSet{Input: "in", Output: "out"}

			gen, err := cpp.New(config, dirs)
			if err != nil {
				t.Fatalf("unexpected error creating generator: %s", err)
			}

			var suite stm.Suite
			onSuite := gen.OnSuite(&suite)

			got := onSuite.Dirs()

			sort.Strings(got)
			sort.Strings(want)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("unexpected directories: got %v, want %v", got, want)
			}
		})
	}
}
