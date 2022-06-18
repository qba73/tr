.PHONY: help check cover test tidy

ROOT			:= $(PWD)
GO_HTML_COV 		:= ./coverage.html
GO_TEST_OUTFILE 	:= ./c.out
GO_DOCKER_IMAGE 	:= golang:1.18
CC_PREFIX := github.com/qba73/testrail


define PRINT_HELP_PYSCRIPT
import re, sys
for line in sys.stdin:
	match = re.match(r'^([a-zA-Z_-]+):.*?## (.*)$$', line)
	if match:
		target, help = match.groups()
		print("%-20s %s" % (target, help))
endef
export PRINT_HELP_PYSCRIPT

default: help

help:
	@python -c "$$PRINT_HELP_PYSCRIPT" < $(MAKEFILE_LIST)

check: ## Run static check analyzer
	staticcheck ./...

cover: ## Run unit tests and generate test coverage report
	go test -v ./... -count=1 -covermode=count -coverprofile=coverage.out
	go tool cover -html coverage.out

test: ## Run unit tests locally
	go test -v -count=1

# MODULES
tidy: ## Run go mod tidy and vendor
	go mod tidy
	go mod vendor

lint: ## Run linter inside container
	docker run --rm -v ${ROOT}:/data cytopia/golint .
