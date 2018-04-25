#!/usr/bin/env bash
export GOROOT=/usr/local/go
export GOPATH=/Users/hao/Documents/Projects/Github/go_lambda_learning

if [ -e lbuild ]
then
    rm lbuild
fi

if [ -e ldeploy ]
then
    rm ldeploy
fi

go build -o lbuild ${GOPATH}/src/github.com/haozzzzzzzz/go-lambda/tool/lambda-build/main.go
go build -o ldeploy ${GOPATH}/src/github.com/haozzzzzzzz/go-lambda/tool/lambda-deploy/main.go
