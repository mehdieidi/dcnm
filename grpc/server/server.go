package main

import (
	"context"
	"log"
	"net"

	pb "github.com/MehdiEidi/dcnm/grpc/keyvalue"
	"github.com/MehdiEidi/dcnm/internal/core"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedKeyValueServer
}

func (s *server) Get(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {
	log.Printf("Received GET key=%v", r.Key)

	value, err := core.Get(r.Key)

	return &pb.GetResponse{Value: value}, err
}

func (s *server) Put(ctx context.Context, r *pb.PutRequest) (*pb.PutResponse, error) {
	log.Printf("Received PUT key=%v value=%v", r.Key, r.Value)

	return &pb.PutResponse{}, core.Put(r.Key, r.Value)
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterKeyValueServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
