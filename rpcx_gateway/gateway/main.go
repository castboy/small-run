package main

import (
	gateway "github.com/rpcxio/rpcx-gateway"
	"github.com/rpcxio/rpcx-gateway/gin"
	"github.com/smallnest/rpcx/client"
	"log"
)

func main() {
	d, err := client.NewZookeeperDiscoveryTemplate("/rpcx_test", []string{"127.0.0.1:2181"}, nil)
	if err != nil {
		log.Fatal(err)
	}

	httpServer := gin.New(":9981")
	gw := gateway.NewGateway("/", httpServer, d, client.FailMode(client.Failover), client.SelectMode(client.RoundRobin), client.DefaultOption)

	gw.Serve()
}
