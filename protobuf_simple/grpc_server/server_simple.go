/*
1. write a proto file
/types/person.proto
2. generate go file for grpc
protoc --go_out=plugins=grpc:. ./types/client_server.proto
3. run server
go run server_simple.go
4. run client
go run client_simple.go
*/
package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/SUN-XIN/go-example/protobuf_simple/types"
)

type server struct {
	StartDate int64
}

func NewServer() *server {
	return &server{
		StartDate: time.Now().Unix(),
	}
}

func (s *server) Sum(ctx context.Context, in *types.SumRequest) (*types.SumReply, error) {
	return &types.SumReply{
		ResInt:    in.ValInt1 + in.ValInt2,
		ResString: fmt.Sprintf("%s%s", in.ValString1, in.ValString2),
		ResFloat:  in.ValFloat1 + in.ValFloat2,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	serv := NewServer()
	types.RegisterSumServer(s, serv)
	s.Serve(lis)
}
