package lib_grpc

import (
	"context"
	pb "golang-gRPC/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8787"
)

type Server struct {
	pb.CalculatorServiceServer
}

func RunServer() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}

func (s *Server) Sum(_ context.Context, c *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {

	resp := &pb.CalculatorResponse{}
	resp.Result = c.A + c.B

	return resp, nil
}
