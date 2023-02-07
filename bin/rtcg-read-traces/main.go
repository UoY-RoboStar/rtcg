// Program rtcg-read-trace reads unfactorised trace and converts them to a
// JSON test.
package main

import (
	"errors"
	"fmt"
	"os"
	"rtcg/internal/cli"
	"rtcg/internal/trace"
)

func main() {
	cli.HandleError(run(), "INPUT-FILE")
}

func run() error {
	fname, err := parseArgs(os.Args)
	if err != nil {
		return err
	}

	traces, err := readTraces(fname)
	if err != nil {
		return fmt.Errorf("couldn't read traces from file %q: %w", fname, err)
	}

	for _, tr := range traces {
		fmt.Println(tr)
	}

	return nil
}

func readTraces(fname string) ([]trace.Trace, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}

	traces, err := trace.Read(f)
	return traces, errors.Join(err, f.Close())
}

func parseArgs(args []string) (string, error) {
	if len(args) != 2 {
		return "", cli.ErrBadArgs
	}
	return args[1], nil
}
