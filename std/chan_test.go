package std

import (
	"testing"
	"time"
)

func TestChanOverAdd(t *testing.T) {
	var ch = make(chan struct{}, 2)
	ch <- struct{}{}
	ch <- struct{}{}
	ch <- struct{}{}
}

func TestName(t *testing.T) {
	receive := make(chan string)
	start := make(chan struct{})
	go func() {
		for {
			select {
			case msg := <-receive:
				t.Log(msg)
			case <-start:
				start <- struct{}{}
			}
		}
	}()

	f := func() string {
		start <- struct{}{}
		defer func() {
			<-start
		}()
		return <-receive
	}


	receive <- "123"

	go func() {
		time.Sleep(2*time.Second)
		receive <- "abc"
	}()

	t.Logf("f:%s",f())

	receive <- "123"
	select{}
}