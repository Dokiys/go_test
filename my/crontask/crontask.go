package crontask

import (
	"context"
	"github.com/pkg/errors"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type Mutex interface {
	Lock(ctx context.Context) error
	Unlock(ctx context.Context) error
	KeepLock(ctx context.Context)
}

type CronTask struct {
	c   *cron.Cron
	l   Mutex
	ctx context.Context
}

func New(ctx context.Context, mutex Mutex) CronTask {
	return CronTask{c: cron.New(), ctx: ctx, l: mutex}
}

func (self *CronTask) AddTask(spec string, f func()) *CronTask {
	err := self.c.AddFunc(spec, f)
	if err != nil {
		panic(errors.Wrapf(err, "CronTask add task fail"))
	}
	return self
}

func (self *CronTask) Run() {
	rand.Seed(time.Now().UnixNano())

	go func() {
		for {
			select {
			case <-time.Tick(time.Duration(rand.Int31n(5)) * time.Second):
				err := self.run()
				if err != nil {
					logrus.Errorf("CronTask run err: %s", err.Error())
				}
			case <- self.ctx.Done():
				logrus.Errorf("CronTask context break")
				return
			}
		}
	}()
}

func (self *CronTask) run() error {
	err := self.l.Lock(self.ctx)
	if err != nil {
		return errors.Wrapf(err, "Lock fail")
	}

	self.c.Start()
	self.l.KeepLock(self.ctx)
	self.c.Stop()

	return errors.Wrapf(self.l.Unlock(self.ctx), "CronTask Unlock err")
}
