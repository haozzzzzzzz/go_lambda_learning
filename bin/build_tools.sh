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

if [ -e api ]
then
    rm api
fi

# lambda
go build -o lbuild ${GOPATH}/src/github.com/haozzzzzzzz/go-lambda/tool/lambda-build/main.go
go build -o ldeploy ${GOPATH}/src/github.com/haozzzzzzzz/go-lambda/tool/lambda-deploy/main.go

# api project
go build -o api ${GOPATH}/src/github.com/haozzzzzzzz/go-rapid-development/tools/api/main.go

