DEPLOY_DIR = $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
ROOT_DIR = $(realpath $(DEPLOY_DIR)/..)

DOCKER = DOCKER_BUILDKIT=1 docker

HUB_ID = ihac
VERSION = 1.0.0

SERVICES = accounts reviews products gateway apollo-gateway

.PHONY: all $(SERVICES)

all: $(SERVICES)

$(SERVICES):
	$(DOCKER) build -f $(DEPLOY_DIR)/$@/Dockerfile -t $(HUB_ID)/$@:$(VERSION) $(ROOT_DIR)