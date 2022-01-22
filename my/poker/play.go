package poker

import (
	"context"
	"github.com/pkg/errors"
	"sync"
)

const base = 1

type RoundSession struct {
	Players   map[int]*Player // key playerId
	Dealer    *Dealer
	handCards map[int]*HandCard // 玩家底牌 key playerId
	PInfo     map[int]*PlayInfo // 本局台面信息 key playerId
	PLog      []string          // 回合操作流水

	seq     []int // 本局playerId顺序
	current int   // 当前步骤玩家index
	l       sync.Mutex
}

type PlayInfo struct {
	Score    int  // 玩家投出积分
	IsViewed bool // 是否看牌
	IsOut    bool // 是否出局
}

func NewRoundSession(players map[int]*Player) *RoundSession {
	return &RoundSession{Players: players}
}

func (self *RoundSession) init(seq []int) error {
	if len(seq) <= 1 {
		return errors.New("人数不够开局！")
	}
	l := len(seq)
	self.handCards = make(map[int]*HandCard, l)
	self.PInfo = make(map[int]*PlayInfo, l)
	self.PLog = make([]string, l*4)
	self.current = 0

	self.seq = seq
	for _, id := range seq {
		handCard, err := self.Dealer.Deal()
		if err != nil {
			return errors.Wrapf(err, "发牌错误：")
		}
		self.handCards[id] = handCard
		self.PInfo[id] = &PlayInfo{}
	}

	return nil
}

func (self *RoundSession) Play(ctx context.Context, seq []int) (winner *Player, err error) {
	// 初始化开局
	if err := self.init(seq);err != nil {
		return nil, errors.Wrapf(err, "初始化开局错误：")
	}

	// 庄家下庄
	self.l.Lock()
	{
		i := ((self.current + len(self.seq)) - 1)%len(self.seq)
		self.getPInfoByIndex(i).Score -= len(self.seq) * base
	}
	self.l.Unlock()

	for {
		// TODO[Dokiy] 2022/1/21:  发送广播消息给所有玩家
		select {
		case a := self.currentPlayer().Action(ctx):
			// TODO[Dokiy] 2022/1/21: 等待当前玩家操作
		}

		if !self.nextPlayer() {
			break
		}
	}

	return self.currentPlayer(), nil
}

func (self *RoundSession) nextPlayer() (ok bool) {
	for i := 1; i < len(self.seq); i++ {
		index := (self.current + i) % len(self.seq)
		if info := self.getPInfoByIndex(index); !info.IsOut {
			self.current = index
			return true
		}
	}

	return false
}

// =========================================================

func (self *RoundSession) currentPlayer() *Player {
	return self.getPlayerByIndex(self.current)
}

func (self *RoundSession) getPInfoByIndex(i int) *PlayInfo {
	return self.PInfo[self.seq[i]]
}

func (self *RoundSession) getPInfoById(id int) *PlayInfo {
	return self.PInfo[id]
}

func (self *RoundSession) getPlayerByIndex(i int) *Player {
	return self.getPlayerById(self.seq[i])
}

func (self *RoundSession) getPlayerById(id int) *Player {
	return self.Players[id]
}
