#!/usr/bin/env bash
export GOROOT=/usr/local/go
export GOPATH=/Users/hao/Documents/Projects/Github/go_lambda_learning

echo generating api
lbuild compile api

echo building...
lbuild compile func

echo deploying...
ldeploy remote func