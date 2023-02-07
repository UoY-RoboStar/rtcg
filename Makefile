.PHONY: all bin/rtcg-gen bin/rtcg-read-traces

all: bin/rtcg-gen bin/rtcg-read-traces

bin/rtcg-gen:
	go build -o bin/ ./cmd/rtcg-gen

bin/rtcg-read-traces:
	go build -o bin/ ./cmd/rtcg-read-traces
