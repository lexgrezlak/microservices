package main

import (
	"context"
	"encoding/json"
	pb "github.com/vnqx/microservices/post-service/pb"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
)

const (
	address       = "localhost:3001"
	pathToPost = "data/post.json"
)

func parseFile(file string) (*pb.Post, error) {
	var post *pb.Post
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &post); err != nil {
		return nil, err
	}
	return post, nil
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewPostServiceClient(conn)

	// Contact the server and print out its response.
	post, err := parseFile(pathToPost)
	if err != nil {
		log.Fatalf("failed to parse file: %v", err)
	}

	r, err := client.CreatePost(context.Background(), post)
	if err != nil {
		log.Fatalf("failed to greet: %v", err)
	}
	log.Printf("created: %v", r.Created)
}
