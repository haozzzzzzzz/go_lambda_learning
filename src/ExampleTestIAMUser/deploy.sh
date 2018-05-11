#!/usr/bin/env bash

stage=$1

if [ -z ${stage} ]
then
    stage="dev"
fi

echo deploy target stage environment: ${stage}

echo generating api
lbuild compile api

echo building...
lbuild compile func -s ${stage}

echo deploying...
ldeploy remote func -s ${stage}
