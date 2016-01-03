package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/awakia/architecture2016/src/rpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	port = flag.Int("port", 9090, "The server port")
)

type server struct{}

func (s *server) List(_ *pb.Empty, stream pb.UserService_ListServer) error {
	users := []*pb.User{
		&pb.User{Id: 1, Name: "Naoyoshi Aikawa"},
		&pb.User{Id: 2, Name: "Hoge Fuga"},
	}
	for _, user := range users {
		if err := stream.Send(user); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) Get(ctx context.Context, param *pb.UserParam) (*pb.User, error) {
	return &pb.User{Id: 1, Name: "Naoyoshi Aikawa"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})
	grpcServer.Serve(lis)
	log.Printf("server running on port:%d", *port)
}
