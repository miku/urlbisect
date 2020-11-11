SHELL := /bin/bash
TARGETS := urlbisect
VERSION := $(shell git rev-parse --short HEAD)
BUILDTIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

GOLDFLAGS += -X main.Version=$(VERSION)
GOLDFLAGS += -X main.Buildtime=$(BUILDTIME)
GOFLAGS = -ldflags "$(GOLDFLAGS)"

.PHONY: all
all: $(TARGETS)

%: cmd/%/main.go bisect.go
	go build -o $@ -ldflags "$(GOLDFLAGS)" $<

.PHONY: clean
clean:
	rm -f $(TARGETS)

