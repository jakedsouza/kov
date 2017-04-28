# vim: ft=make
# atom: grammar=Makefile

GO ?= go
GOVERSION ?= go1.8
SHELL := /bin/bash
GIT_VERSION = $(shell git describe --tags)

.DEFAULT_GOAL := help

.PHONY: check
check: goversion checkfmt swagger-validate ## Runs static code analysis checks
	@echo running metalint ...
	gometalinter --vendored-linters --install
	gometalinter --vendored-linters --vendor --disable=gotype --errors --fast --deadline=60s   ./...
	@echo running header check ...
	hack/header-check.sh

.PHONY: goversion
goversion: ## Checks if installed go version is latest
	@echo Checking go version...
	@( $(GO) version | grep -q $(GOVERSION) ) || ( echo "Please install $(GOVERSION) (found: $$($(GO) version))" && exit 1 )

.PHONY: checkfmt
checkfmt: ## Checks code format
	hack/gofmtcheck.sh

.PHONY: fmt
fmt: ## format go code
	goimports -v -w $$(find . -name '*.go' -not -path './vendor/*')

.PHONY: swagger-validate
swagger-validate: # Validates swagger files
	swagger validate ./swaggerspec/swagger.yml
	
.PHONY: distclean
distclean: ## Clean ALL files including ignored ones
	git clean -f -d -x .

.PHONY: clean
clean: ## Clean all modified files
	git clean -f -d .

.PHONY: generate
generate: ## run go generate
	$(GO) generate ./cmd

.PHONY: generate-fmt
generate-fmt: generate fmt ## Run go generate and fix go-fmt and headers
	./hack/header-check.sh fix

.PHONY: cli-dev
cli-dev: check generate-fmt ## Generates the cli for dev
	@DEV=1 ./hack/build.sh

.PHONY: cli
cli: check generate-fmt ## Generates the cli binary 
	@DEV=0 ./hack/build.sh

.PHONY: test
test: ## Run unit tests
	@echo running tests...
ifdef DRONE
	hack/coverage
else
	$(GO) test -v -race $(shell go list -v ./... | grep -v /vendor/)
endif

help: ## Display make help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
