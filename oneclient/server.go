package main

import (
	"context"
	"flag"

	example "github.com/bigdot123456/RPCWork"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "0.0.0.0:8972", "server address")
)

type Echo int

func (*Echo) Say(ctx context.Context, args string, reply *string) error {
	*reply = args
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()
	//s.RegisterName("Arith", new(example.Arith), "")
	s.Register(new(example.Arith), "")
	s.RegisterName("echo", new(Echo), "")
	s.Serve("tcp", *addr)
}
