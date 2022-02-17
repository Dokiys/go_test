package hellogrpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net"
	"testing"
	"time"
)

const addr = "localhost:50055"

type server struct {
	addr string
	UnimplementedGreeterServer
}

// SayHello implements hellogrpc.GreeterServer
func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &HelloReply{Message: s.addr + ":Hello " + in.GetName()}, nil
}

// SayMoreHello implements hellogrpc.GreeterServer
func (s *server) SayMoreHello(stream Greeter_SayMoreHelloServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("Received: %v", in.GetName())
		reply := &HelloReply{Message: "Hello " + in.GetName()}
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
		RegisterGreeterServer(s, &Server{Addr: addr})

		// 注册服务
		//addrM = make(map[string]string, 1)
		//addrM[myAddrKey] = addr
	}

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
		conn, err = grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	if err != nil {
		log.Fatalf("failed to Dial: %v", err)
	}
	c := NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	const name = "zhangsan"
	r, err := c.SayHello(ctx, &HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

func TestGrpcStreamClient(t *testing.T) {
	//Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to Dial: %v", err)
	}
	c := NewGreeterClient(conn)

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
		if err := stream.Send(&HelloRequest{Name: name}); err != nil {
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