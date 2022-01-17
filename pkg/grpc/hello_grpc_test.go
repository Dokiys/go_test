package grpc

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc/hellogrpc"
	"io"
	"log"
	"net"
	"testing"
	"time"
)

const addr = "localhost:50055"

type server struct {
	pb.UnimplementedGreeterServer
}
// SayHello implements hellogrpc.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// SayMoreHello implements hellogrpc.GreeterServer
func (s *server) SayMoreHello(stream pb.Greeter_SayMoreHelloServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("Received: %v", in.GetName())
		reply := &pb.HelloReply{Message: "Hello " + in.GetName()}
		for i := 0; i < 3; i++ {
			if err := stream.Send(reply); err != nil {
				return err
			}
			time.Sleep(time.Second)
		}
	}

}

func TestGrpcServer(t *testing.T) {
	var s *grpc.Server
	// Create the TLS credentials
	//{
	//	creds, err := credentials.NewServerTLSFromFile("./hellogrpc/zchd.crt", "./hellogrpc/ca.key")
	//	if err != nil {
	//		log.Fatalf("failed to new tls creds: %v", err)
	//	}
	//	s = grpc.NewServer(grpc.Creds(creds))
	//}

	// Create the insecure server
	{
		s = grpc.NewServer()
	}
	pb.RegisterGreeterServer(s, &server{})
	
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func TestGrpcClient(t *testing.T) {
	var conn *grpc.ClientConn
	var err error
	//Set up a TLS connection to the server.
	//{
	//	creds, err := credentials.NewClientTLSFromFile("./hellogrpc/zchd.crt", "www.zchd.ltd")
	//	if err != nil {
	//		log.Fatalf("failed to new tls creds: %v", err)
	//	}
	//	conn, err = grpc.Dial(addr, grpc.WithTransportCredentials(creds))
	//}

	//Set up a connection to the server.
	{
		conn, err = grpc.Dial(addr, grpc.WithInsecure())
	}
	if err != nil {
		log.Fatalf("failed to Dial: %v", err)
	}
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	const name = "zhangsan"
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

func TestGrpcClient2(t *testing.T) {
	//Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to Dial: %v", err)
	}
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	stream, err := c.SayMoreHello(context.Background())
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			log.Printf("Greeting: %s", in.GetMessage())
		}
	}()

	names := []string{"zhangsan", "lisi", "wangwu"}
	for _, name := range names {
		if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
			log.Fatalf("Failed to send a req: %v", err)
		}
		time.Sleep(time.Second)
	}
	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("could not stop greet: %v", err)
	}
	<-waitc
}
