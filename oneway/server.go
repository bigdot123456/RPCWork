package main

import (
	"context"
	"flag"
	"fmt"

	example "github.com/bigdot123456/RPCWork"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "0.0.0.0:8972", "server address")
)

type Arith struct{}

// the second parameter is not a pointer
func (t *Arith) Mul(ctx context.Context, args example.Args, reply *interface{}) error {
	fmt.Printf("args: %v, oneway:true\n", args)
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()
	//s.Register(new(Arith), "")
	s.RegisterName("Arith", new(Arith), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}
