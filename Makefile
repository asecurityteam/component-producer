TAG := $(shell git rev-parse --short HEAD)
DIR := $(shell pwd -L)
.PHONY := dep lint test integration coverage
# SDCLI
SDCLI_VERSION=v1.5
SDCLI=docker run --rm -v "$(DIR):$(DIR)" -w "$(DIR)"  asecurityteam/sdcli:$(SDCLI_VERSION)

vendor:
	$(SDCLI) go mod vendor

dep: vendor

lint: dep
	$(SDCLI) go lint

test: dep
	$(SDCLI) go test

integration: dep
	$(SDCLI) go integration

coverage: dep
	$(SDCLI) go coverage

doc: ;

build-dev: ;

build: ;

run: ;

deploy-dev: ;

deploy: ;
