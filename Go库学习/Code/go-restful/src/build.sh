#!/bin/bash
PROJ_DIR=$(cd ..;pwd)
export GOPATH=$GOPATH:$PROJ_DIR
go build main.go