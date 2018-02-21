package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/pwelch/truck/protobuf"
	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewFileClient(conn)
	stream, err := c.FileTransfer(context.Background(), &pb.Request{Fetch: true})
	if err != nil {
		fmt.Println("error")
	}

	var content []byte
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("fail")
		}
		content = resp.GetContent()
	}

	fmt.Printf("%s \n", content)
}
