
# Create version tag from git tag
VERSION=$(shell git describe | sed 's/^v//')
REPO=cybermaggedon/evs-elasticsearch
DOCKER=docker
GO=GOPATH=$$(pwd)/go go

all: evs-elasticsearch build

SOURCE=evs-elasticsearch.go es-model.go es-load.go es-mapping.go

evs-elasticsearch: ${SOURCE} go.mod go.sum
	${GO} build -o $@ ${SOURCE}

build: evs-elasticsearch
	${DOCKER} build -t ${REPO}:${VERSION} -f Dockerfile .

push:
	${DOCKER} push ${REPO}:${VERSION}

