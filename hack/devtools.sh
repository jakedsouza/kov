#!/bin/sh

go get -u -v github.com/alecthomas/gometalinter
go get github.com/Masterminds/glide
go get -u -v github.com/go-swagger/go-swagger/cmd/swagger
go get -u -v github.com/mitchellh/gox

# install all the linters
gometalinter --install --update
