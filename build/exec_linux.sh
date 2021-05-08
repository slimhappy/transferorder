#!/bin/zsh
export GOOS="linux"
go build transferorder
if [[ $? -ne 0 ]]
then
    echo "build faild"
    exit
fi