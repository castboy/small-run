package main

import (
	"context"
	"fmt"
)

type Arith int

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

func (i *Arith) Plus(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	fmt.Println(args.A, args.B, *reply)

	return nil
}

