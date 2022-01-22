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

	action chan string
	// Buffered channel of outbound messages.
	//send chan []byte
}

func NewPlayer(conn *websocket.Conn) *Player {
	return &Player{
		Id:   0,
		conn: conn,
		//send: nil,
	}
}

func (self *Player) WaitAction(ctx context.Context) string {
	for {
		select {
		case action := <- self.action:
			return action
		case <- time.After(10*time.Second):
			// TODO[Dokiy] 2022/1/22: 定义一个包含当前玩家可以操作信息的结构体
			// TODO[Dokiy] 2022/1/22: 填写添加内容
			err := self.conn.WriteJSON(``)
			logrus.Errorf("推送玩家操作提醒错误：",err)
		}
	}
}
