package validate

import (
	"github.com/UoY-RoboStar/rtcg/internal/structure"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

// Suite is a validated test suite.
type Suite map[string]*Test

// FullSuite performs a full validation of the given test suite.
func FullSuite(suite testlang.Suite) (Suite, error) {
	return structure.TryOverMapValues(suite, Full)
}
