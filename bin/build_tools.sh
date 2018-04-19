#!/usr/bin/env bash
export GOPATH=/Users/hao/Documents/Projects/Github/go_lambda_learning
go build -o lambda-build ${GOPATH}/src/github.com/haozzzzzzzz/go-lambda/tool/lambda-build/main.go
go build -o lambda-deploy ${GOPATH}/src/github.com/haozzzzzzzz/go-lambda/tool/lambda-build/main.go