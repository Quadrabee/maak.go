PLATFORM=darwin/amd64

build: maak

maak: make/bindata.go go.* *.go **/*.go
	go build

make/bindata.go: make/templates/*.tpl
	go-bindata -o make/bindata.go -pkg make make/templates

.PHONY: bin/maak

bin/maak:
	@DOCKER_BUILDKIT=1 docker build . --target bin \
		--output bin/ \
		--platform ${PLATFORM}

