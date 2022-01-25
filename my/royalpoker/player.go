package royalpoker

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Player interface {
	GetId() int
	Send(ctx context.Context, data []byte)
	Receive(ctx context.Context) []byte
	Close(ctx context.Context)
}
type PlayerWs struct {
	Id int
	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send    chan []byte
	receive chan []byte

	start chan struct{}
}

func NewPlayerWs(conn *websocket.Conn) *PlayerWs {
	player := &PlayerWs{
		Id:   0,
		conn: conn,
		send: make(chan []byte),
	}
	go func() {
		for {
			select {
			case msg := <-player.send:
				err := conn.WriteMessage(websocket.TextMessage, msg)
				logrus.Errorf("推送玩家消息[%s]错误：%s", msg, err.Error())
			}
		}
	}()

	go func() {
		for {
			select {
			case <-player.start:
				player.start <- struct{}{}
				_, msg, err := player.conn.ReadMessage()
				logrus.Errorf("主动接收玩家消息错误：%s", err.Error())
				player.receive <- msg
			default:
				_, _, err := player.conn.ReadMessage()
				logrus.Errorf("被动接收玩家消息错误：%s", err.Error())
			}
		}
	}()
	return player
}

func (self *PlayerWs) GetId() int {
	return self.Id
}

func (self *PlayerWs) Send(ctx context.Context, data []byte) {
	self.send <- data
}

func (self *PlayerWs) Receive(ctx context.Context) []byte {
	self.start <- struct{}{}
	defer func() {
		<-self.start
	}()
	return <-self.receive
}

func (self *PlayerWs) Close(ctx context.Context) {
	err := self.conn.Close()
	if err != nil {
		logrus.Errorf("关闭链接失败：%s", err.Error())
	}
}
