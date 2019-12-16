#!/bin/bash

echo "installing go version 1.12.7..."
apt install -y wget gcc musl-dev openssl
wget -O go.tgz https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz
tar -C /usr/local -xzf go.tgz
export GOPATH=$HOME/work
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

go version
