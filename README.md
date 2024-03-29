# `rtcg`

**NOTE (2023-05-02)**: The main developer of `rtcg` (Matt) is no longer
employed on the project for which `rtcg` was originally developed.  As such,
development efforts on `rtcg` may be intermittent.
See [the corresponding GitHub issue](https://github.com/UoY-RoboStar/rtcg/issues/12)
for details.


`rtcg` (RoboStar test code generation) is a series of tools for manipulating
forbidden-trace tests, and using them to generate C code.

`rtcg` is licenced under the MIT licence; see `LICENSE.md`.

For more in-depth documentation, try the [wiki](https://github.com/UoY-RoboStar/rtcg/wiki).
We welcome all contributions to the tool and its documentation.


## How to build

Building `rtcg` requires [Go](https://go.dev) 1.20.
All other dependencies are handled automatically using the `go` tool.

To build all tools in one go, type `make`.  The tools will appear in `bin`.


### Regenerating `go generate` code

This project makes use of the following `go generate` tools.  Usually, you
won't need them installed, as we check in the results of the last time we
applied them.  However, you may want to update the generated code after making
changes to the source:

- [go-enum](https://github.com/abice/go-enum)


## The tools

### `rtcg.sh`: generate C tests from traces

Usage: `rtcg-gen CONFIG-FILE [TRACES-FILE]`

This `sh` script automates a pipeline of the following steps:

1. `rtcg-read-traces` on `TRACES-FILE`
2. `rtcg-make-stms` on the output of step 1
3. `rtcg-gen` on `CONFIG-FILE` and the output of step 2

95% of the time, this is the correct workflow.


### `rtcg-read-traces`: convert traces to tests

Usage: `rtcg-read-traces [TRACE-FILE]`

This tool takes in a series of forbidden traces, one per line, and produces
a `rtcg` test suite in JSON format on stdout.


### `rtcg-make-stms`: convert tests to state mahines

Usage: `rtcg-make-stms [TEST-FILE]`

This tool takes in a test JSON file (e.g. from `rtcg-read-traces`) and emits
a JSON file describing test state machines.  The distinction between the two
is that the former is a faithful tree representation of the test in the CSP
testing theory, and the latter is a more straightforwardly useful list of
states and transitions which can be used to generate code.


### `rtcg-gen`: generate C code

Usage: `rtcg-gen [-clean | -output DIR] CONFIG-FILE [STM-FILE]`

**Tool under development**

This tool takes in a `rtcg` state machine suite in JSON format (e.g. from
`rtcg-make-stms`), as well as an configuration file telling it which templates
and other auxiliary data to use, and emits automatically generated test code.

The configuration file format is not yet documented, but you can see examples
in `examples`: for instance, `examples/bmon/gen.xml`.

To get a quick idea of how `rtcg-gen` works, try:

```shell
$ make examples-cpp
$ ls out/bmon/ros/src
```


## Q&A

### Is this ready for production use?

Not yet!  This is a prototype being explored as part of some work on
testing RoboChart models, and should not be considered a main-line RoboStar
tool in the same way as RoboTool.


### Why is this written in Go, given RoboTool is an Eclipse plugin set?

I (Matt) chose to write this tool in Go because its main use is to expand out
a series of textual templates in a data-driven manner, and this is a very good
use case for Go's standard `text/template` library.


### Will there be binaries available?

Once the tool is in a stable state, I'll start producing binaries for the
various platforms supported by RoboTool.

