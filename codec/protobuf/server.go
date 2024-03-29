package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/bigdot123456/RPCWork/codec/protobuf/pb"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "0.0.0.0:8972", "server address")
)

type Arith int

func (t *Arith) Mul(ctx context.Context, args *pb.ProtoArgs, reply *pb.ProtoReply) error {
	reply.C = args.A * args.B
	fmt.Printf("call: %d * %d = %d\n", args.A, args.B, reply.C)
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()
	//s.RegisterName("Arith", new(example.Arith), "")
	s.Register(new(Arith), "")
	s.Serve("tcp", *addr)
}
