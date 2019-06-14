package main

import (
	"context"
	"fmt"
	"github.com/jin-cg/grpc-demo/app/demo/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewHelloWorldServiceClient(conn)

	resp, _ := c.Hei(
		context.Background(),
		&proto.Request{Name: "Euler"},
	)
	fmt.Println("NormalCase:", resp.GetContent())
}
