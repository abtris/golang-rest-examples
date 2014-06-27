# GoLang Apiary REST Examples

- [API documentation](http://docs.restapi3.apiary.io) at [Apiary](http://apiary.io)

## Install go

    brew install go

As of go 1.2, a valid GOPATH is required to use the `go get` command:
  http://golang.org/doc/code.html#GOPATH

`go vet` and `go doc` are now part of the go.tools sub repo:
  http://golang.org/doc/go1.2#go_tools_godoc

To get `go vet` and `go doc` run:

    go get code.google.com/p/go.tools/cmd/godoc
    go get code.google.com/p/go.tools/cmd/vet

You may wish to add the GOROOT-based install location to your PATH:

    export PATH=$PATH:/usr/local/opt/go/libexec/bin

## Compile example

    cd ./get

    go build

## Run example

    ./get/get
