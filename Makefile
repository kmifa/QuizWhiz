MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
ROOT_DIR := $(abspath $(MAKEFILE_DIR)..)

MODULE_NAME := github.com/kmifa/QuizWhiz

.PHONY: install
install: ## Install dependencies
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@go install golang.org/x/tools/cmd/goimports@latest
