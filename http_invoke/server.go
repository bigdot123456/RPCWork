package main

import (
	"flag"

	example "github.com/bigdot123456/RPCWork"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "0.0.0.0:8972", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	s.Register(new(example.Arith), "")
	s.Serve("tcp", *addr)
}
