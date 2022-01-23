package common

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"time"
)

type Player struct {
	Id int
	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send   chan []byte
	recive chan []byte
}

func NewPlayer(conn *websocket.Conn) *Player {
	player := &Player{
		Id:   0,
		conn: conn,
		send: make(chan []byte),
	}
	go func() {
		for {
			select {
			case msg := <-player.send:
				err := player.conn.WriteJSON(msg)
				logrus.Errorf("推送玩家消息[%s]错误：%s", msg, err.Error())
			}
		}
	}()
	return player
}
func (self *Player) Send(data []byte) {
	self.send <- data
}

func (self *Player) WaitAction(ctx context.Context, data []byte) []byte {
	self.Send(data)

	self.recive = make(chan []byte)
	for {
		select {
		case action := <-self.recive:
			close(self.recive)
			return action
		case <-time.After(10 * time.Second):
			self.Send(data)
		}
	}
}
