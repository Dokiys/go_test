package win3cards

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

type ActionType string
type Action struct {
	ActionType ActionType `json:"action_type"`
	Bet        int        `json:"bet"`
	ShowIndex  int        `json:"show_index"` // 被开牌玩家的顺序 1开始计算
}

const (
	W3C_ACTION_READY = "W3C_ACTION_READY"

	ACTION_IN   = "ACTION_IN"
	ACTION_OUT  = "ACTION_OUT"
	ACTION_VIEW = "ACTION_VIEW"
	ACTION_SHOW = "ACTION_SHOW"
)

func (self Action) do(ctx context.Context, rs *RoundSession) error {
	switch self.ActionType {
	case ACTION_IN:
		rs.l.Lock()
		defer rs.l.Unlock()

		pInfo := rs.currentPInfo()
		if pInfo.IsOut {
			return errors.New("已经出局的玩家不能进行下注！")
		}

		// 下注
		{
			var base = 2
			if pInfo.IsViewed {
				base = 1
			}
			if base*self.Bet < rs.MaxBet {
				return errors.New("下注必须大于当前最大注码！")
			}

			pInfo.Score += self.Bet
			rs.MaxBet = self.Bet * base
		}

	case ACTION_OUT:
		rs.currentPInfo().IsOut = true

	case ACTION_VIEW:
		rs.currentPInfo().IsViewed = true
		handCard := rs.handCards[rs.current]
		rs.Caller(ctx, rs.currentPlayer(), GenActionViewMsg(handCard))

	case ACTION_SHOW:
		rs.l.Lock()
		defer rs.l.Unlock()

		if rs.current == self.ShowIndex {
			return errors.New("不能开自己的牌！")
		}

		pInfo1, pInfo2 := rs.currentPInfo(), rs.getPInfoByIndex(self.ShowIndex)
		if pInfo2 == nil {
			return errors.Errorf("错误：未找到该玩家！")
		}
		if pInfo1.IsOut || pInfo2.IsOut {
			return errors.New("已经出局的玩家不能(被)开牌！")
		}

		// 下注
		{
			var base = 2
			if pInfo1.IsViewed {
				base = 1
			}
			if base*self.Bet < rs.MaxBet {
				return errors.New("下注必须大于当前最大注码！")
			}

			pInfo1.Score += self.Bet
		}

		// 开牌, 并记录给开牌输家看牌
		{
			h1, h2 := rs.handCards[rs.current], rs.handCards[self.ShowIndex]
			if h2.v == "" {
				return errors.Errorf("错误：未找到该玩家底牌！")
			}

			// 开牌者可以看到被开者到牌
			rs.ViewLog[rs.current] = append(rs.ViewLog[rs.current], self.ShowIndex)
			if Compare(h1, h2) {
				// 被开牌者如果输了也可以看到开牌者的牌
				pInfo2.IsOut = true
				rs.ViewLog[self.ShowIndex] = append(rs.ViewLog[self.ShowIndex], rs.current)
			} else {
				pInfo1.IsOut = true
				// 如果当前玩家输了，设置下一个玩家
				rs.next()
			}
		}
	}

	rs.PLog = append(rs.PLog, self.genPLog(rs))
	return nil
}

// ==========================================

func (self *Action) genPLog(rs *RoundSession) (plog string) {
	plog = fmt.Sprintf("玩家[%d]", rs.GetPlayerName(rs.Players[rs.current]))
	switch self.ActionType {
	case ACTION_IN:
		plog = fmt.Sprintf("%s【跟注】：【%d】", plog, self.Bet)
	case ACTION_OUT:
		plog = fmt.Sprintf("%s【弃牌】", plog)
	case ACTION_VIEW:
		plog = fmt.Sprintf("%s进行了【看牌】", plog)
	case ACTION_SHOW:
		var outId int
		if rs.currentPInfo().IsOut {
			outId = rs.currentPlayer()
		} else {
			outId = self.ShowIndex
		}
		plog = fmt.Sprintf("%s【开牌】[%d]号玩家：[%d]号玩家出局", plog, self.ShowIndex+1, outId+1)
	}

	return plog
}

func toAction(data []byte) (Action, error) {
	action := &Action{}
	return *action, errors.Wrapf(json.Unmarshal(data, action), "解析操作消息错误：")
}

func (self Action) isContinued() bool {
	return self.ActionType == ACTION_VIEW
}

func (self Action) isShow() bool {
	return self.ActionType == ACTION_SHOW
}

func (self Action) isW3CReady() bool {
	return self.ActionType == W3C_ACTION_READY
}
