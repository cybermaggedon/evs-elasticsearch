
# Create version tag from git tag
VERSION=$(shell git describe | sed 's/^v//')
REPO=cybermaggedon/evs-elasticsearch
DOCKER=docker
GO=GOPATH=$$(pwd)/go go

all: evs-elasticsearch build

SOURCE=main.go model.go load.go mapping.go config.go

evs-elasticsearch: ${SOURCE} go.mod go.sum
	${GO} build -o $@ ${SOURCE}

build: evs-elasticsearch
	${DOCKER} build -t ${REPO}:${VERSION} -f Dockerfile .
	${DOCKER} tag ${REPO}:${VERSION} ${REPO}:latest

push:
	${DOCKER} push ${REPO}:${VERSION}
	${DOCKER} push ${REPO}:latest

