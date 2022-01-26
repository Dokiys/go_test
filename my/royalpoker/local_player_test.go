package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	file, err := os.OpenFile("/Users/admin/dokiy/go_test/my/royalpoker/log", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(file)

	ctx := context.Background()
	hub := NewHub(1)

	// 注册
	player := NewLocalPlayer("zhangsan")
	player2 := NewLocalPlayer("lisi")
	//a := player.Receive(ctx)
	//t.Log(a)

	err = hub.Register(player)
	if err != nil {
		panic(err)
	}

	err = hub.Register(player2)
	if err != nil {
		panic(err)
	}

	err = hub.Start(ctx)
	if err != nil {
		panic(err)
	}
}

//	ACTION_IN   = "ACTION_IN"
//	ACTION_OUT  = "ACTION_OUT"
//	ACTION_VIEW = "ACTION_VIEW"
//	ACTION_SHOW = "ACTION_SHOW"

// 准备： {"action_type":"W3C_ACTION_READY","bet": 0, "show_index":0}
// 跟： {"action_type":"ACTION_IN","bet": 1, "show_index":0}
// 看牌： {"action_type":"ACTION_VIEW","bet": 0, "show_index":0}
// 开： {"action_type":"ACTION_SHOW","bet": 2, "show_index":0}