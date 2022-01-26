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
	Type MsgType
	Msg  string
	Data HubSessionMsgData
}

type HubSessionMsgData struct {
	Owner   int
	Players []PlayerMsg
}

type PlayerMsg struct {
	Id   int
	Name string
}

// =========================================================

func GenHubSessionMsg(hub *Hub, msg string) []byte {
	playerMsg := make([]PlayerMsg, len(hub.Players))
	for i, player := range hub.Players {
		playerMsg[i] = PlayerMsg{
			Id:   player.GetId(),
			Name: player.GetName(),
		}
	}
	act := HubSessionMsg{
		Type: MSGTYPE_HUB_SESSION,
		Msg:  msg,
		Data: HubSessionMsgData{
			Owner:   hub.Owner,
			Players: playerMsg,
		},
	}
	bytes, err := json.Marshal(act)
	if err != nil {
		logrus.Errorf("序列化ActionViewMsg失败: ", err.Error())
	}

	return bytes
}
