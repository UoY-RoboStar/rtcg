package gen

import (
	"time"

	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

// Context is the context passed into the template.
type Context struct {
	// Name is the name of the test case being generated.
	Name string
	// Date is the time of generation.
	Date time.Time
	// Stm is the state machine being generated.
	Stm *stm.Stm
}
