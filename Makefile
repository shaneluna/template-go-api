.DEFAULT_GOAL := help

.PHONY: help localchi localgin

help: ## : Show this help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_%-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

localchi: ## : Compile and run local chi binary.
	go build -o=build/localchi ./cmd/localchi/.
	./build/localchi

localgin: ## : Compile and run local gin binary.
	go build -o=build/localgin ./cmd/localgin/.
	./build/localgin