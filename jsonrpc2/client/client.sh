#!/bin/sh

curl --header "X-JSONRPC-2.0: true" "http://0.0.0.0:8972/" \
-d '{"jsonrpc": "2.0", "method": "Arith.Mul", "params": {"A": 10, "B": 20}, "id": 3}'