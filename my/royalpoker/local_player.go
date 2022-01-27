package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/dokiy/royalpoker/common"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var l sync.Mutex

type LocalPlayer struct {
	Id   int
	Name string

	// Buffered channel of outbound messages.
	send    chan []byte
	receive chan []byte
	close   chan struct{}
	start   chan struct{}
}

func NewLocalPlayer(name string) *LocalPlayer {
	player := &LocalPlayer{
		Id:      common.RandNum(6),
		Name:    name,
		send:    make(chan []byte),
		receive: make(chan []byte),
		start:   make(chan struct{}),
		close:   make(chan struct{}),
	}
	go func() {
		for {
			select {
			case msg := <-player.send:
				// 写出数据到控制台
				logrus.Infof("对玩家[%d]发出消息：%s", player.Id, msg)
			case <-player.close:
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case <-player.start:
				l.Lock()
				print(fmt.Sprintf("请对玩家[%d]进行操作\n", player.Id))
				var msg string
				reader := bufio.NewReader(os.Stdin)
				{
					data, _, _ := reader.ReadLine()
					msg = string(data)
				}
				//msg := `{"action_type":"W3C_ACTION_READY"}`
				//logrus.Infof("接收到玩家[%d]消息：%s", player.Id, msg)
				print("输入结束\n")

				player.receive <- []byte(msg)
				player.start <- struct{}{}
				l.Unlock()

			case <-player.close:
				return
			}
		}
	}()
	return player
}

func (self *LocalPlayer) GetId() int {
	return self.Id
}

func (self *LocalPlayer) GetName() string {
	return self.Name
}

func (self *LocalPlayer) Send(ctx context.Context, data []byte) {
	self.send <- data
}

func (self *LocalPlayer) Receive(ctx context.Context) ([]byte, error) {
	self.start <- struct{}{}
	defer func() {
		<-self.start
	}()
	return <-self.receive, nil
}

func (self *LocalPlayer) Close(ctx context.Context) {
	self.close <- struct{}{}
}
