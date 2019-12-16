#!/bin/bash

echo "installing go version 1.12.7..."
apk add --no-cache --virtual .build-deps bash gcc musl-dev openssl go
wget -O go.tgz https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz
tar -C /usr/local -xzf go.tgz
cd /usr/local/go/src/
pwd
ls -al
./make.bash
# export PATH="/usr/local/go/bin:$PATH"
# export GOPATH="/opt/go/"
# export PATH="$PATH:$GOPATH/bin"
export PATH=/usr/local/go/bin:$PATH
export GOPATH=/usr/bin/go
export PATH=$PATH:$GOPATH/bin
# apk del .build-deps
# go version

# curl -O https://dl.google.com/go/go1.11.13.linux-amd64.tar.gz
#
# tar -C /usr/local -xzf go1.11.13.linux-amd64.tar.gz 

export PATH=$PATH:/usr/local/go/bin
