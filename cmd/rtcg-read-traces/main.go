// Program rtcg-read-trace reads unfactorised trace and converts them to a
// JSON test.
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/UoY-RoboStar/rtcg/internal/cli"
	"github.com/UoY-RoboStar/rtcg/internal/trace"
)

const (
	usage            = "[INPUT-FILE]"
	numAnonymousArgs = 1
)

func main() {
	cli.HandleError(run(), usage)
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

	if err := suite.Write(os.Stdout); err != nil {
		return fmt.Errorf("couldn't write expanded traces: %w", err)
	}

	return nil
}

func readTraces(fname string) ([]trace.Trace, error) {
	file, err := cli.OpenFileOrStdin(fname)
	if err != nil {
		return nil, fmt.Errorf("couldn't open input: %w", err)
	}

	traces, err := trace.Read(file)

	return traces, errors.Join(err, file.Close())
}

func parseArgs(args []string) (string, error) {
	path, err := cli.ParseFileArgument(args, numAnonymousArgs+1)
	if err != nil {
		return "", fmt.Errorf("couldn't parse args: %w", err)
	}

	return path, nil
}
