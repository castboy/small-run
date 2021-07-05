package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	"go.uber.org/fx"
)

var addr = flag.String("addr", "localhost:8972", "server address")

func main() {
	flag.Parse()

	fx.New(Modules, Register)
}

var Modules = fx.Provide(
	NewRPCPlugin,
	NewRPCServer,
	NewRPCClient,
	NewRedisClient,
	NewRedisMutex,
	NewArith,
)

var Register = fx.Invoke(
	func(arith *Arith, s *server.Server) {
		s.RegisterName("Arith", arith, "")
		s.Serve("tcp", *addr)
	},
)
