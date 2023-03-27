package gencommon

import (
	"fmt"

	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

// TestGenerator is the interface of things that can generate on a test-by-test basis.
type TestGenerator interface {
	// GenerateTest generates for the test with the given name and STM.
	GenerateTest(string, *stm.Stm) error
}

// GenerateTests runs the given generation function on each
func GenerateTests(suite stm.Suite, gen TestGenerator) error {
	for name, test := range suite {
		if err := gen.GenerateTest(name, test); err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}
	}

	return nil
}
