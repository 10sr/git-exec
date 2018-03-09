.PHONY: fmt build installdeps

fmt:
	gofmt -s -w *.go

build:
	go build -v .

installdeps:
	go get -d ./...
