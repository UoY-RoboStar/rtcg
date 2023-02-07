# `rtcg`

`rtcg` (RoboStar test code generation) is a series of tools for manipulating
forbidden-trace tests, and using them to generate C code.

`rtcg` is licenced under the MIT licence; see `LICENSE.md`.


## Requirements

Building `rtcg` requires [Go](https://go.dev) 1.20.
All other dependencies are handled automatically using the `go` tool.

To build all tools in one go, type `make`.  The tools will appear in `bin`.


## The tools


### `rtcg-read-traces`: convert traces to tests

Usage: `rtcg-read-traces [INPUT-FILE]`

This tool takes in a series of forbidden traces, one per line, and produces
a `rtcg` test suite in JSON format on stdout.


### `rtcg-gen`: generate C code

Usage: `rtcg-gen TEMPLATE-DIR [INPUT-FILE]`

**Tool under development**

This tool takes in a `rtcg` test suite in JSON format (such as those produced
by `rtcg-read-traces`), as well as some `text/template` templates for
generating code, and emits automatically generated test code.