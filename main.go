package main

import (
	"fmt"
	"golang-gRPC/lib_grpc"
	pb "golang-gRPC/proto"
)

func main() {
	s := &pb.Student{
		Name:   "Peng Jie",
		Age:    24,
		Gender: "Male",
		Number: 99,
	}

	fmt.Println(
		s.GetName(),
		s.GetAge(),
		s.GetGender(),
		s.GetNumber(),
	)

	lib_grpc.Server()
}
