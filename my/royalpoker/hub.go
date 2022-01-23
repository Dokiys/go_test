package royalpoker

import (
	"context"
	"github.com/dokiy/royalpoker/common"
	"github.com/dokiy/royalpoker/win3cards"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type PlaySession interface {
	Run(ctx context.Context, players map[int]*common.Player) error
}
type Hub struct {
	Id      int
	Owner   int
	Players map[int]*common.Player
	//Broadcast chan []byte

	isStarted   bool
	playSession PlaySession
}

func NewHub(ownerId int) *Hub {
	id := common.RandNum(6)
	return &Hub{
		Id:          id,
		Owner:       ownerId,
		Players:     make(map[int]*common.Player),
		playSession: win3cards.NewW3CSession(),
		//Broadcast:  make(chan []byte),
	}
}

func (self *Hub) Register(conn *websocket.Conn) (*common.Player, error) {
	if self.isStarted {
		return nil, errors.New("游戏已开始！")
	}
	player := common.NewPlayer(conn)
	self.Players[player.Id] = player
	return player, nil
}

func (self *Hub) Unregister(player *common.Player) error {
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
	return errors.Wrapf(self.playSession.Run(ctx, self.Players), "开局失败：")
}
