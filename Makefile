.DEFAULT_GOAL := help

.PHONY: help local

help: ## : Show this help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_%-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

local: ## : Compile and run local binary.
	go build -o=build/local ./cmd/local/.
	./build/local