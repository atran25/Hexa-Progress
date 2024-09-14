default: help

.PHONY: help
help: ## Show help for each of the Makefile recipes.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

test: ## Run tests with coverage
	go test -v ./... -coverprofile=coverage.out

coverage: ## Show coverage
	go tool cover -func coverage.out

docker: ## Build & run docker image
	docker build -t hexaprogress .
	docker run --rm -it -p 8080:8080 -e HEXAPROGRESS_DISCORD_TOKEN=$(dt) -e HEXAPROGRESS_GUILD_ID=$(gid) hexaprogress