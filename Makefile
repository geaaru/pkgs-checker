NAME := pkgs-checker
PACKAGE_NAME ?= $(NAME)
GOLANG_VERSION=$(shell go env GOVERSION)
REVISION := $(shell git rev-parse --short HEAD || echo dev)
VERSION := $(shell git describe --tags || echo $(REVISION))
VERSION := $(shell echo $(VERSION) | sed -e 's/^v//g')
BUILD_PLATFORMS ?= -osarch="linux/amd64" -osarch="linux/386" -osarch="linux/arm"

# go tool nm ./luet | grep Commit
override LDFLAGS += -X "github.com/geaaru/pkgs-checker/cmd.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')"
override LDFLAGS += -X "github.com/geaaru/pkgs-checker/cmd.BuildCommit=$(shell git rev-parse HEAD)"
override LDFLAGS += -X "github.com/geaaru/pkgs-checker/cmd.BuildGoVersion=$(GOLANG_VERSION)"


.PHONY: all

all: pkgs-checker

.PHONY: pkgs-checker
pkgs-checker:
	# pkgs-checker uses go-sqlite3 that require CGO
	CGO_ENABLED=1 go build -ldflags '$(LDFLAGS)'

.PHONY: test
test:
	go test -v -tags all -cover -race ./...

.PHONY: coverage
coverage:
	go test ./... -race -coverprofile=coverage.txt -covermode=atomic

.PHONY: test-coverage
test-coverage:
	scripts/ginkgo.coverage.sh --codecov

.PHONY: clean
clean:
	-rm pkgs-checker
	-rm -rf release/ dist/

.PHONY: build
build:
	CGO_ENABLED=0 go build -ldflags '$(LDFLAGS)'

.PHONY: build-small
build-small:
	@$(MAKE) LDFLAGS+="-s -w" build
	upx --brute -1 $(NAME)

.PHONY: deps
deps:
	go env
	# Installing dependencies...
	GO111MODULE=on go install -mod=mod golang.org/x/lint/golint
	GO111MODULE=on go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo
	go get github.com/onsi/gomega/...
	ginkgo version

.PHONY: goreleaser-snapshot
goreleaser-snapshot:
	rm -rf dist/ || true
	GOVERSION=$(GOLANG_VERSION) goreleaser release --skip=validate,publish --snapshot --verbose
