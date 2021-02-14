package main

import (
	"context"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	"github.com/Jimeux/app-mesher/svc-pii/config"
	"github.com/Jimeux/app-mesher/svc-pii/rpc"
)

type server struct {
	rpc.UnimplementedPIIServiceServer
}

func (s *server) GetData(_ context.Context, in *rpc.GetDataRequest) (*rpc.GetDataReply, error) {
	log.Printf("Received request from ID:%d", in.GetId())
	return &rpc.GetDataReply{Data: "data:" + strconv.FormatInt(in.GetId(), 10)}, nil
}

func main() {
	conf := config.New()

	lis, err := net.Listen("tcp", conf.Server.Host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	rpc.RegisterPIIServiceServer(s, &server{})
	log.Printf("Listening on %s...\n", conf.Server.Host)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
