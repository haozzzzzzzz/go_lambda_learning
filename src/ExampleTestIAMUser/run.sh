#!/usr/bin/env bash
echo generating api
lbuild compile api

echo running main
go build -o main main.go
./main