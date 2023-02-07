// Package cli contains various helpers for the command-line interface of rtcg tools.
package cli

import (
	"errors"
	"fmt"
	"os"
)

// HandleError handles a top-level error err in the tool.
//
// It exits with a nonzero status code if the error is non-nil.
func HandleError(err error, usageStr string) {
	if err == nil {
		return
	}
	_, _ = fmt.Fprintln(os.Stderr, err)
	if errors.Is(err, ErrBadArgs) {
		usage(usageStr)
	}
	os.Exit(1)
}

func usage(usageStr string) {
	_, _ = fmt.Fprintln(os.Stderr, "usage:", os.Args[0], usageStr)
}

// ErrBadArgs is an error used when invalid arguments have been supplied.
var ErrBadArgs = errors.New("invalid arguments supplied")
