PLATFORM := $(or ${PLATFORM},${PLATFORM},darwin/amd64)

build: maak

maak: make/bindata.go go.* *.go **/*.go
	go build

make/bindata.go: make/templates/*.tpl
	go-bindata -o make/bindata.go -pkg make make/templates

.PHONY: bin/maak

test: bin/maak
	cd examples && \
	PATH=${PWD}/bin/:${PATH} ./test.sh

bin/maak: go.* *.go **/*.go
	@DOCKER_BUILDKIT=1 docker build . --target bin \
		--output bin/ \
		--platform ${PLATFORM}

