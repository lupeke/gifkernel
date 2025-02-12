#!/bin/bash

go build -o gifkernel server.go

ops pkg from-run --name gifkernel --version v0.0 gifkernel

ops instance create gifkernel -c config.json -t aws -z sa-east-1 --show-debug
