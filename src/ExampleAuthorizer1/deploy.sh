#!/usr/bin/env bash
echo building...
lbuild compile func

echo deploying...
ldeploy remote func