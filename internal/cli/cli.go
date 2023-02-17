// Package cli contains various helpers for the command-line interface of rtcg tools.
package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

// stdinPath is the conventional path used to request input from stdin.
const stdinPath = "-"

// ParseFileArgument parses a common pattern where the final argument is optional and is the input file.
//
// If the final argument is missing (judged by comparing length of args against maxArgs), it becomes stdinPath.
func ParseFileArgument(args []string, maxArgs int) (string, error) {
	switch len(args) {
	case maxArgs:
		// Last argument is the file.
		return args[maxArgs-1], nil
	case maxArgs - 1:
		// Missing file argument, try stdin.
		return stdinPath, nil
	default:
		return "", ErrBadArgs
	}
}

// OpenFileOrStdin opens the file at path or, if it is '-', returns stdin instead.
func OpenFileOrStdin(path string) (io.ReadCloser, error) {
	if path == stdinPath {
		return os.Stdin, nil
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("when opening file %q: %w", path, err)
	}

	return file, nil
}

// HandleError handles a top-level error err in the tool.
//
// It exits with a nonzero status code if the error is non-nil.
//
// If the error has something to do with arguments, we call usage.
func HandleError(err error, usage func()) {
	if err == nil {
		return
	}

	_, _ = fmt.Fprintln(os.Stderr, err)

	if errors.Is(err, ErrBadArgs) {
		usage()
	}

	os.Exit(1)
}

// Usage is a basic usage function.
func Usage(usageStr string) {
	_, _ = fmt.Fprintln(os.Stderr, "usage:", os.Args[0], usageStr)
}

// FlagUsage is a usage function for commands that use flag.
func FlagUsage(usageStr string) {
	Usage("(flags) " + usageStr)
	flag.PrintDefaults()
}

// ErrBadArgs is an error used when invalid arguments have been supplied.
var ErrBadArgs = errors.New("invalid arguments supplied")
