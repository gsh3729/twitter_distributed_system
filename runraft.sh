#!/bin/sh

export ETCDPATH=/tmp/etcd-download-test/
export PATH=$ETCDPATH:$PATH
go install github.com/mattn/goreman@latest
goreman start