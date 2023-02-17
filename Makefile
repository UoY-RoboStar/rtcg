.POSIX:

GO          = go
INPUT_DIR   = ./cmd
BIN_DIR     = bin
GEN         = $(BIN_DIR)/rtcg-gen
MAKE_STMS   = $(BIN_DIR)/rtcg-make-stms
READ_TRACES = $(BIN_DIR)/rtcg-read-traces

.PHONY: \
  all \
  rtcg-gen \
  rtcg-make-stms \
  rtcg-read-traces \
  examples \
  examples-cpp \
  bmon \
  bmon-cpp \


# Makes all commands.
all: \
  rtcg-gen \
  rtcg-make-stms \
  rtcg-read-traces

rtcg-gen:
	$(GO) build -o $(BIN_DIR)/ $(INPUT_DIR)/rtcg-gen

rtcg-make-stms:
	$(GO) build -o $(BIN_DIR)/ $(INPUT_DIR)/rtcg-make-stms

rtcg-read-traces:
	$(GO) build -o $(BIN_DIR)/ $(INPUT_DIR)/rtcg-read-traces

#
# Examples
#

# Makes all examples.
examples: \
  bmon

bmon: examples/bmon/stms.json examples/bmon/tests.json examples/bmon/traces

examples/bmon/stms.json: examples/bmon/tests.json
	$(MAKE_STMS) $< > $@

examples/bmon/tests.json: examples/bmon/traces
	$(READ_TRACES) $< > $@

#
# C++ code for examples
#

examples-cpp: \
  examples \
  bmon-cpp

bmon-cpp: examples/bmon/stms.json
	$(GEN) -clean -output "out/bmon/ros" "templates/ros" $<