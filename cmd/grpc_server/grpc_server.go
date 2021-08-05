package grpc_server

import (
	"context"
	"final/internal/message_sender"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"final/config"
	pb "final/grpc_service"
)

type server struct {
	pb.UnimplementedMessageServiceServer
	sender *message_sender.MsgSender
}

func (s *server) SendAction(_ context.Context, in *pb.Action) (*pb.Result, error) {
	log.Printf("Received action from client: %s\n", in.Type)

	s.sender.Send(in)

	return &pb.Result{Success: true}, nil
}

func StartGRPCServer(config config.Config, sender *message_sender.MsgSender) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.GRPC.Addr, config.GRPC.Port))
	if err != nil {
		log.Fatal("grpc server listen", err)
	}

	s := grpc.NewServer()

	pb.RegisterMessageServiceServer(s, &server{
		sender: sender,
	})

	defer s.Stop()

	if err := s.Serve(lis); err != nil {
		log.Fatal("grpc serve", err)
	} else {
		fmt.Println("started")
	}
}
