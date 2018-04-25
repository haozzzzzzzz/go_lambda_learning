#!/usr/bin/env bash
echo generating api
lamb compile api

echo building...
lamb compile func

echo deploying...
lamd remote func