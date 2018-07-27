package main

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/SUN-XIN/go-example/protobuf_simple/types"
)

const (
	address = "localhost:8081"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := types.NewSumClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &types.SumRequest{ValInt1: 1, ValInt2: 2}
	r, err := c.Sum(ctx, req)
	if err != nil {
		log.Fatalf("Failed Sum: %v", err)
	}
	log.Printf("Int Res of %+v: %d", *req, r.ResInt)

	req = &types.SumRequest{ValString1: "a", ValString2: "b"}
	r, err = c.Sum(ctx, req)
	if err != nil {
		log.Fatalf("Failed Sum: %v", err)
	}
	log.Printf("String Res of %+v: %s", *req, r.ResString)

	req = &types.SumRequest{ValFloat1: 1.1, ValFloat2: 2.2}
	r, err = c.Sum(ctx, req)
	if err != nil {
		log.Fatalf("Failed Sum: %v", err)
	}
	log.Printf("Float Res of %+v: %f", *req, r.ResFloat)
}
