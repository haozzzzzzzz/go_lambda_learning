#!/usr/bin/env bash
export GOROOT=/usr/local/go
export GOPATH=/Users/hao/Documents/Projects/Github/go_lambda_learning

if [ -e lamb ]
then
    rm lamb
fi

if [ -e lamd ]
then
    rm lamd
fi

go build -o lamb ${GOPATH}/src/github.com/haozzzzzzzz/go-lambda/tool/lambda-build/main.go
go build -o lamd ${GOPATH}/src/github.com/haozzzzzzzz/go-lambda/tool/lambda-deploy/main.go
