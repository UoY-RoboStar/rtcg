package gencommon

import (
	"fmt"

	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

// TestGenerator is the interface of things that can generate on a test-by-test basis.
type TestGenerator interface {
	// GenerateTest generates for the test with the given name and STM.
	// The directory set is that for the test itself.
	GenerateTest(DirSet, string, *stm.Stm) error
}

// GenerateTests runs the generator gen on each item in suite, passing subdirectories from root.
func GenerateTests(root DirSet, suite stm.Suite, gen TestGenerator) error {
	for name, test := range suite {
		dirs := root.Subdir(name)

		if err := gen.GenerateTest(dirs, name, test); err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}
	}

	return nil
}
