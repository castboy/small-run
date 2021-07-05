package main

import (
	"context"
	"fmt"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/smallnest/rpcx/client"
	"log"
)

type Arith struct {
	xclient client.XClient
	rclient *goredislib.Client
	rmutex  *redsync.Mutex
}

func NewArith(xclient client.XClient, rclient *goredislib.Client, rmutex  *redsync.Mutex) *Arith {
	return &Arith{
		xclient: xclient,
		rclient: rclient,
		rmutex: rmutex,
	}
}

type PlusArgs struct {
	A int
	B int
}

type Args struct {
	A int
}

type Reply struct {
	C int
}

func (t *Arith) Plus(ctx context.Context, args *Args, reply *Reply) error {
	if err := t.rmutex.Lock(); err != nil {
		panic(err)
	}

	defer func() {
		if ok, err := t.rmutex.Unlock(); !ok || err != nil {
			panic("unlock failed")
		}
	}()

	b, _ := t.rclient.Get(ctx, "plus-B").Int()

	plusArgs := &PlusArgs{
		A: args.A,
		B: b,
	}

	plusReply := &Reply{}

	err := t.xclient.Call(context.Background(), "Plus", plusArgs, plusReply)
	if err != nil {
		return err
	}

	err = t.rclient.Set(ctx, "plus-B", plusReply.C, 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	reply.C = plusReply.C

	fmt.Println(args.A, reply.C)

	return nil
}
