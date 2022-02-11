package frontend

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/MehdiEidi/dcnm/core"
	pb "github.com/MehdiEidi/dcnm/grpc/keyvalue"
	"google.golang.org/grpc"
)

type grpcFrontEnd struct {
	store *core.KeyValueStore
	pb.UnimplementedKeyValueServer
}

func (s *grpcFrontEnd) Start(store *core.KeyValueStore) error {
	s.store = store

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	gs := grpc.NewServer()

	pb.RegisterKeyValueServer(gs, &grpcFrontEnd{})
	if err := gs.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

func (s *grpcFrontEnd) Get(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {
	log.Printf("Received GET key=%v", r.Key)

	value, err := s.store.Get(r.Key)

	return &pb.GetResponse{Value: value}, err
}

func (s *grpcFrontEnd) Put(ctx context.Context, r *pb.PutRequest) (*pb.PutResponse, error) {
	log.Printf("Received PUT key=%v value=%v", r.Key, r.Value)

	return &pb.PutResponse{}, s.store.Put(r.Key, r.Value, false)
}
