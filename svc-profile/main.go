package main

import (
	"context"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	"github.com/Jimeux/app-mesher/svc-profile/config"
	"github.com/Jimeux/app-mesher/svc-profile/rpc"
)

type server struct {
	rpc.UnimplementedProfileServiceServer
	piiSvc rpc.PIIServiceClient
}

func (s *server) GetProfile(ctx context.Context, in *rpc.GetProfileRequest) (*rpc.GetProfileReply, error) {
	log.Printf("Received request from ID:%d", in.GetId())
	piiReply, err := s.piiSvc.GetData(ctx, &rpc.GetDataRequest{
		Id: in.GetId(),
	})
	if err != nil {
		return nil, err
	}
	return &rpc.GetProfileReply{
		Profile: "profile:" + strconv.FormatInt(in.GetId(), 10),
		Data:    piiReply.GetData(),
	}, nil
}

func main() {
	conf := config.New()

	piiConn, err := grpc.Dial(conf.Server.PIIHost,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("could not connect to profile svc: %v", err)
	}
	defer piiConn.Close()

	piiSvc := rpc.NewPIIServiceClient(piiConn)
	profileServer := &server{piiSvc: piiSvc}

	lis, err := net.Listen("tcp", conf.Server.Host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	rpc.RegisterProfileServiceServer(s, profileServer)
	log.Printf("Listening on %s...\n", conf.Server.Host)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
