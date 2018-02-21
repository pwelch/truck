package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

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

	outFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("error creating file")
	}
	defer outFile.Close()

	w := bufio.NewWriter(outFile)

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("fail")
		}
		w.Write(resp.GetContent())
		w.Flush()
	}
}
