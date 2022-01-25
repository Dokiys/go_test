package win3cards

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"sync"
)

const baseRound = 3

type W3cSession struct {
	Players      []int // 玩家id
	Poker        *Win3Cards
	count        int          // 玩家人数
	Round        int          // 当前回合数
	ScoreMap     map[int]int  // 玩家分数 key playerId
	ReadyInfo    map[int]bool // 准备信息， 全部准备表示已经开局
	RoundSession *RoundSession

	Caller   func(ctx context.Context, id int, msg []byte) error // 向Player发送消息
	Receiver func(ctx context.Context, id int) ([]byte, error)   // 向Player发送消息
	l        sync.Mutex
}

func NewW3CSession(caller func(context.Context, int, []byte) error, receiver func(context.Context, int) ([]byte, error)) *W3cSession {
	return &W3cSession{Caller: caller, Receiver: receiver}
}

func (self *W3cSession) init(players []int) error {
	// Init
	length := len(players)
	self.Players = players
	self.count = length
	self.ReadyInfo = make(map[int]bool, length)
	self.ScoreMap = make(map[int]int, length)
	self.Round = 1
	self.Poker = NewPoker()

	// 开局
	for _, id := range self.Players {
		self.ScoreMap[id] = 0
		self.ReadyInfo[id] = true
	}
	self.RoundSession = NewRoundSession(players, self.Caller, self.Receiver)
	return nil
}

func (self *W3cSession) Run(ctx context.Context, players []int) error {
	// TODO[Dokiy] 2022/1/25: 处理panic
	//defer func() {
	//	if err := recover(); err != nil {
	//		logrus.Errorf("系统错误：", err)
	//	}
	//}()
	err := self.init(players)
	if err != nil {
		return errors.Wrapf(err, "初始化开局信息失败：")
	}

	// 发送局面信息
	self.BroadcastSession(ctx)
	for r := 1; r <= self.count*baseRound; r++ {
		self.Round = r
		// 等待玩家准备
		self.WaitReady(ctx)
		self.BroadcastMsg(ctx, "游戏开始！")

		// 开始
		winner, err := self.Play(ctx, r)
		if err != nil {
			return errors.Wrapf(err, "开局失败：")
		}

		// 结算
		self.Settle(winner)
		// 发送局面信息
		self.BroadcastSession(ctx)
	}

	// 结束信息
	self.BroadcastResult(ctx)
	return nil
}

func (self *W3cSession) WaitReady(ctx context.Context) {
	var wg sync.WaitGroup
	for _, playerId := range self.Players {
		if self.ReadyInfo[playerId] {
			continue
		}
		wg.Add(1)
		go func(id int) {
			for {
				data, err := self.Receiver(ctx, id)
				if err != nil {
					logrus.Errorf("等待玩家准备时，接收操作错误：", err.Error())
					continue
				}

				action, err := toAction(data)
				if err != nil {
					logrus.Errorf("解析玩家操作消息错误：%s", err.Error())
					continue
				}
				if !action.isW3CReady() {
					logrus.Errorf("玩家操作错误，需要进行准备操作！")
					continue
				}

				self.ReadyInfo[id] = true
				self.BroadcastSession(ctx)
				break
			}
			wg.Done()
		}(playerId)
	}
	wg.Wait()
}

func (self *W3cSession) Play(ctx context.Context, round int) (int, error) {
	l := len(self.Players)
	players := make([]int, l)
	for i, id := range self.Players {
		players[(round+i)%l] = id
	}
	// 发送位序
	self.BroadcastSeq(ctx, players)

	self.Poker.CutTheDeck()
	winner, err := self.RoundSession.Run(ctx, self.Poker, players)
	if err != nil {
		return 0, err
	}

	return winner, nil
}

func (self *W3cSession) InfoPlayer(ctx context.Context, id int, msg string) {
	go self.Caller(ctx, id, []byte(msg))
}

func (self *W3cSession) BroadcastMsg(ctx context.Context, msg string) {
	data := GenInfoMsg(msg)
	for _, id := range self.Players {
		go self.Caller(ctx, id, data)
	}
}

func (self *W3cSession) BroadcastSeq(ctx context.Context, players []int) {
	data := GenSeqMsg(players)
	for _, id := range players {
		go self.Caller(ctx, id, data)
	}
}

func (self *W3cSession) BroadcastSession(ctx context.Context) {
	data := GenW3cSessionMsg(self)
	for _, id := range self.Players {
		go self.Caller(ctx, id, data)
	}
}

func (self *W3cSession) BroadcastResult(ctx context.Context) {
	data := GenW3cResultMsg(self)
	for _, id := range self.Players {
		go self.Caller(ctx, id, data)
	}
}

func (self *W3cSession) Settle(winner int) {
	self.l.Lock()
	{
		var bet int
		for id, info := range self.RoundSession.PInfo {
			bet += info.Score
			self.ScoreMap[id] -= info.Score
		}
		self.ScoreMap[winner] += bet

		// 取消准备
		for id, _ := range self.ReadyInfo {
			self.ReadyInfo[id] = false
		}
	}
	self.l.Unlock()
}