package main

import (
    `context`
    `errors`
    `fmt`
    "log"
    "time"

    "go.etcd.io/etcd/api/v3/mvccpb"
    "go.etcd.io/etcd/client/v3"
)

const (
    dialTimeout = time.Second * 3
)

var endpoints = []string{"127.0.0.1:2379"}

func NewConfig(dialTimeout time.Duration, endpoints []string) clientv3.Config {
    return clientv3.Config{
        DialTimeout: dialTimeout,
        Endpoints:   endpoints,
    }
}

func NewClient(cfg clientv3.Config) *clientv3.Client {
    client, err := clientv3.New(cfg)
    if err != nil {
        log.Fatal(err)
    }
    return client
}

type Service struct {
    key string
    val string
    client *clientv3.Client
    keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
    leaseID       clientv3.LeaseID
}

func NewService(key, val string) *Service {
    cli := NewClient(NewConfig(dialTimeout, endpoints))
    return &Service{
        client: cli,
        key: key,
        val: val,
    }
}

// 注册服务并给定一个租期（租约）
func (s *Service) RegisterServiceWithLease(lease int64) error {
    // 申请一个 lease 时长的租约
    respLease, err := s.client.Grant(context.TODO(), lease)
    if err != nil {
        return err
    }
    // 注册服务并绑定租约
    _, err = s.client.Put(context.TODO(), s.key, s.val, clientv3.WithLease(respLease.ID))
    if err != nil {
        return err
    }

    // 续租
    keepAliveChan, err := s.client.KeepAlive(context.TODO(), respLease.ID)
    if err != nil {
        return err
    }
    s.keepAliveChan = keepAliveChan
    s.leaseID = respLease.ID
    return nil
}

// 获取服务，根据前缀获取所有服务
func (s *Service) GetServiceWithPrefix(prefix string) (map[string]string, error) {
    serviceMap := make(map[string]string, 0)

    resp, err := s.client.Get(context.TODO(), prefix, clientv3.WithPrefix())
    if err != nil {
        return serviceMap, err
    }
    if resp.Kvs == nil || len(resp.Kvs) == 0 {
        return serviceMap, errors.New("this is no service")
    }

    for i := range resp.Kvs {
        if val := resp.Kvs[i].Value; val != nil {
            key := string(resp.Kvs[i].Key)
            serviceMap[key] = string(resp.Kvs[i].Value)
        }
    }
    return serviceMap, nil
}

// 监听续租情况
func (s *Service) ListenLease() error {
    go func() {
        for {
            select {
            case keepResp := <-s.keepAliveChan:
                if keepResp == nil {
                    log.Println("lease is failed")
                    return
                } else {
                    log.Println("grant lease: ", s.leaseID)
                }
            }
        }
    }()

    return nil
}

// 撤销租约
func (s *Service) RevokeLease() error {
    if _, err := s.client.Revoke(context.TODO(), s.leaseID); err != nil {
        return err
    }
    return nil
}

// 删除服务，根据前缀删除服务
func (s *Service) DeleteServiceWithPrefix(prefix string) error {
    _, err := s.client.Delete(context.TODO(), prefix, clientv3.WithPrefix())
    if err != nil {
        log.Println(err)
        return err
    }

    return nil
}

// 监听服务变化，根据前缀来监听服务操作变化
func (s *Service) ListenServiceWithPrefix(prefix string) error {
    watchRespChan := s.client.Watch(context.TODO(), prefix, clientv3.WithPrefix())

    for watchResp := range watchRespChan {
        for _, event := range watchResp.Events {
            switch event.Type {
            case mvccpb.PUT:
                log.Println("etcd put operation", string(event.Kv.Value))
            case mvccpb.DELETE:
                log.Println("etcd delete operation")
            }
        }
    }

    return nil
}

func main() {
    srv := NewService("charging/nmott-1", "charging-nmott-1")
    err := srv.RegisterServiceWithLease(10)
    if err != nil {
        panic(err)
    }

    srv.ListenLease()

    srv2 := NewService("charging/nmott-2", "charging-nmott-2")
    err = srv2.RegisterServiceWithLease(10)
    if err != nil {
        panic(err)
    }

    srv2.ListenLease()

    srvs, err := srv.GetServiceWithPrefix("charging")
    if err != nil {
        panic(err)
    }

    fmt.Println(srvs)

    go srv.ListenServiceWithPrefix("charging")

    err = srv.RevokeLease()
    if err != nil {
        panic(err)
    }

    srvs, err = srv.GetServiceWithPrefix("charging")
    if err != nil {
        panic(err)
    }

    fmt.Println(srvs)

    err = srv2.RevokeLease()
    if err != nil {
        panic(err)
    }

    err = srv2.DeleteServiceWithPrefix("charging")
    if err != nil {
        panic(err)
    }

    time.Sleep(time.Minute)
}