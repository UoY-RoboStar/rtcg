package cpp_test

import (
	"path/filepath"
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/gen/config/cpp"
)

// TestVariant_Dir tests the Dir method on CppVariant.
func TestVariant_Dir(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		variant cpp.Variant
		dir     string
		want    string
	}{
		"animate-empty": {variant: cpp.VariantAnimate, dir: "", want: "animate"},
		"animate-root":  {variant: cpp.VariantAnimate, dir: "/", want: "/animate"},
		"ros-empty":     {variant: cpp.VariantRos, dir: "", want: "ros"},
		"ros-root":      {variant: cpp.VariantRos, dir: "/", want: "/ros"},
	}
	for name, test := range tests {
		name := name
		test := test

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := test.variant.Dir(filepath.FromSlash(test.dir))
			want := filepath.FromSlash(test.want)

			if got != want {
				t.Errorf("VariantDir %s: got %v, want %v", name, got, want)
			}
		})
	}
}
