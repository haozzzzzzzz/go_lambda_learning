#!/usr/bin/env bash
source params.sh

export GOROOT=${goRoot}
export GOPATH=${goPath}
export GOOS=linux
export GOARCH=amd64

api compile -p ../
go build -o ../stage/${stage}/main ../main.go
cd ../stage/${stage}/
zip -r deploy_${stage}_${serviceName}.zip main config
