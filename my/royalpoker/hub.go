package royalpoker

import (
	"context"
	"fmt"
	"github.com/dokiy/royalpoker/common"
	"github.com/dokiy/royalpoker/win3cards"
	"github.com/pkg/errors"
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

func NewHub(ownerId int) *Hub {
	id := common.RandNum(6)
	hub := &Hub{
		Id:      id,
		Owner:   ownerId,
		Players: make(map[int]Player),
	}
	hub.playSession = win3cards.NewW3CSession(hub.CallPlayer, hub.ReceivePlayer)
	return hub
}

func (self *Hub) Register(player Player)  error {
	if self.isStarted {
		return errors.New("游戏已开始！")
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
	for i, player := range self.Players {
		players[i] = player.GetId()
	}
	err := self.playSession.Run(ctx, players)
	if err != nil {
		return errors.Wrapf(err, "开局失败：")
	}

	for _, player := range self.Players {
		go player.Close(ctx)
	}
	return nil
}

func (self *Hub) CallPlayer(ctx context.Context, id int, msg []byte) error {
	player, ok := self.Players[id]
	if !ok {
		return errors.New(fmt.Sprintf("接收数据错误：未找到玩家[%d]", id))
	}
	player.Send(ctx, msg)
	return nil
}

func (self *Hub) ReceivePlayer(ctx context.Context, id int) ([]byte,error) {
	player, ok := self.Players[id]
	if !ok {
		return nil, errors.New(fmt.Sprintf("接收数据错误：未找到玩家[%d]", id))
	}
	return player.Receive(ctx), nil
}
