package hellogrpc

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	etcd "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"testing"
	"time"
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

	err := registerServer(context.Background())
	if err != nil {
		log.Fatalf("register server err:%s", err)
	}
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

func registerServer(ctx context.Context) error {
	client, err := etcd.New(etcd.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		return err
	}

	manager, err := endpoints.NewManager(client, "grpc/hello_etcd")
	if err != nil {
		return err
	}

	lease := etcd.NewLease(client)
	leaseResp, err := lease.Grant(ctx, 30)
	if err != nil {
		return errors.Wrapf(err, "EtcdClient获取租约失败！")
	}
	go func() {
		alive, _ := lease.KeepAlive(ctx, leaseResp.ID)
		for {
			select {
			case <- alive:
			}
		}
	}()

	err = manager.AddEndpoint(ctx, "grpc/hello_etcd/localhost:50052", endpoints.Endpoint{Addr: "localhost:50052"}, etcd.WithLease(leaseResp.ID))
	if err != nil {
		return err
	}

	return nil
}

func TestGateway(t *testing.T) {
	etcdClient, err := etcd.New(etcd.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: time.Second * 5,
	})
	builder, err := resolver.NewBuilder(etcdClient)
	if err != nil {
		log.Fatal(err)
	}

	endpoint := "etcd:///grpc/hello_etcd"

	gatewayMux := runtime.NewServeMux()
	dopts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(builder),
	}
	err = RegisterGreeterHandlerFromEndpoint(context.Background(), gatewayMux, endpoint, dopts)
	if err != nil {
		t.Fatal(err)
	}

	err = http.ListenAndServe(":8081", gatewayMux)
	if err != nil {
		t.Fatal(err)
	}
}