BINARY   := openspl
VERSION  ?= 0.1.0-dev
COMMIT   ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo none)
DATE     ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS  := -s -w \
	-X github.com/be-thomas/openspl/pkg/version.Version=$(VERSION) \
	-X github.com/be-thomas/openspl/pkg/version.Commit=$(COMMIT) \
	-X github.com/be-thomas/openspl/pkg/version.Date=$(DATE)

.PHONY: all build test race vet fmt fmt-check check hooks conformance conformance-update e2e clean help

all: check

build:
	go build -trimpath -ldflags "$(LDFLAGS)" -o $(BINARY) ./cmd/openspl

test:
	go test ./...

race:
	go test -race ./...

vet:
	go vet ./...

fmt:
	gofmt -w .

fmt-check:
	@files="$$(gofmt -l .)"; \
	if [ -n "$$files" ]; then \
		echo "gofmt required for:"; echo "$$files"; exit 1; \
	fi

check: fmt-check vet test build
	@echo "ok: format, vet, tests, build"

hooks:
	chmod +x .githooks/commit-msg .githooks/pre-commit scripts/check-commit-message.sh conformance/engines/run.sh
	git config core.hooksPath .githooks
	@echo "installed: core.hooksPath=.githooks"

conformance:
	go test ./conformance -count=1

conformance-update:
	go test ./conformance -count=1 -update

e2e:
	./conformance/engines/run.sh

clean:
	rm -f $(BINARY)
	rm -rf dist

help:
	@printf '%s\n' \
		'build      compile the CLI' \
		'test       run unit and conformance tests' \
		'race       run tests with the race detector' \
		'vet        run go vet' \
		'fmt        format Go sources' \
		'fmt-check  fail if sources are not formatted' \
		'check      format check + vet + test + build' \
		'hooks      install the commit hooks' \
		'conformance  run the offline conformance suite' \
		'conformance-update  regenerate goldens and expectations' \
		'e2e        run the Docker engine harness' \
		'clean      remove build output'
