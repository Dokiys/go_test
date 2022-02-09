package main

import (
	"context"
	"fmt"
	"github.com/dokiy/royalpoker/common"
	"github.com/dokiy/royalpoker/win3cards"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type PlaySession interface {
	Run(ctx context.Context, players []int) error
	BroadcastSession(ctx context.Context)
	InfoPlayerSession(ctx context.Context, id int)
}
type Hub struct {
	Id      int
	Owner   int
	Players map[int]Player
	playSession PlaySession
	IsStarted   bool

	ctx         context.Context
	close       chan struct{}
	l           sync.Mutex
}

var hubMap map[int]*Hub

func init() {
	hubMap = make(map[int]*Hub)
}
func NewHub(owner int) *Hub {
	id := common.RandNum(9999)
	hub := &Hub{
		Id:      id,
		Owner:   owner,
		Players: make(map[int]Player),
		close: make(chan struct{}),
	}
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		select {
		case <-hub.close:
			cancel()
		}
	}()
	hub.ctx = ctx
	hub.playSession = win3cards.NewW3CSession(hub.callPlayer, hub.receivePlayer, hub.getPlayerName)

	hubMap[id] = hub
	go func() {
		select {
		case <-time.After(2 * time.Hour):
			delete(hubMap, id)
		}
	}()
	return hub
}

var hubLock sync.Mutex
func (self *Hub) Register(player Player) error {
	hubLock.Lock()
	defer hubLock.Unlock()
	if self.IsStarted {
		return errors.New("游戏已开始！")
	}
	// 如果之前对用户连接存在，则需要关闭原来对连接
	//p, ok := self.Players[player.GetId()]
	//if ok {
	//	p.Close(context.TODO())
	//}
	self.Players[player.GetId()] = player
	if self.IsStarted {
		self.playSession.InfoPlayerSession(self.ctx, player.GetId())
	}
	return nil
}

func (self *Hub) Unregister(playerId int) error {
	hubLock.Lock()
	defer hubLock.Unlock()
	if self.IsStarted {
		return errors.New("游戏已开始！")
	}
	_, ok := self.Players[playerId]
	if !ok {
		return nil
	}
	delete(self.Players, playerId)

	return nil
}

func (self *Hub) Start() error {
	hubLock.Lock()
	if self.IsStarted == true {
		hubLock.Unlock()
		return nil
	}
	self.IsStarted = true
	hubLock.Unlock()
	var players = make([]int, len(self.Players))
	i := 0
	for _, player := range self.Players {
		players[i] = player.GetId()
		i++
	}
	err := self.playSession.Run(self.ctx, players)
	if errors.As(err, &context.Canceled) {
		logrus.Errorf("游戏被取消", err.Error())
	}
	if err != nil {
		self.IsStarted = false
		return errors.Wrapf(err, "开局失败")
	}

	self.IsStarted = false
	for id, _ := range self.Players {
		delete(userHubMap, id)
	}

	return nil
}

func (self *Hub) Close(force bool) {
	if self.IsStarted && !force {
		return
	}
	for _, player := range self.Players {
		go player.Close(self.ctx)
	}
	delete(hubMap, self.Id)
	self.close <- struct{}{}
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
func (self *Hub) InfoPlayerRelinkSession(ctx context.Context, id int) {
	if !self.IsStarted {
		return
	}
	self.playSession.InfoPlayerSession(ctx, id)
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
	return player.Receive(ctx)
}

func (self *Hub) getPlayerName(id int) string {
	player, ok := self.Players[id]
	if !ok {
		return ""
	}
	return player.GetName()
}
