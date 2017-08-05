#!/bin/bash

export BUILDSH_PATH=$(realpath "${0}")
export GOPATH=$(dirname "${BUILDSH_PATH}")

go build -o transfer.sh "${GOPATH}/src/transfersh/main.go"
