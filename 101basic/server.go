package main

import (
	"flag"
	// example "rpcx-work"

	myRPC "github.com/bigdot123456/RPCWork/rpcService"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "0.0.0.0:8972", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	//s.RegisterName("Arith", new(example.Arith), "")
	s.Register(new(myRPC.Arith), "")
	s.Serve("tcp", *addr)
}
