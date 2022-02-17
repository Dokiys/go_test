package hellogrpc

import (
	"context"
	"io"
	"log"
	"time"
)

type Server struct {
	Addr string
	UnimplementedGreeterServer
}

// SayHello implements hellogrpc.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &HelloReply{Message: s.Addr + ":Hello " + in.GetName()}, nil
}

// SayMoreHello implements hellogrpc.GreeterServer
func (s *Server) SayMoreHello(stream Greeter_SayMoreHelloServer) error {
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
