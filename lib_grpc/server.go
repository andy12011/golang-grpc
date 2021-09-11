package lib_grpc

import (
	pb "golang-gRPC/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":8787"
)

type server struct {
	pb.UnimplementedSaveStudentDataServiceServer
}

func Server() {
	// Create gRPC Server

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Println("gRPC server is running.")

	pb.RegisterSaveStudentDataServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
