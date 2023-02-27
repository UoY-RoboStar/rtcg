package validate

import (
	"fmt"

	"github.com/UoY-RoboStar/rtcg/internal/structure"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

// Suite is a validated test suite.
type Suite map[string]*Test

// FullSuite performs a full validation of the given test suite.
func FullSuite(suite testlang.Suite) (Suite, error) {
	validated, err := structure.TryOverMapValues(suite, Full)
	if err != nil {
		return nil, fmt.Errorf("couldn't validate suite: %w", err)
	}

	return validated, nil
}
