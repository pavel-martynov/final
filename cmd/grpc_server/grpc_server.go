package grpc_server

import (
	"context"
	"final/internal/message_sender"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"final/config"
	pb "final/grpc_service"
)

type server struct {
	pb.UnimplementedMessageServiceServer
	sender *message_sender.MsgSender
}

func (s *server) SendAction(_ context.Context, in *pb.Action) (*pb.Result, error) {
	logrus.Infof("Received action from client: %s", in.Type)

	s.sender.Send(in)

	return &pb.Result{Success: true}, nil
}

func StartGRPCServer(config config.Config, sender *message_sender.MsgSender) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.GRPC.Addr, config.GRPC.Port))
	if err != nil {
		logrus.Fatal(err)
	}

	s := grpc.NewServer()

	pb.RegisterMessageServiceServer(s, &server{
		sender: sender,
	})

	defer s.Stop()

	if err := s.Serve(lis); err != nil {
		logrus.Fatal(err)
	} else {
		fmt.Println("started")
	}
}
