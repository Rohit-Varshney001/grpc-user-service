package main

import (
	"context"
	"log"
	"net"

	pb "grpc-user-service/pb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
	users []*pb.User
}

func (s *server) GetUser(ctx context.Context, req *pb.UserIdRequest) (*pb.User, error) {
	for _, user := range s.users {
		if user.Id == req.Id {
			return user, nil
		}
	}
	return nil, nil
}

func (s *server) GetUsers(ctx context.Context, req *pb.UserIdsRequest) (*pb.Users, error) {
	var users []*pb.User
	for _, id := range req.Ids {
		for _, user := range s.users {
			if user.Id == id {
				users = append(users, user)
				break
			}
		}
	}
	return &pb.Users{Users: users}, nil
}

func (s *server) SearchUsers(ctx context.Context, req *pb.UserSearchRequest) (*pb.Users, error) {
	var users []*pb.User
	for _, user := range s.users {
		if (req.Fname == "" || user.Fname == req.Fname) && (req.City == "" || user.City == req.City) {
			users = append(users, user)
		}
	}
	return &pb.Users{Users: users}, nil
}

func main() {
	users := []*pb.User{
		{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
		{Id: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 6.1, Married: false},
		// Add more users as needed
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{users: users})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
