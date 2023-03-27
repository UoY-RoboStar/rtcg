.POSIX:

VERSION = 0.1.1

GO          = go
INPUT_DIR   = ./cmd
BIN_DIR     = bin

REL_DIR             = release
REL_DIR_LINUX_AMD64 = $(REL_DIR)/linux-amd64
REL_DIR_WIN_AMD64   = $(REL_DIR)/windows-amd64
REL_DIR_MAC_ARM64   = $(REL_DIR)/macos-arm64

GEN         = $(BIN_DIR)/rtcg-gen
MAKE_STMS   = $(BIN_DIR)/rtcg-make-stms
READ_TRACES = $(BIN_DIR)/rtcg-read-traces

ALL_INPUTS = $(INPUT_DIR)/rtcg-gen $(INPUT_DIR)/rtcg-make-stms $(INPUT_DIR)/rtcg-read-traces

.PHONY: \
  all \
  rtcg-gen \
  rtcg-make-stms \
  rtcg-read-traces \
  examples \
  examples-cpp \
  bmon \
  bmon-cpp \
  release \
  release-linux-amd64 \
  release-windows-amd64 \
  release-macos-arm64 \
  release-examples


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

bmon-cpp: examples/bmon/gen.xml examples/bmon/stms.json
	$(GEN) -clean -output "out/bmon" examples/bmon/gen.xml examples/bmon/stms.json


#
# Release binary production
#

release: \
  release-linux-amd64 \
  release-windows-amd64 \
  release-macos-arm64 \
  release-examples

release-linux-amd64:
	env GOOS=linux GOARCH=amd64 $(GO) build -o $(REL_DIR_LINUX_AMD64)/rtcg-$(VERSION)/ $(ALL_INPUTS)
	rm -f $(REL_DIR_LINUX_AMD64)/rtcg-$(VERSION)-linux-amd64.tar.gz
	cd $(REL_DIR_LINUX_AMD64) && tar -czvf rtcg-$(VERSION)-linux-amd64.tar.gz rtcg-$(VERSION)/

release-windows-amd64:
	env GOOS=windows GOARCH=amd64 $(GO) build -o $(REL_DIR_WIN_AMD64)/rtcg-$(VERSION)/ $(ALL_INPUTS)
	rm -f $(REL_DIR_WIN_AMD64)/rtcg-$(VERSION)-windows-amd64.zip
	cd $(REL_DIR_WIN_AMD64) && zip -r rtcg-$(VERSION)-windows-amd64.zip rtcg-$(VERSION)/

release-macos-arm64:
	env GOOS=darwin GOARCH=arm64 $(GO) build -o $(REL_DIR_MAC_ARM64)/rtcg-$(VERSION)/ $(ALL_INPUTS)
	rm -f $(REL_DIR_MAC_ARM64)/rtcg-$(VERSION)-macos-arm64.tar.gz
	cd $(REL_DIR_MAC_ARM64) && tar -czvf rtcg-$(VERSION)-macos-arm64.tar.gz rtcg-$(VERSION)/

release-examples: examples-cpp
	mkdir -p $(REL_DIR)/rtcg-$(VERSION)-examples
	cp -r out $(REL_DIR)/rtcg-$(VERSION)-examples/out
	cp -r examples $(REL_DIR)/rtcg-$(VERSION)-examples/in
	cd $(REL_DIR) && zip -r rtcg-$(VERSION)-examples.zip rtcg-$(VERSION)-examples/
	cd $(REL_DIR) && tar -czvf rtcg-$(VERSION)-examples.tar.gz rtcg-$(VERSION)-examples/
