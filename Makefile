.PHONY: build help
.SHELLFLAGS = -c # Run commands in a -c flag 
.SILENT: ; # no need for @
.ONESHELL: ; # recipes execute in same shell
.NOTPARALLEL: ; # wait for this target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell
.DEFAULT_GOAL := help

deps: ## Install dependencies for Chart development
	helm plugin install https://github.com/instrumenta/helm-kubeval
	helm plugin install https://github.com/instrumenta/helm-conftest

lint: ## Lints the helm chart using kubeval
	helm lint
	helm kubeval . -v 1.18.0
 
conftest: ## Runs conftest tool on Helm chart
	helm conftest . --update

terratest: ## Runs Testing suite with terratest
	cd test
	go test .
	go test . -v -timeout 5m --tags=integration

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $, $}'