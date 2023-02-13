.POSIX:
GO         = go
INPUT_DIR  = ./cmd
OUTPUT_DIR = bin

.PHONY: \
  all \
  rtcg-gen \
  rtcg-make-stms \
  rtcg-read-traces

all: \
  rtcg-gen \
  rtcg-make-stms \
  rtcg-read-traces

rtcg-gen:
	$(GO) build -o $(OUTPUT_DIR)/ $(INPUT_DIR)/rtcg-gen

rtcg-make-stms:
	$(GO) build -o $(OUTPUT_DIR)/ $(INPUT_DIR)/rtcg-make-stms

rtcg-read-traces:
	$(GO) build -o $(OUTPUT_DIR)/ $(INPUT_DIR)/rtcg-read-traces
