package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	example "github.com/bigdot123456/RPCWork"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "0.0.0.0:8972", "server address")
)

type Arith struct{}

// the second parameter is not a pointer
func (t *Arith) Mul(ctx context.Context, args example.Args, reply *example.Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()
	//s.Register(new(Arith), "")
	s.RegisterName("Arith", new(Arith), "")
	go s.Serve("tcp", *addr)

	time.Sleep(time.Minute)
	err := s.Shutdown(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("restart ...")
	s = server.NewServer()
	//s.Register(new(Arith), "")
	s.RegisterName("Arith", new(Arith), "")
	s.Serve("tcp", *addr)

}
