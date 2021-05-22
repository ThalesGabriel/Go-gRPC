package services

import (
	"context"
	"fmt"
	"github/ThalesGabriel/go-grpc/pb/pb"
	"io"
	"log"
	"time"
)

// type UserServiceServer interface {
// 	AddUser(context.Context, *User) (*User, error)
// 	AddUserVerbose(*User, UserService_AddUserVerboseServer) error
// 	mustEmbedUnimplementedUserServiceServer()
//  AddUserStreamBoth(ctx context.Context, *User) error
// }

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	// Mock insert

	fmt.Println(req.GetName())

	return &pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil

}

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {

	stream.Send(&pb.UserResultStream{
		Status: "Inited",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Inserting",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "User has been inserted",
		User: &pb.User{
			Id:    "123",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Completed",
		User: &pb.User{
			Id:    "123",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})

	return nil

}

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {

	users := []*pb.User{}

	for {

		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}

		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}

		users = append(users, &pb.User{
			Id:    req.GetId(),
			Email: req.GetEmail(),
			Name:  req.GetName(),
		})

		fmt.Println("Adding: ", req.GetName())

	}

}

func (*UserService) AddUserStreamBoth(stream pb.UserService_AddUserStreamBothServer) error {

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error receiving stream from the client: %v", err)
		}

		err = stream.Send(&pb.UserResultStream{
			Status: "Added",
			User:   req,
		})

		if err != nil {
			log.Fatalf("Error sentind stream to the client: %v", err)
		}
	}

}
