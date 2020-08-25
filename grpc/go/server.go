package main

import (
  "fmt"
  "log"
  "net"
  "google.golang.org/grpc"
  "./chat"
)

func main() {

  fmt.Println("Go gRPC Beginners Tutorial!")

  lis, err := net.Listen("tcp", ":9000")
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }

  grpcServer := grpc.NewServer()

  s := chat.Server{}

  chat.RegisterChatServiceServer(grpcServer, &s)

  if err := grpcServer.Serve(lis); err != nil {
	log.Fatalf("failed to serve: %s", err)
  }
}
