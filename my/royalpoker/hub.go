package royalpoker

import (
	"context"
	"github.com/dokiy/royalpoker/common"
	"github.com/dokiy/royalpoker/win3cards"
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

	Register    chan *common.Player
	Unregister  chan *common.Player
	start       chan struct{}
	playSession PlaySession
}

func (self *Hub) isStarted() bool {
	_, r := <-self.start
	return !r
}

func NewHub(ownerId int) *Hub {
	id := common.RandNum(6)
	return &Hub{
		Id:          id,
		Owner:       ownerId,
		Players:     make(map[int]*common.Player),
		playSession: win3cards.NewW3CSession(),
		//Broadcast:  make(chan []byte),
		Register:   make(chan *common.Player),
		Unregister: make(chan *common.Player),
		start:      make(chan struct{}),
	}
}

func (self *Hub) addPlayer(player *common.Player) {
	self.Players[player.Id] = player
}

func (self *Hub) deletePlayer(player *common.Player) {
	if _, ok := self.Players[player.Id]; ok {
		delete(self.Players, player.Id)
	}
}

func (self *Hub) Run() error {
	ctx := context.Background()
	for {
		select {
		case player := <-self.Register:
			if self.isStarted() {
				// TODO[Dokiy] 2022/1/22: 返回消息给玩家
				continue
			}
			self.addPlayer(player)
		case player := <-self.Unregister:
			if self.isStarted() {
				// TODO[Dokiy] 2022/1/22: 返回消息给玩家
				continue
			}
			self.deletePlayer(player)
		case _, ok := <-self.start:
			if !ok {
				return errors.New("获取开始状态信息失败！")
			}

			err := self.Start(ctx)
			if err != nil {
				return errors.Wrapf(err, "开始游戏失败！")
			}

			// 发送结果给所有玩家
			// TODO[Dokiy] 2022/1/21:

		}
	}
}

func (self *Hub) Start(ctx context.Context) error {
	close(self.start)
	return errors.Wrapf(self.playSession.Run(ctx, self.Players), "开局失败：")
}
