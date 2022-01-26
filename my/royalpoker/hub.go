package main

import (
	"context"
	"fmt"
	"github.com/dokiy/royalpoker/common"
	"github.com/dokiy/royalpoker/win3cards"
	"github.com/pkg/errors"
	"time"
)

type PlaySession interface {
	Run(ctx context.Context, players []int) error
}
type Hub struct {
	Id      int
	Owner   int
	Players map[int]Player

	isStarted   bool
	playSession PlaySession
}

var hubMap map[int]*Hub

func init() {
	hubMap = make(map[int]*Hub)
}
func NewHub(owner int) *Hub {
	id := common.RandNum(6)
	hub := &Hub{
		Id:      id,
		Owner:   owner,
		Players: make(map[int]Player),
	}
	hub.playSession = win3cards.NewW3CSession(hub.callPlayer, hub.receivePlayer)

	hubMap[id] = hub
	go func() {
		select {
		case <-time.After(2 * time.Hour):
			delete(hubMap, id)
		}
	}()
	return hub
}

func (self *Hub) Register(player Player) error {
	p, ok := self.Players[player.GetId()]
	if self.isStarted && !ok {
		return errors.New("游戏已开始！")
	}
	// 如果之前对用户连接存在，则需要关闭原来对连接
	if ok {
		p.Close(context.TODO())
	}
	self.Players[player.GetId()] = player
	return nil
}

func (self *Hub) Unregister(player Player) error {
	if self.isStarted {
		return errors.New("游戏已结束！")
	}
	_, ok := self.Players[player.GetId()]
	if !ok {
		return errors.New("未找到该玩家！")
	}
	delete(self.Players, player.GetId())

	return nil
}

func (self *Hub) Run() error {
	ctx := context.Background()
	err := self.Start(ctx)
	if err != nil {
		return errors.Wrapf(err, "开始游戏失败！")
	}

	return nil
}

func (self *Hub) Start(ctx context.Context) error {
	self.isStarted = true
	var players = make([]int, len(self.Players))
	i := 0
	for _, player := range self.Players {
		players[i] = player.GetId()
		i++
	}
	err := self.playSession.Run(ctx, players)
	if err != nil {
		return errors.Wrapf(err, "开局失败：")
	}

	// 等待一会儿，让消息发送完成
	time.Sleep(10 * time.Second)
	for _, player := range self.Players {
		go player.Close(ctx)
	}
	return nil
}

func GetHub(id int) (*Hub, bool) {
	hub, ok := hubMap[id]
	return hub, ok
}

func (self *Hub) BroadcastHubSession(ctx context.Context, msg string) {
	data := GenHubSessionMsg(self, msg)
	for id, _ := range self.Players {
		go self.callPlayer(ctx, id, data)
	}
}

func (self *Hub) callPlayer(ctx context.Context, id int, msg []byte) error {
	player, ok := self.Players[id]
	if !ok {
		return errors.New(fmt.Sprintf("接收数据错误：未找到玩家[%d]", id))
	}
	player.Send(ctx, msg)
	return nil
}

func (self *Hub) receivePlayer(ctx context.Context, id int) ([]byte, error) {
	player, ok := self.Players[id]
	if !ok {
		return nil, errors.New(fmt.Sprintf("接收数据错误：未找到玩家[%d]", id))
	}
	return player.Receive(ctx), nil
}
