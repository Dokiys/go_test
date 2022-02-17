package main

import (
	"flag"
	"google.golang.org/grpc"
	"grpc/hellogrpc"
	"log"
	"net"
)

var addr = flag.String("addr", "localhost:50055", "http service address")

func main() {
	flag.Parse()
	var s *grpc.Server
	// Create the TLS credentials
	//{
	//	creds, err := credentials.NewServerTLSFromFile("../zchd.crt", "../ca.key")
	//	if err != nil {
	//		log.Fatalf("failed to new tls creds: %v", err)
	//	}
	//	s = grpc.NewServer(grpc.Creds(creds))
	//}

	// Create the insecure server
	{
		s = grpc.NewServer()
		hellogrpc.RegisterGreeterServer(s, &hellogrpc.Server{Addr: *addr})

		// 注册服务
		//addrM = make(map[string]string, 1)
		//addrM[myAddrKey] = addr
	}

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
