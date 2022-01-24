package win3cards

import (
	"context"
	"github.com/pkg/errors"
	"sync"
)

const baseRound = 3

type w3cSession struct {
	Players      []int // 玩家id
	Poker        *Win3Cards
	count        int          // 玩家人数
	Round        int          // 当前回合数
	ScoreMap     map[int]int  // 玩家分数 key playerId
	Seq          []int        // 玩家顺序
	ReadyInfo    map[int]bool // 准备信息， 全部准备表示已经开局
	RoundSession *RoundSession

	Caller   func(ctx context.Context, id int, msg []byte) // 向Player发送消息
	Receiver func(ctx context.Context, id int) []byte      // 向Player发送消息
	l        sync.Mutex
}

func NewW3CSession(caller func(context.Context, int, []byte), receiver func(context.Context, int) []byte) *w3cSession {
	return &w3cSession{Caller: caller, Receiver: receiver}
}

func (self *w3cSession) init(players []int) error {
	// Init
	length := len(players)
	self.Players = players
	self.count = length
	self.ReadyInfo = make(map[int]bool, length)
	self.Seq = make([]int, length)
	self.Poker = NewPoker()

	// 开局
	i := 0
	for id, _ := range self.Players {
		self.Seq[i] = id
		self.ScoreMap[id] = 0
		self.ReadyInfo[id] = true

		i++
	}
	self.RoundSession = NewRoundSession(players, self.Caller, self.Receiver)
	return nil
}

func (self *w3cSession) Run(ctx context.Context, players []int) error {
	err := self.init(players)
	if err != nil {
		return errors.Wrapf(err, "初始化开局信息失败：")
	}
	for r := 0; r < self.count*baseRound; r++ {
		self.Round = r
		// TODO[Dokiy] 2022/1/23: 等待玩家准备开始新游戏

		l := len(self.Seq)
		seq := make([]int, l)
		for i, id := range self.Seq {
			seq[(r+i)%l] = id
		}
		self.Poker.CutTheDeck()
		winner, err := self.RoundSession.Play(ctx, self.Poker, seq)
		if err != nil {
			return err
		}

		self.settle(winner)
	}

	// TODO[Dokiy] 2022/1/23: 更新结果，并发送准备消息
	return nil
}

func (self *w3cSession) settle(winner int) {
	self.l.Lock()
	{
		var bet int
		for id, info := range self.RoundSession.PInfo {
			bet += info.Score
			self.ScoreMap[id] -= info.Score
		}
		self.ScoreMap[winner] += bet
	}
	self.l.Unlock()
}
