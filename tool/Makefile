

.PHONY: build install

build:
	CGO_ENABLED=0 go build -o bin/gopenai

install: build
	cp bin/gopenai ${GOPATH}/bin/