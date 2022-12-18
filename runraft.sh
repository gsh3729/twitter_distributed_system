#!/bin/sh

cd "raft/src/go.etcd.io/etcd/contrib/raftexample"
go build -o raftexample
go install github.com/mattn/goreman@latest
goreman start