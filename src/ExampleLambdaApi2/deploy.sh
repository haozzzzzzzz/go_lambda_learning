#!/usr/bin/env bash
echo generating api
lbuild compile api

echo building...
lbuild compile func

echo deploying...
ldeploy remote func