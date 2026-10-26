BINARY   := openspl
VERSION  ?= 0.1.0-dev
COMMIT   ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo none)
DATE     ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS  := -s -w \
	-X github.com/be-thomas/openspl/pkg/version.Version=$(VERSION) \
	-X github.com/be-thomas/openspl/pkg/version.Commit=$(COMMIT) \
	-X github.com/be-thomas/openspl/pkg/version.Date=$(DATE)

.PHONY: all build test race vet fmt fmt-check check hooks generate generate-check conformance conformance-update e2e clean help

all: check

build:
	go build -trimpath -ldflags "$(LDFLAGS)" -o $(BINARY) ./cmd/openspl

test:
	go test ./...

race:
	go test -race ./...

# Generated ANTLR parsers trip vet's unreachable analyzer: upstream codegen
# emits a defensive `goto` after a `return` in every rule. They are vetted with
# that one analyzer disabled; hand-written packages get the full set.
vet:
	@handwritten="$$(go list ./... | grep -v '/gen$$' || true)"; \
	if [ -n "$$handwritten" ]; then go vet $$handwritten; fi
	@generated="$$(go list ./... | grep '/gen$$' || true)"; \
	if [ -n "$$generated" ]; then go vet -unreachable=false $$generated; fi

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
	chmod +x .githooks/commit-msg .githooks/pre-commit scripts/check-commit-message.sh scripts/generate-parsers.sh conformance/engines/run.sh
	git config core.hooksPath .githooks
	@echo "installed: core.hooksPath=.githooks"

# Regenerate the ANTLR parsers. Needs Java 17+; builds never do.
generate:
	./scripts/generate-parsers.sh

generate-check:
	./scripts/generate-parsers.sh
	@changed="$$(git status --porcelain -- pkg/parser/spl1/gen pkg/parser/spl2/gen)"; \
	if [ -n "$$changed" ]; then \
		echo "generated parsers are out of date; run make generate and commit the result"; \
		echo "$$changed"; exit 1; \
	fi
	@echo "generated parsers are current"

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
		'generate   regenerate the ANTLR parsers (needs Java 17+)' \
		'generate-check  fail if committed parsers differ from the grammars' \
		'conformance  run the offline conformance suite' \
		'conformance-update  regenerate goldens and expectations' \
		'e2e        run the Docker engine harness' \
		'clean      remove build output'
