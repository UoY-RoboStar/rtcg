// Program rtcg-make-stms generates a testing state machine given a test tree.
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/UoY-RoboStar/rtcg/internal/cli"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"github.com/UoY-RoboStar/rtcg/internal/validate"
)

const numAnonymousArgs = 1

func main() {
	cli.HandleError(run(), func() { cli.Usage("[INPUT-FILE]") })
}

func run() error {
	a, err := parseArgs(os.Args)
	if err != nil {
		return err
	}

	return a.run()
}

func parseArgs(argv []string) (*makeStmAction, error) {
	var (
		action makeStmAction
		err    error
	)

	action.inputFile, err = cli.ParseFileArgument(argv, numAnonymousArgs+1)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse args: %w", err)
	}

	return &action, nil
}

type makeStmAction struct {
	inputFile string
}

func (a *makeStmAction) run() error {
	s, err := a.readSuite()
	if err != nil {
		return fmt.Errorf("couldn't read test suite: %w", err)
	}

	vs, err := validate.FullSuite(s)
	if err != nil {
		return fmt.Errorf("malformed test suite: %w", err)
	}

	stms, err := a.buildStms(vs)
	if err != nil {
		return fmt.Errorf("couldn't build stms for test suite: %w", err)
	}

	if err := stms.Write(os.Stdout); err != nil {
		return fmt.Errorf("couldn't write state machines: %w", err)
	}

	return nil
}

func (a *makeStmAction) readSuite() (testlang.Suite, error) {
	file, err := cli.OpenFileOrStdin(a.inputFile)
	if err != nil {
		return nil, fmt.Errorf("couldn't open input: %w", err)
	}

	suite, err := testlang.ReadSuite(file)

	return suite, errors.Join(err, file.Close())
}

func (a *makeStmAction) buildStms(tests validate.Suite) (stm.Suite, error) {
	var bs stm.Builder

	suite, err := bs.BuildSuite(tests)
	if err != nil {
		return nil, fmt.Errorf("couldn't build state machines: %w", err)
	}

	return suite, nil
}
