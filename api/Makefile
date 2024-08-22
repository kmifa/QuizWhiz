MODULE_NAME := github.com/kmifa/QuizWhiz

GCP_PROJECT_ID ?= aaaaaaa
GCP_REGION ?= asia-northeast1

SERVICE_NAME ?= quiz-whiz-api

GIT_REF := $(shell git describe --always)
VERSION ?= commit-$(GIT_REF)

PACKAGES ?= $(shell go list ./...)

.PHONY: fmt
fmt:
	@cd src; go mod tidy
	@cd src; goimports -local '$(MODULE_NAME)' -w .
	@cd src; gofmt -w .

.PHONY: gen
gen:
	@cd src; go generate ./...
	@make fmt

.PHONY: test
test:
	@cd src; go test -run=. $(PACKAGES)

.PHONY: test-coverage
test-coverage:
	@mkdir -p ./test-coverage
	@cd src; go test -coverprofile=../test-coverage/coverage.out -covermode=atomic $(PACKAGES)
	@cd src; go tool cover -html=../test-coverage/coverage.out -o ../test-coverage/coverage.html

.PHONY: mock
mock: ## Generate a mock corresponding to a constant
	@cd src; go generate -x -run="mockgen.*" ./...
	@make fmt

.PHONY: build
build:
	@cd src; go build -v -o bin/api $(MODULE_NAME)/cmd

.PHONY: docker/build
docker/build:
	@docker build \
	--build-arg VERSION=$(VERSION) \
	--build-arg MODULE_NAME=$(MODULE_NAME) \
	-t $(IMAGE_LOCAL) \
	-f docker/api.Dockerfile \
	./src