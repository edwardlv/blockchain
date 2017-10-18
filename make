#!/usr/bin/env bash

if [ ! -f make ]; then
    echo 'make must be run within its container folder' 1>&2
    exit 1
    fi

    CURDIR=`pwd`
    OLDGOPATH="$GOPATH"
    export GOPATH="$CURDIR"

    gofmt  -w src

# go install 包含go文件的所有目录
go install block
go install main


export GOPATH="$OLDGOPATH"

echo 'finished'
