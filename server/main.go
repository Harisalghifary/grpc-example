package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Harisalghifary/grpc-example/protofiles"
	"google.golang.org/grpc"
)

type server struct{}

func main() {

	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	pb.RegisterMoneyTransactionServer(grpcServer, &server{})

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to server: %v", err)

	}
}

// declare function from protobuf definition
func (s *server) MakeTransaction(ctx context.Context, input *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	// Business logic will come here
	fmt.Println("Got amount ", input.Amount)
	fmt.Println("Got from ", input.From)
	fmt.Println("For ", input.To)

	return &pb.TransactionResponse{Confirmation: true}, nil
}
