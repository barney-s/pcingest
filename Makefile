# Copyright 2019 The Kubernetes Authors.
# SPDX-License-Identifier: Apache-2.0
#
# Makefile for pcingest

COVER_FILE ?= cover.out
ARCH ?= amd64
ALL_ARCH = amd64 arm arm64 ppc64le s390x
IMAGE_NAME = pcingest

VER ?= v0.0.1
ifeq ($(ARCH), amd64)
IMAGE_TAG ?= $(VER)
else
IMAGE_TAG ?= $(ARCH)-$(VER)
endif

IMAGE ?= $(REGISTRY)/$(IMAGE_NAME):$(IMAGE_TAG)


.DEFAULT_GOAL := all
.PHONY: all
all: test bin/pcingest

# Run tests
.PHONY: test
test:
	go test -v ./... -coverprofile $(COVER_FILE)

# Build ingest binary
.PHONY: bin/pcingest
bin/pcingest: 
	go build -o bin/pcingest main.go

## --------------------------------------
## Docker
## --------------------------------------
.PHONY: docker-build
docker-build: test
	docker build --network=host --pull --build-arg ARCH=$(ARCH) . -t $(IMAGE)

.PHONY: docker-push
docker-push: ## Push the docker image
	docker push $(IMAGE)

.PHONY: clean
clean:
	go clean --cache
