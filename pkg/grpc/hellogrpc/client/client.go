package main

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"grpc/hellogrpc"
	"log"
)

func init() {
	resolver.Register(&MyResolveBuilder{})
	balancer.Register(newMyBalanceBuilder())
}
func main() {
	var conn *grpc.ClientConn
	var err error

	conn, err = grpc.Dial("myresolve:///mytarget",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"mybalancer"}`),
		//grpc.WithBalancerName("mybalancer"),			// same as above
	)
	if err != nil {
		log.Fatalf("failed to Dial: %v", err)
	}
	c := hellogrpc.NewGreeterClient(conn)

	for i := 0; i < 5; i++ {
		r, err := c.SayHello(context.Background(), &hellogrpc.HelloRequest{Name: "zhangsan"})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	}
}

/*
	Server discovery
*/
const scheme = "myresolve"

// Implement ResolveBuilder
type MyResolveBuilder struct{}

func (self *MyResolveBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	// 发现服务
	//state = resolver.State{
	//	Addresses: []resolver.Address{{Addr: addrM[target.Endpoint]}},
	//}
	var state resolver.State
	if target.Endpoint == "mytarget" {
		state = resolver.State{
			Addresses: []resolver.Address{{Addr: "localhost:50055"}, {Addr: "localhost:50056"}},
		}
	}
	err := cc.UpdateState(state)
	if err != nil {
		cc.ReportError(errors.Wrapf(err, "更新State失败："))
	}
	return &MyResolver{}, nil
}
func (self *MyResolveBuilder) Scheme() string { return scheme }

// Implement Resolver
type MyResolver struct{}

func (self *MyResolver) ResolveNow(resolver.ResolveNowOptions) {}
func (self *MyResolver) Close()                                {}

/*
	Load Balancer
*/
const lbName = "mybalancer"

func newMyBalanceBuilder() balancer.Builder {
	return base.NewBalancerBuilder(lbName, &MyPickerBuilder{}, base.Config{HealthCheck: true})
}

// Implement PickerBuilder
type MyPickerBuilder struct{}

func (self *MyPickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	scs := make([]balancer.SubConn, 0, len(info.ReadySCs))
	for sc := range info.ReadySCs {
		scs = append(scs, sc)
	}
	return &MyPicker{subConns: scs}
}

// Implement Picker
type MyPicker struct {
	subConns []balancer.SubConn
	next     int
}

func (self *MyPicker) Pick(_ balancer.PickInfo) (balancer.PickResult, error) {
	sc := self.subConns[self.next]
	self.next = (self.next + 1) % len(self.subConns)
	return balancer.PickResult{SubConn: sc}, nil
}
