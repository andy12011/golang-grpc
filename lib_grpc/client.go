package lib_grpc

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "golang-gRPC/proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8787"
)

func Client() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Sum(ctx, &pb.CalculatorRequest{A: 100, B: 200})
	if err != nil {
		log.Fatalf("Could not get nonce: %v", err)
	}

	fmt.Println("Response:", r.GetResult())
}
