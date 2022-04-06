package crontask

import (
	"context"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

func TestCronTask(t *testing.T) {
	f("123")
}

func TestCronTask2(t *testing.T) {
	f("abc")
}

func TestCronTask3(t *testing.T) {
	f("xyz")
}

func TestCronTask4(t *testing.T) {
	f("+-*/")
}
func f(str string) {
	ctx := context.Background()
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	mutex, err := NewEtcdMutex(cli, 10, "test")
	if err != nil {
		panic(err)
	}

	ct := New(ctx, mutex)
	ct.AddTask("*/5 * * * * ?", func() { logrus.Info(str) })
	ct.Run()

	select {}
}
