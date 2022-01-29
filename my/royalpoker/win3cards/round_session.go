package win3cards

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"sync"
)

const base = 1

type RoundSession struct {
	Players   []int             // 玩家Id，顺序未玩家开始顺序
	handCards map[int]HandCard  // 玩家底牌 key index
	MaxBet    int               // 当前轮注码(开牌值计算)
	PInfo     map[int]*PlayInfo // 本局玩家信息,因为在w3c_session中需要结算，所以key使用playerId
	PLog      []string          // 回合操作流水

	ViewLog map[int][]int // 看牌记录 key index

	Caller   func(ctx context.Context, id int, msg []byte) error // 向Player发送消息
	Receiver func(ctx context.Context, id int) ([]byte, error)   // 从Player接收消息
	current  int                                                 // 当前步骤玩家index
	l        sync.Mutex
}

type PlayInfo struct {
	Score    int  `json:"score"`     // 玩家投出积分
	IsViewed bool `json:"is_viewed"` // 是否看牌
	IsOut    bool `json:"is_out"`    // 是否出局
}

func NewRoundSession(players []int,
	caller func(context.Context, int, []byte) error,
	receiver func(context.Context, int) ([]byte, error)) *RoundSession {

	return &RoundSession{
		Caller:   caller,
		Receiver: receiver,
	}
}

func (self *RoundSession) init(poker *Win3Cards, players []int) error {
	l := len(players)
	self.handCards = make(map[int]HandCard, l)
	self.PInfo = make(map[int]*PlayInfo, l)
	self.PLog = make([]string, 0)
	self.ViewLog = make(map[int][]int, l)
	self.current = 0
	self.MaxBet = base * 2

	for i, id := range players {
		handCard, err := poker.Deal()
		if err != nil {
			return errors.Wrapf(err, "发牌错误：")
		}
		self.handCards[i] = handCard
		self.PInfo[id] = &PlayInfo{}
	}

	return nil
}

func (self *RoundSession) Run(ctx context.Context, poker *Win3Cards, players []int) (winner int, err error) {
	// 初始化开局
	if err := self.init(poker, players); err != nil {
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
		data, err := self.waitAction(ctx)
		if err != nil {
			if errors.As(err, &context.Canceled) {
				return 0, errors.Wrapf(err, "游戏被关闭：")
			}
			logrus.Errorf("接收玩家操作消息错误：%s", err.Error())
			continue
		}
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
			showWinner = action.isShow()
			break
		}
	}

	// 发送台面消息给所有玩家
	self.BroadcastSession(ctx)
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

func (self *RoundSession) BroadcastInfo(ctx context.Context, msg string) {
	data := GenInfoMsg(msg)
	for _, id := range self.Players {
		go self.Caller(ctx, id, data)
	}
}

// =========================================================

func (self *RoundSession) blind() {
	self.l.Lock()
	defer self.l.Unlock()

	l := len(self.Players)
	//i := ((self.current + l) - 1) % l
	self.getPInfoByIndex(self.current).Score += l * base
}

func (self *RoundSession) waitAction(ctx context.Context) ([]byte, error) {
	go self.BroadcastInfo(ctx, fmt.Sprintf("轮到[%d]号玩家操作", self.current+1))
	data, err := self.Receiver(ctx, self.currentPlayer())
	if err != nil {
		return nil, errors.Wrapf(err, "接收操作失败！")
	}
	return data, nil
}

func (self *RoundSession) next() (ok bool) {
	for i := 1; i < len(self.Players); i++ {
		index := (self.current + i) % len(self.Players)
		if info := self.getPInfoByIndex(index); !info.IsOut {
			self.current = index
			return true
		}
	}

	return false
}

func (self *RoundSession) showdown(ctx context.Context, showWinner bool) {
	for index, playerIndexs := range self.ViewLog {
		handCards := make(map[int]HandCard, len(playerIndexs)+2)
		for _, playerIndex := range playerIndexs {
			handCards[playerIndex] = self.handCards[playerIndex]
		}

		// 看自己的牌
		handCards[index] = self.handCards[index]
		// 看赢家的牌
		if showWinner {
			handCards[self.current] = self.handCards[self.current]
		}
		go self.Caller(ctx, self.Players[index], GenViewLogMsg(handCards))
	}
}

// =========================================================

func (self *RoundSession) currentPlayer() int {
	return self.Players[self.current]
}

func (self *RoundSession) currentPInfo() *PlayInfo {
	return self.getPInfoByIndex(self.current)
}

func (self *RoundSession) getPInfoByIndex(i int) *PlayInfo {
	return self.PInfo[self.Players[i]]
}
