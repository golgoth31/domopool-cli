#-----------------------------------------------------------------------------
# Global Variables
#-----------------------------------------------------------------------------

DOCKER_USER ?= $(DOCKER_USER)
DOCKER_PASS ?=

DOCKER_BUILD_ARGS := --build-arg HTTP_PROXY=$(http_proxy) --build-arg HTTPS_PROXY=$(https_proxy)
PROTOC_IMAGE := thethingsindustries/protoc:3.1.21
APP_VERSION := latest
PACKAGE ?= $(shell go list ./... | grep configs)
VERSION ?= $(shell git describe --tags --always || git rev-parse --short HEAD)
GIT_COMMIT=$(shell git rev-parse HEAD)
BUILD_DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

GOLINTER:=$(shell command -v golangci-lint 2> /dev/null)

override LDFLAGS += \
  -X ${PACKAGE}.Version=${VERSION} \
  -X ${PACKAGE}.BuildDate=${BUILD_DATE} \
  -X ${PACKAGE}.GitCommit=${GIT_COMMIT} \


#-----------------------------------------------------------------------------
# BUILD
#-----------------------------------------------------------------------------

.PHONY: default build test publish build_local lint artifact_linux artifact_darwin deploy
default:  test lint build swagger

test:
	go test -v ./...

build_assets:
	cd web && yarn build

run:
	go mod download
	go run main.go

build:
	cd web && yarn build
	go build -ldflags '${LDFLAGS}' -o ./domopool main.go

build_local:
	go build -ldflags '${LDFLAGS}' -o ./domopool main.go
build_docker:
	docker build $(DOCKER_BUILD_ARGS) -t localhost/domopool:$(APP_VERSION) -f ./build/Dockerfile .
lint:
ifndef GOLINTER
	GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.15
endif
	golangci-lint run

artifact_linux:
	GOPROXY=https://proxy.golang.org CGO_ENABLED=0 GOOS=linux go build -ldflags '${LDFLAGS}' -o domopool-linux

artifact_darwin:
	GOPROXY=https://proxy.golang.org CGO_ENABLED=0 GOOS=darwin go build -ldflags '${LDFLAGS}' -o domopool-darwin

#-----------------------------------------------------------------------------
# DEPLOY
#-----------------------------------------------------------------------------
deploy:
	kubectl apply -f deployments/k8s/domosense.yml
#-----------------------------------------------------------------------------
# PUBLISH
#-----------------------------------------------------------------------------

.PHONY: publish

publish:
	docker push $(DOCKER_USER)/domopool:$(APP_VERSION)

#-----------------------------------------------------------------------------
# CLEAN
#-----------------------------------------------------------------------------

.PHONY: clean

clean:
	rm -rf domosense
