#!/bin/bash

gpm help &> /dev/null
if [ $? -eq 0 ]; then

  rm -rf $GOPATH
  mkdir -p $GOPATH

  gpm install
  mkdir -p $GOPATH/src/github.com/jarosser06
  ln -s ${PWD} $GOPATH/src/github.com/jarosser06/gcm
else
  echo "gpm not found"
fi
