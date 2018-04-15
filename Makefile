PKGS := $(shell go list ./... | grep -v /vendor)

VERSION := `git rev-parse HEAD`
LDFLAGS=-ldflags "-X=github.com/khoiracle/dido/cmd.version=$(VERSION)"

.PHONY: test install

all: test install

test:
	go test -v $(PKGS)

install:
	go install $(LDFLAGS)
