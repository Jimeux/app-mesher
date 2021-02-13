package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/Jimeux/app-mesher/svc-identity/config"
	"github.com/Jimeux/app-mesher/svc-identity/rpc"
)

type server struct {
	rpc.UnimplementedIdentityServiceServer
}

func (s *server) IssueToken(_ context.Context, in *rpc.IssueTokenRequest) (*rpc.IssueTokenReply, error) {
	log.Printf("Received request from: %s", in.GetUsername())
	return &rpc.IssueTokenReply{Token: "token:" + in.GetUsername()}, nil
}

func main() {
	conf := config.New()

	lis, err := net.Listen("tcp", conf.Server.Host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	rpc.RegisterIdentityServiceServer(s, &server{})
	log.Printf("Listening on %s...\n", conf.Server.Host)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
