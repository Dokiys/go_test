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
	Players   map[int]PlayerMsg `json:"players"`
}

type PlayerMsg struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// =========================================================

func GenHubSessionMsg(hub *Hub, msg string) []byte {
	playerMsg := make(map[int]PlayerMsg, len(hub.Players))
	for id, player := range hub.Players {
		playerMsg[id] = PlayerMsg{
			Id:   player.GetId(),
			Name: player.GetName(),
		}
	}
	act := HubSessionMsg{
		Type: MSGTYPE_HUB_SESSION,
		Msg:  msg,
		Data: HubSessionMsgData{
			Owner:     hub.Owner,
			Players:   playerMsg,
		},
	}
	bytes, err := json.Marshal(act)
	if err != nil {
		logrus.Errorf("序列化ActionViewMsg失败: ", err.Error())
	}

	return bytes
}
