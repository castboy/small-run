package main

import (
	"bytes"
	gateway "github.com/rpcxio/rpcx-gateway"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/smallnest/rpcx/codec"
)

type Args struct {
	A int
}

type Reply struct {
	C int
}

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			cc := &codec.MsgpackCodec{}

			args := &Args{
				A: 1,
			}

			data, _ := cc.Encode(args)

			req, err := http.NewRequest("POST", "http://127.0.0.1:9981/", bytes.NewReader(data))
			if err != nil {
				log.Fatal("failed to create request: ", err)
				return
			}

			h := req.Header
			// h.Set(gateway.XMessageID, "10000")
			// h.Set(gateway.XMessageType, "0")
			h.Set(gateway.XSerializeType, "3")
			h.Set(gateway.XServicePath, "Arith")
			h.Set(gateway.XServiceMethod, "Plus")

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Fatal("failed to call: ", err)
			}
			defer res.Body.Close()

			// handle http response
			replyData, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal("failed to read response: ", err)
			}

			reply := &Reply{}
			err = cc.Decode(replyData, reply)
			if err != nil {
				log.Fatal("failed to decode reply: ", err)
			}

			log.Printf("%d , total: %d", args.A, reply.C)
		}()
	}

	select {

	}
}
