package main

import (
	"context"
	"log"
	"time"

	pb "grpc-user-service/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Test GetUser
	r, err := c.GetUser(ctx, &pb.UserIdRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("User: %v", r)

	// Test GetUsers
	ids := []int32{1, 2}
	users, err := c.GetUsers(ctx, &pb.UserIdsRequest{Ids: ids})
	if err != nil {
		log.Fatalf("could not get users: %v", err)
	}
	log.Printf("Users: %v", users.Users)

	// Test SearchUsers
	searchReq := &pb.UserSearchRequest{Fname: "Steve", City: "LA"}
	searchRes, err := c.SearchUsers(ctx, searchReq)
	if err != nil {
		log.Fatalf("could not search users: %v", err)
	}
	log.Printf("Search Result: %v", searchRes.Users)
}
