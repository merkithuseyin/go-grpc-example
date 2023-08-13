package main

import (
	"context"
	"fmt"
	"time"

	pb "grpc-example/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	const myRequestParameter = "merkit"
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: myRequestParameter})
	if err != nil {
		fmt.Printf("could not get response: %v\r\n", err)
	}
	fmt.Printf("Response: %s\r\n", r.GetMessage())

	// call another function
	r, err = c.SayHelloAgain(ctx, &pb.HelloRequest{Name: myRequestParameter})
	if err != nil {
		fmt.Printf("could not get response: %v\r\n", err)
	}
	fmt.Printf("Response: %s\r\n", r.GetMessage())
}
