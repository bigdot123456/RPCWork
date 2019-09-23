package main

import (
	"context"
	"flag"
	"log"
	example "rpcx-work"
	"time"

	myRPC "github.com/bigdot123456/RPCWork/rpcService"
	"github.com/smallnest/rpcx/client"
)

var (
//	addr = flag.String("addr", "0.0.0.0:8972", "server address")
	addr = flag.String("addr", "10.0.0.4:8972", "server address")
)

func main() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &myRPC.Args{
		A: 100,
		B: 200,
	}

	for {
		reply := &example.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		time.Sleep(1e9)
	}

}
