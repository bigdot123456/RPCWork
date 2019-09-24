package main

import (
	"flag"

	//myRPC "github.com/bigdot123456/RPCWork/rpcService"
	"RPCWork/rpcService"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "0.0.0.0:8972", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	//s.RegisterName("Arith", new(example.Arith), "")
	s.Register(new(rpcService.Arith), "")
	s.Serve("tcp", *addr)
}
