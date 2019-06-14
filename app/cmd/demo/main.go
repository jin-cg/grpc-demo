package main

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jin-cg/grpc-demo/app/demo/controllers"
	"github.com/jin-cg/grpc-demo/app/demo/proto"
	"github.com/jin-cg/grpc-demo/app/demo/services"
	"github.com/jin-cg/grpc-demo/app/pkg/grpc/interceptor"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	addr := fmt.Sprintf(":%d", 50051)
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Cannot listen to address %s", addr)
	}
	db, err := gorm.Open("postgres", nil)
	if err != nil {
		panic(err)
	}

	service := services.NewService(db)
	controller := controllers.NewController(service)
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			interceptor.DBInterceptor(db),
		)),
	)
	proto.RegisterHelloWorldServiceServer(server, controller)

	if err := server.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
