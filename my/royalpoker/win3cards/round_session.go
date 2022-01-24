package win3cards

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"sync"
)

const base = 1

type RoundSession struct {
	Players   []int             // 玩家Id
	handCards map[int]HandCard  // 玩家底牌 key playerId
	MaxBet    int               // 当前轮注码(开牌值计算)
	PInfo     map[int]*PlayInfo // 本局玩家信息 key playerId
	PLog      []string          // 回合操作流水

	ViewLog map[int][]int // 看牌记录

	Caller   func(ctx context.Context, id int, msg []byte) // 向Player发送消息
	Receiver func(ctx context.Context, id int) []byte      // 从Player接收消息
	seq      []int                                         // 本局playerId顺序
	current  int                                           // 当前步骤玩家index
	l        sync.Mutex
}

type PlayInfo struct {
	Score    int  // 玩家投出积分
	IsViewed bool // 是否看牌
	IsOut    bool // 是否出局
}

func NewRoundSession(players []int,
	caller func(context.Context, int, []byte),
	receiver func(context.Context, int) []byte) *RoundSession {

	return &RoundSession{
		Players:  players,
		Caller:   caller,
		Receiver: receiver,
	}
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
	self.MaxBet = base * 2

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

func (self *RoundSession) Play(ctx context.Context, poker *Win3Cards, seq []int) (winner int, err error) {
	// 初始化开局
	if err := self.init(poker, seq); err != nil {
		return 0, errors.Wrapf(err, "初始化开局错误：")
	}

	// 庄家下庄
	self.blind()

	// 开始叫牌
	var showWinner bool
	for {
		// 发送台面消息给所有玩家
		self.BroadcastSession(ctx)

		// 等待当前玩家操作
		data := self.waitAction(ctx)

		// 处理玩家操作
		action, err := toAction(data)
		if err != nil {
			logrus.Errorf("解析玩家操作消息错误：%s", err.Error())
			continue
		}
		err = action.do(ctx, self)
		if err != nil {
			go self.Caller(ctx, self.currentPlayer(), GenInfoMsg(err.Error()))
			continue
		}

		if action.isContinued() {
			continue
		}

		// 如果没有下一个玩家，结束游戏
		if !self.next() {

			// 如果最后开牌结束，则需要看赢家的牌
			if action.isShow() {
				showWinner = true
			}
			break
		}
	}

	// 摊牌
	self.showdown(ctx, showWinner)

	return self.currentPlayer(), nil
}

func (self *RoundSession) BroadcastSession(ctx context.Context) {
	msg := GenRoundSessionMsg(self)
	for _, id := range self.Players {
		go self.Caller(ctx, id, msg)
	}
}

// =========================================================

func (self *RoundSession) blind() {
	self.l.Lock()
	defer self.l.Unlock()

	l := len(self.seq)
	i := ((self.current + l) - 1) % l
	self.getPInfoByIndex(i).Score += l * base
}

func (self *RoundSession) waitAction(ctx context.Context) []byte {
	self.Caller(ctx, self.currentPlayer(), GenInfoMsg("It's your turn！"))
	return self.Receiver(ctx, self.currentPlayer())
}

func (self *RoundSession) next() (ok bool) {
	for i := 1; i < len(self.seq); i++ {
		index := (self.current + i) % len(self.seq)
		if info := self.getPInfoByIndex(index); !info.IsOut {
			self.current = index
			return true
		}
	}

	return false
}

func (self *RoundSession) showdown(ctx context.Context, showWinner bool) {
	for id, playerIds := range self.ViewLog {
		handCards := make(map[int]HandCard, len(playerIds))
		for _, playerId := range playerIds {
			handCards[playerId] = self.handCards[playerId]
		}

		// 看自己的牌
		handCards[id] = self.handCards[id]
		// 看赢家的牌
		if showWinner {
			handCards[self.current] = self.handCards[self.current]
		}
		self.Caller(ctx, id, GenViewLogMsg(handCards))
	}
}

// =========================================================

func (self *RoundSession) currentPlayer() int {
	return self.getPlayerByIndex(self.current)
}

func (self *RoundSession) currentPInfo() *PlayInfo {
	return self.getPInfoByIndex(self.current)
}

func (self *RoundSession) getPInfoByIndex(i int) *PlayInfo {
	return self.PInfo[self.seq[i]]
}

func (self *RoundSession) getPInfoById(id int) *PlayInfo {
	return self.PInfo[id]
}

func (self *RoundSession) getPlayerByIndex(i int) int {
	return self.seq[i]
}
