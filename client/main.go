package main

import (
	"context"
	"log"

	pb "github.com/Harisalghifary/grpc-example/protofiles"
	"google.golang.org/grpc"
)

const address = "localhost:8000"

func main() {
	// setup connection with grpc server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error while making connection: %v", err)
	}

	// create client instance
	clientServer := pb.NewMoneyTransactionClient(conn)

	// invoke remote func from client to server
	clientServer.MakeTransaction(context.Background(), &pb.TransactionRequest{
		From:   "masamune",
		To:     "xbox",
		Amount: float32(50.5),
	})

}
