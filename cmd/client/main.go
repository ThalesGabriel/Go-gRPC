package main

import (
	"context"
	"fmt"
	"github/ThalesGabriel/go-grpc/pb/pb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

var reqs = []*pb.User{
	{
		Id:    "123",
		Name:  "Thales 1",
		Email: "thales@exemplo1.com",
	},
	{
		Id:    "1234",
		Name:  "Thales 2",
		Email: "thales@exemplo4.com",
	},
	{
		Id:    "1235",
		Name:  "Thales 3",
		Email: "thales@exemplo3.com",
	},
	{
		Id:    "1236",
		Name:  "Thales 4",
		Email: "thales@exemplo4.com",
	},
	{
		Id:    "1237",
		Name:  "Thales 5",
		Email: "thales@exemplo5.com",
	},
}

func main() {

	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect with gRPC %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	// AddUser(client)
	// AddUserVerbose(client)
	// AddUsers(client)
	AddUserStreamBoth(client)
}

func AddUser(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "0",
		Name:  "Thales",
		Email: "thal3s49@hotmail.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not add User %v", err)
	}

	fmt.Println(res)

}

func AddUserVerbose(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "0",
		Name:  "Thales",
		Email: "thal3s49@hotmail.com",
	}

	restream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not add User with Verbose %v", err)
	}
	for {

		stream, err := restream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not receive message %v", err)
		}

		fmt.Println("Status: ", stream.GetStatus())

	}

}

func AddUsers(client pb.UserServiceClient) {

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)

}

func AddUserStreamBoth(client pb.UserServiceClient) {

	stream, err := client.AddUserStreamBoth(context.Background())

	if err != nil {
		log.Fatalf("Could creating request %v", err)
	}

	channel := make(chan int)

	go func() {

		for _, req := range reqs {
			fmt.Println("Sending user: ", req.GetName())
			stream.Send(req)
			time.Sleep(time.Second * 3)
		}
		stream.CloseSend()

	}()

	go func() {

		for {

			res, erro := stream.Recv()

			if erro == io.EOF {
				break
			}

			if erro != nil {
				log.Fatalf("Could not receive message %v", err)
				break
			}

			fmt.Printf("Receiving user %v with status %v\n", res.GetUser().GetName(), res.GetStatus())

		}

	}()

	<-channel

}
