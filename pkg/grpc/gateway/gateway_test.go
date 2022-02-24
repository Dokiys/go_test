package hellogrpc

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"testing"
)

type Server struct {
	UnimplementedGreeterServer
}

// SayHello implements hellogrpc.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &HelloReply{Message: "Hello " + in.GetName()}, nil
}
func TestServer(t *testing.T) {
	var addr = "localhost:50052"
	var s *grpc.Server

	// Create the insecure server
	s = grpc.NewServer()
	RegisterGreeterServer(s, &Server{})

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func TestGateway(t *testing.T) {
	endpoint := "localhost:50052"
	gatewayMux := runtime.NewServeMux()
	dopts := []grpc.DialOption{grpc.WithInsecure()}
	err := RegisterGreeterHandlerFromEndpoint(context.Background(), gatewayMux, endpoint, dopts)
	if err != nil {
		t.Fatal(err)
	}

	err = http.ListenAndServe(":8081", gatewayMux)
	if err != nil {
		t.Fatal(err)
	}
}