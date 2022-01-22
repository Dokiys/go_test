package win3cards

import (
	"context"
	"github.com/dokiy/royalpoker/common"
	"github.com/pkg/errors"
	"sync"
)

const base = 1

type RoundSession struct {
	Players   map[int]*common.Player // key playerId
	handCards map[int]HandCard        // 玩家底牌 key playerId
	PInfo     map[int]*PlayInfo       // 本局台面信息 key playerId
	PLog      []string                // 回合操作流水

	seq     []int // 本局playerId顺序
	current int   // 当前步骤玩家index
	l       sync.Mutex
}

type PlayInfo struct {
	Score    int  // 玩家投出积分
	IsViewed bool // 是否看牌
	IsOut    bool // 是否出局
}

func NewRoundSession(players map[int]*common.Player) *RoundSession {
	return &RoundSession{Players: players}
}

func (self *RoundSession) init(poker *Win3Cards, seq []int) error {
	if len(seq) <= 1 {
		return errors.New("人数不够开局！")
	}
	l := len(seq)
	self.handCards = make(map[int]HandCard, l)
	self.PInfo = make(map[int]*PlayInfo, l)
	self.PLog = make([]string, l*4)
	self.current = 0

	self.seq = seq
	for _, id := range seq {
		handCard, err := poker.Deal()
		if err != nil {
			return errors.Wrapf(err, "发牌错误：")
		}
		self.handCards[id] = handCard
		self.PInfo[id] = &PlayInfo{}
	}

	return nil
}

func (self *RoundSession) Play(ctx context.Context, poker *Win3Cards, seq []int) (winner *common.Player, err error) {
	// 初始化开局
	if err := self.init(poker, seq); err != nil {
		return nil, errors.Wrapf(err, "初始化开局错误：")
	}

	// 庄家下庄
	self.l.Lock()
	{
		i := ((self.current + len(self.seq)) - 1) % len(self.seq)
		self.getPInfoByIndex(i).Score -= len(self.seq) * base
	}
	self.l.Unlock()

	for {
		// TODO[Dokiy] 2022/1/21:  发送广播消息给所有玩家
		// ...

		// TODO[Dokiy] 2022/1/21: 等待当前玩家操作
		action := self.currentPlayer().WaitAction(ctx)
		print(action)
		// TODO[Dokiy] 2022/1/22: 处理当前玩家的操作
		// 1 跟，2 加注，3 弃牌，4，看牌 5，开其他玩家的牌，
		{

		}

		if isContinued(action) {
			continue
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

func (self *RoundSession) currentPlayer() *common.Player {
	return self.getPlayerByIndex(self.current)
}

func (self *RoundSession) getPInfoByIndex(i int) *PlayInfo {
	return self.PInfo[self.seq[i]]
}

func (self *RoundSession) getPInfoById(id int) *PlayInfo {
	return self.PInfo[id]
}

func (self *RoundSession) getPlayerByIndex(i int) *common.Player {
	return self.getPlayerById(self.seq[i])
}

func (self *RoundSession) getPlayerById(id int) *common.Player {
	return self.Players[id]
}
