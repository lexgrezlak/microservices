package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "instagram-microservices/src/post-service/pb"
	"log"
	"net"
	"sync"
)

const (
	port = ":3001"
)

type repository interface {
	Create(*pb.Post) (*pb.Post, error)
}

type Repository struct {
	mu sync.RWMutex
	posts []*pb.Post
}

func (r *Repository) Create(post *pb.Post) (*pb.Post, error)  {
	r.mu.Lock()
	updatedPosts := append(r.posts, post)
	r.posts = updatedPosts
	r.mu.Unlock()
	return post, nil
}

type service struct {
	r *Repository
}

func (s *service) CreatePost(ctx context.Context, req *pb.Post) (*pb.Response, error) {
	post, err := s.r.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Created: true,
		Post:    post,
	}, nil
}

func main() {
	// Set up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	r := &Repository{}
	pb.RegisterPostServiceServer(s, &service{r})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("Running on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}