package crontask

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"time"
)

type EtcdMutex struct {
	ttl int
	l   *concurrency.Mutex
}

func NewEtcdMutex(client *clientv3.Client, ttl int, key string) (*EtcdMutex, error) {
	s, err := concurrency.NewSession(client, concurrency.WithTTL(ttl))
	if err != nil {
		return nil, errors.Wrapf(err, "EtcdMutex create etcd session err")
	}
	l := concurrency.NewMutex(s, "/"+key)
	return &EtcdMutex{ttl: ttl, l: l}, nil
}

func (self *EtcdMutex) Lock(ctx context.Context) error {
	return errors.Wrapf(self.l.Lock(ctx), "EtcdMutex lock err")
}

func (self *EtcdMutex) Unlock(ctx context.Context) error {
	return errors.Wrapf(self.l.Unlock(ctx), "EtcdMutex unlock err")
}

func (self *EtcdMutex) KeepLock(ctx context.Context) {
	for {
		select {
		case <-time.Tick(time.Duration(self.ttl*3/4) * time.Second):
			if err := self.l.TryLock(ctx); err != nil {
				logrus.Errorf("Fail to keep EtcdMutex lock, %v", err)
				return
			}
		case <-ctx.Done():
			return
		}
	}
}