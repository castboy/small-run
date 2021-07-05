package main

import (
	"context"
	"go.uber.org/fx"
	"time"

	metrics "github.com/rcrowley/go-metrics"
	"github.com/rpcxio/rpcx-etcd/serverplugin"
	"github.com/smallnest/rpcx/server"
)

func NewRPCPlugin() *serverplugin.EtcdV3RegisterPlugin {
	return &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		EtcdServers:    []string{"localhost:2379"},
		BasePath:       "/rpcx_test",
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
}

func NewRPCServer(lifecycle fx.Lifecycle, plugin *serverplugin.EtcdV3RegisterPlugin) (*server.Server, error) {
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