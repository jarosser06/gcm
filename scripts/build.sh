#!/bin/bash

if ! [ -a ${PWD}/.gopath ]; then
  echo "Fetching dependencies..."
  scripts/deps.sh
fi

mkdir -p ${PWD}/bin
go build -o bin/gcm gcm.go
