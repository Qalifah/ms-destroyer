package main

import (
	"context"
	"destroyer/database"
	"destroyer/handler"
	"destroyer/lib/pubsub/pulsar"
	"destroyer/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	pubSubSrv, err := pulsar.New("pulsar://localhost:6650", "targets.acquired")
	if err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()
	db, err := database.New(ctx, "mongodb://localhost:27017", "service")
	if err != nil {
		log.Fatalln(err)
	}
	ctrl := handler.New(pubSubSrv, db)
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterDestroyerServer(s, ctrl)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
