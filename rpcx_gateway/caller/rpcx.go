package main

import (
	"context"
	metrics "github.com/rcrowley/go-metrics"
	client2 "github.com/rpcxio/rpcx-etcd/client"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"go.uber.org/fx"
	"time"
)

func NewRPCPlugin() *serverplugin.ZooKeeperRegisterPlugin {
	return &serverplugin.ZooKeeperRegisterPlugin{
		ServiceAddress:   "tcp@" + *addr,
		ZooKeeperServers: []string{"localhost:2181"},
		BasePath:         "/rpcx_test",
		Metrics:          metrics.NewRegistry(),
		UpdateInterval:   time.Minute,
	}
}

func NewRPCServer(lifecycle fx.Lifecycle, plugin *serverplugin.ZooKeeperRegisterPlugin) (*server.Server, error) {
	s := server.NewServer()
	err := plugin.Start()
	if err != nil {
		return nil, err
	}

	s.Plugins.Add(plugin)

	lifecycle.Append(
		fx.Hook{
			OnStop: func(ctx context.Context) error {
				err := s.UnregisterAll()
				if err != nil {
					return err
				}

				return s.Close()
			},
		})

	return s, nil
}

func NewRPCClient() (client.XClient, error) {
	d, err := client2.NewEtcdV3Discovery("/rpcx_test", "Arith", []string{"localhost:2379"}, false, nil)
	if err != nil {
		return nil, err
	}

	xclient := client.NewXClient("Arith", client.Failover, client.RoundRobin, d, client.DefaultOption)

	return xclient, nil
}
