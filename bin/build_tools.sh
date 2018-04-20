#!/usr/bin/env bash
export GOROOT=/usr/local/go
export GOPATH=/Users/hao/Documents/Projects/Github/go_lambda_learning

if [ -e lambda-build ]
then
    rm lambda-build
fi

if [ -e lambda-deploy ]
then
    rm lambda-deploy
fi

go build -o lambda-build ${GOPATH}/src/github.com/haozzzzzzzz/go-lambda/tool/lambda-build/main.go
#go build -o lambda-deploy ${GOPATH}/src/github.com/haozzzzzzzz/go-lambda/tool/lambda-deploy/main.go
