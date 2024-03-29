package main

import (
	"flag"
	"log"
	"time"

	metrics "github.com/rcrowley/go-metrics"
	example "github.com/bigdot123456/RPCWork"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var (
	addr     = flag.String("addr", "0.0.0.0:8972", "server address")
	etcdAddr = flag.String("etcdAddr", "localhost:2379", "etcd address")
	basePath = flag.String("base", "/rpcx_test", "prefix path")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	addRegistryPlugin(s)

	s.RegisterName("Arith", new(example.Arith), "")
	s.Serve("tcp", *addr)
}

func addRegistryPlugin(s *server.Server) {

	r := serverplugin.NewMDNSRegisterPlugin("tcp@"+*addr, 8972, metrics.NewRegistry(), time.Minute, "")
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
