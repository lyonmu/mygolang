package main

import (
	"context"
	"flag"
	"log"
	"time"

	proto "testgrpc/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ConToPro(ctx, &proto.MyConsumer{Mmber: 23})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("收到的数字为: %v", r.GetNnber())
}
