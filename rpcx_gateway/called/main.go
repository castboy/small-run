package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	"go.uber.org/fx"
)

var addr = flag.String("addr", "localhost:8973", "server address")

func main() {
	flag.Parse()

	fx.New(Modules, Server).Run()
}

var Modules = fx.Provide(
	NewRPCPlugin,
	NewRPCServer,
)

var Server = fx.Invoke(
	func(s *server.Server) {
		s.RegisterName("Arith", new(Arith), "")
		s.Serve("tcp", *addr)
	},
)