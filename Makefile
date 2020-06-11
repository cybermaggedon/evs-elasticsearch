
# Create version tag from git tag
VERSION=$(shell git describe | sed 's/^v//')
REPO=cybermaggedon/evs-elasticsearch
DOCKER=docker
GO=GOPATH=$$(pwd)/go go

all: evs-elasticsearch build

evs-elasticsearch: evs-elasticsearch.go es-model.go go.mod go.sum
	${GO} build -o $@ evs-elasticsearch.go es-model.go

build: evs-elasticsearch
	${DOCKER} build -t ${REPO}:${VERSION} -f Dockerfile .

push:
	${DOCKER} push ${REPO}:${VERSION}

