package pkg

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"testing"
	"time"
)

func TestEtcdClient(t *testing.T) {
	ctx := context.Background()
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	defer cli.Close()

	s, err := concurrency.NewSession(cli, concurrency.WithTTL(10))
	if err != nil {
		//
	}
	l := concurrency.NewMutex(s, "/hello_mutex")
	l.Lock(ctx)
	time.Sleep(time.Second*5)
	l.Unlock(ctx)
	select{}
}


