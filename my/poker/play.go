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

func (self *RoundSession) Play(ctx context.Context, seq []int, round int) (winner *Player, err error) {
	l := len(seq)
	button := round%l - 1

	self.seq = make([]int, l)
	for i := 0; i < l; i++ {
		id := seq[(i+button)%l]

		self.seq[i] = id
		handCard, err := self.Dealer.Deal()
		if err != nil {
			return nil, errors.Wrapf(err, "发牌错误：")
		}
		self.handCards[id] = handCard
	}

	{
		buttonId := seq[button]
		self.l.Lock()
		self.PInfo[buttonId].Score -= l * base
		self.current = (button + 1) % l
		self.l.Unlock()
	}

	for {
		// TODO[Dokiy] 2022/1/21:  发送广播消息给所有玩家
		select {
		case a := self.currentPlayer().WaitDo(ctx):
			// TODO[Dokiy] 2022/1/21: 等待当前玩家操作
		}

		if !self.next() {
			break
		}
	}

	return self.currentPlayer(), nil
}

func (self *RoundSession) next() (ok bool) {
	for i := 1; i < len(self.seq); i++ {
		index := (self.current + i) % len(self.seq)
		if info:= self.getPInfoByIndex(index); !info.IsOut {
			self.current = index
			return true
		}
	}

	return false
}

func (self *RoundSession) currentPlayer() *Player {
	return self.getPlayerByIndex(self.current)
}

func (self *RoundSession) getPInfoByIndex(i int) *PlayInfo {
	return self.PInfo[self.seq[i]]
}

func (self *RoundSession) getPlayerByIndex(i int) *Player {
	return self.getPlayerById(self.seq[i])
}

func (self *RoundSession) getPlayerById(id int) *Player {
	return self.Players[id]
}

