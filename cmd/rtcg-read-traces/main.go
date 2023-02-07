// Program rtcg-read-trace reads unfactorised trace and converts them to a
// JSON test.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"rtcg/internal/cli"
	"rtcg/internal/testlang"
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

	suite := trace.ExpandAll(traces)
	return dumpSuite(suite)
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

func dumpSuite(suite testlang.Suite) error {
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "\t")
	return w.Encode(suite)
}
