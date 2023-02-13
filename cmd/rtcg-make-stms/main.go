// Program rtcg-make-stms generates a testing state machine given a test tree.
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/UoY-RoboStar/rtcg/internal/cli"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

const (
	usage            = "[INPUT-FILE]"
	numAnonymousArgs = 1
)

func main() {
	cli.HandleError(run(), usage)
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

	stms := a.buildStms(s)

	return stms.Write(os.Stdout)
}

func (a *makeStmAction) readSuite() (testlang.Suite, error) {
	file, err := cli.OpenFileOrStdin(a.inputFile)
	if err != nil {
		return nil, fmt.Errorf("couldn't open input: %w", err)
	}

	suite, err := testlang.ReadSuite(file)

	return suite, errors.Join(err, file.Close())
}

func (a *makeStmAction) buildStms(tests testlang.Suite) stm.Suite {
	var bs stm.Builder

	return bs.BuildSuite(tests)
}