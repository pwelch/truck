package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"net"
	"os"

	pb "github.com/pwelch/truck/protobuf"
	grpc "google.golang.org/grpc"
)

// bufferSize for file transfer, default 0.25 MB
const bufferSize = 262144

// FileServer type
type FileServer struct{}

// FileTransfer transfer things
func (s *FileServer) FileTransfer(req *pb.Request, stream pb.File_FileTransferServer) error {
	fmt.Println("Entering FileTransfer")

	// Open our file to transfer
	file, err := os.Open("./fixtures/time_machine.txt")
	if err != nil {
		fmt.Println("err reading file")
	}
	defer file.Close()

	// Create our buffer to chunk the file
	buffer := make([]byte, bufferSize)

	// Read the file until we hit EOF
	for {
		// Read up to len(buffer) bytes from the file each pass
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}

			break
		}

		// Send this chunk of bytes.
		// buffer[:bytesRead] ensures we send the remaining bytes if buffer is not full
		stream.Send(&pb.Response{Content: buffer[:bytesRead]})
	}

	b2 := make([]byte, bufferSize)
	f2, err := os.Open("./fixtures/time_machine.txt")
	if err != nil {
		fmt.Println("err reading file")
	}
	defer f2.Close()
	h := sha256.New()
	if _, err := io.CopyBuffer(h, f2, b2); err != nil {
		fmt.Println("error creating sha")
	}

	fmt.Printf("\n %x \n", h.Sum(nil))

	fmt.Println("Finished FileTransfer")

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
