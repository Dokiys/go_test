package main

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Player interface {
	GetId() int
	GetName() string
	Send(ctx context.Context, data []byte)
	Receive(ctx context.Context) ([]byte,error)
	Close(ctx context.Context)
}
type PlayerWs struct {
	Id   int
	Name string
	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send    chan []byte
	receive chan []byte
	close   chan struct{}

	start chan struct{}
}

func NewPlayerWs(conn *websocket.Conn, id int, name string) *PlayerWs {
	player := &PlayerWs{
		Id:      id,
		Name:    name,
		conn:    conn,
		send:    make(chan []byte),
		receive: make(chan []byte),
		start:   make(chan struct{}),
		close:   make(chan struct{}),
	}
	go func() {
		for {
			select {
			case msg := <-player.send:
				err := conn.WriteMessage(websocket.TextMessage, msg)
				logrus.Errorf("推送玩家消息[%s]错误：%s", msg, err.Error())
			case <-player.close:
				close(player.start)
				close(player.close)
				close(player.start)
				close(player.receive)
				player.conn.Close()
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case <-player.start:
				_, msg, err := player.conn.ReadMessage()
				if err != nil {
					logrus.Errorf("主动接收玩家消息错误：%s", err.Error())
				}
				player.receive <- msg
				player.start <- struct{}{}

			case <-player.close:
				close(player.start)
				close(player.close)
				close(player.start)
				close(player.receive)
				player.conn.Close()
				return
			}
		}
	}()
	return player
}

func (self *PlayerWs) GetId() int {
	return self.Id
}

func (self *PlayerWs) GetName() string {
	return self.Name
}

func (self *PlayerWs) Send(ctx context.Context, data []byte) {
	self.send <- data
}

func (self *PlayerWs) Receive(ctx context.Context) ([]byte, error) {
	self.start <- struct{}{}
	defer func() {
		<-self.start
	}()

	select {
	case data := <-self.receive:
		return data, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (self *PlayerWs) Close(ctx context.Context) {
	self.close <- struct{}{}
	err := self.conn.Close()
	if err != nil {
		logrus.Errorf("关闭链接失败：%s", err.Error())
	}
}
