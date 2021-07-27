package grpcServer

import (
	"context"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"final/config"
	pb "final/grpc"
)

type server struct {
	pb.UnimplementedMessageServiceServer
}

func (s *server) SendAction(_ context.Context, in *pb.Action) (*pb.Result, error) {
	logrus.Infof("Received action from client: %s", in.Type)

	return &pb.Result{Success: true}, nil
}

func StartGRPCServer(config config.Config) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.GRPC.Addr, config.GRPC.Port))
	if err != nil {
		logrus.Fatal(err)
	}

	s := grpc.NewServer()

	pb.RegisterMessageServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		logrus.Fatal(err)
	} else {
		fmt.Println("started")
	}
}