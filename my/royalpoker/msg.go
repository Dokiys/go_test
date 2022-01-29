package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

const (
	MSGTYPE_HUB_SESSION = "HUB_SESSION"
)

type MsgType string
type HubSessionMsg struct {
	Type MsgType           `json:"type"`
	Msg  string            `json:"msg"`
	Data HubSessionMsgData `json:"data"`
}

type HubSessionMsgData struct {
	Owner     int         `json:"owner"`
	Players   []PlayerMsg `json:"players"`
	IsStarted bool        `json:"is_started"`
}

type PlayerMsg struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// =========================================================

func GenHubSessionMsg(hub *Hub, msg string) []byte {
	playerMsg := make([]PlayerMsg, len(hub.Players))
	var i = 0
	for _, player := range hub.Players {
		playerMsg[i] = PlayerMsg{
			Id:   player.GetId(),
			Name: player.GetName(),
		}
		i++
	}
	act := HubSessionMsg{
		Type: MSGTYPE_HUB_SESSION,
		Msg:  msg,
		Data: HubSessionMsgData{
			Owner:     hub.Owner,
			Players:   playerMsg,
			IsStarted: hub.IsStarted,
		},
	}
	bytes, err := json.Marshal(act)
	if err != nil {
		logrus.Errorf("序列化ActionViewMsg失败: ", err.Error())
	}

	return bytes
}
