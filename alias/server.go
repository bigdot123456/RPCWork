package main

import (
	"flag"

	example "github.com/bigdot123456/RPCWork"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var (
	addr = flag.String("addr", "0.0.0.0:8972", "server address")
)

func main() {
	flag.Parse()

	a := serverplugin.NewAliasPlugin()
	a.Alias("a.b.c.D", "Times", "Arith", "Mul")
	s := server.NewServer()
	s.Plugins.Add(a)
	s.RegisterName("Arith", new(example.Arith), "")
	err := s.Serve("reuseport", *addr)
	if err != nil {
		panic(err)
	}
}
