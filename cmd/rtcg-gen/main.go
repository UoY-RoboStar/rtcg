// Program rtcg-gen performs test code generation given JSON input.
package main

import (
	"errors"
	"fmt"
	"os"
	"rtcg/internal/cli"
	"rtcg/internal/stm"
	"rtcg/internal/testlang"
)

func main() {
	cli.HandleError(run(), "TEMPLATE-DIR [INPUT-FILE]")
}

func run() error {
	args, err := parseArgs(os.Args)
	if err != nil {
		return err
	}

	s, err := readSuite(args.inputFile)
	if err != nil {
		return fmt.Errorf("couldn't read test suite: %w", err)
	}

	var sb stm.Builder
	stms := sb.BuildSuite(s)

	for k, v := range stms {
		fmt.Println(k, "=", v)
	}

	return nil
}

func parseArgs(argv []string) (*args, error) {
	var args args

	argc := len(argv)
	if argc == 2 {
		args.inputFile = "-" // stdin
	} else if argc == 3 {
		args.inputFile = argv[2]
	} else {
		return nil, cli.ErrBadArgs
	}

	args.templateDir = argv[1]
	return &args, nil
}

type args struct {
	templateDir string
	inputFile   string
}

func readSuite(path string) (testlang.Suite, error) {
	f, err := cli.OpenFileOrStdin(path)
	if err != nil {
		return nil, err
	}
	suite, err := testlang.ReadSuite(f)
	return suite, errors.Join(err, f.Close())
}
