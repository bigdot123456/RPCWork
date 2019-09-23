#!/bin/bash

xgen -o cmd/main.go -r "etcd" -pkg github.com/bigdot123456/RPCWork/xgen

# go run -tags "etcd"  cmd/main.go