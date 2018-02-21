package main

import (
	"flag"
	"fmt"
	"net"

	pb "github.com/pwelch/truck/protobuf"
	grpc "google.golang.org/grpc"
)

// FileServer type
type FileServer struct{}

// FileTransfer transfer things
func (s *FileServer) FileTransfer(req *pb.Request, stream pb.File_FileTransferServer) error {
	fmt.Println("Entering FileTransfer")
	stream.Send(&pb.Response{Content: []byte("Our Message")})

	return nil
}

func newFileServer() *FileServer {
	return &FileServer{}
}

func main() {
	port := flag.Int("port", 50051, "Port for the server to run on")
	flag.Parse()

	fmt.Printf("Starting Server on port %d...\n", *port)
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	grpcServer := grpc.NewServer()
	pb.RegisterFileServer(grpcServer, newFileServer())
	grpcServer.Serve(conn)
}
