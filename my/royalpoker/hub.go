package royalpoker

import (
	"context"
	"github.com/dokiy/royalpoker/common"
	"github.com/dokiy/royalpoker/win3cards"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type PlaySession interface {
	Run(ctx context.Context, players []int) error
}
type Hub struct {
	Id      int
	Owner   int
	Players map[int]*Player

	isStarted   bool
	playSession PlaySession
}

func NewHub(ownerId int) *Hub {
	id := common.RandNum(6)
	hub := &Hub{
		Id:      id,
		Owner:   ownerId,
		Players: make(map[int]*Player),
	}
	hub.playSession = win3cards.NewW3CSession(hub.CallPlayer, hub.ReceivePlayer)
	return hub
}

func (self *Hub) Register(conn *websocket.Conn) (*Player, error) {
	if self.isStarted {
		return nil, errors.New("游戏已开始！")
	}
	player := NewPlayer(conn)
	self.Players[player.Id] = player
	return player, nil
}

func (self *Hub) Unregister(player *Player) error {
	if self.isStarted {
		return errors.New("游戏已结束！")
	}
	_, ok := self.Players[player.Id]
	if !ok {
		return errors.New("未找到该玩家！")
	}
	delete(self.Players, player.Id)

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
		players[i] = player.Id
	}
	return errors.Wrapf(self.playSession.Run(ctx, players), "开局失败：")
}

func (self *Hub) CallPlayer(ctx context.Context, id int, msg []byte) {
	self.Players[id].Send(ctx, msg)
}

func (self *Hub) ReceivePlayer(ctx context.Context, id int) []byte {
	return self.Players[id].Receive(ctx)
}
