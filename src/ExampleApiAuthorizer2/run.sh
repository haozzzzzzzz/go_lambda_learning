#!/usr/bin/env bash
echo building...
lbuild compile func

echo running main
go build -o main main.go
./main