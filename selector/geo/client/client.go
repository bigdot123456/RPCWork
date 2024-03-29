package main

import (
	"context"
	"flag"
	"log"
	"time"

	example "github.com/bigdot123456/RPCWork"
	"github.com/smallnest/rpcx/client"
)

var (
	addr1 = flag.String("addr1", "tcp@0.0.0.0:8972", "server address")
	addr2 = flag.String("addr2", "tcp@localhost:8973", "server address")
)

func main() {
	flag.Parse()

	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1, Value: "latitude=39.9289&longitude=116.3883"},
		{Key: *addr2, Value: "latitude=139.3453&longitude=23.3243"}})
	xclient := client.NewXClient("Arith", client.Failtry, client.ConsistentHash, d, client.DefaultOption)
	defer xclient.Close()
	xclient.ConfigGeoSelector(39.30, 116.40)

	args := &example.Args{
		A: 10,
		B: 20,
	}

	for i := 0; i < 10; i++ {
		reply := &example.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)

		time.Sleep(time.Second)
	}

}
