#!/bin/bash

CLOUD=aws
ZONE=sa-east-1

go build -o gifkernel server.go

ops pkg from-run --name gifkernel --version v0.0 gifkernel

ops image create -t $CLOUD -c config.json --package gifkernel_v0.0 -i gifkernel --local --show-debug

ops instance create gifkernel -c config.json -t $CLOUD -z $ZONE --show-debug
