package win3cards

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type ActionMsg string
type Action struct {
	ActionMsg ActionMsg
	Bet       int
	ShowId    int
}

const (
	ACTION_IN   = "ACTION_IN"
	ACTION_OUT  = "ACTION_OUT"
	ACTION_VIEW = "ACTION_VIEW"
	ACTION_SHOW = "ACTION_SHOW"
)

func (self Action) do(rs *RoundSession) error {
	switch self.ActionMsg {
	case ACTION_IN:
		rs.l.Lock()
		defer rs.l.Unlock()

		pInfo := rs.CurrentPInfo()
		if pInfo.IsOut {
			return errors.New("已经出局的玩家不能进行下注！")
		}

		// 下注
		{
			var base = 1
			if pInfo.IsViewed {
				base = 2
			}
			if base*self.Bet < rs.MaxBet {
				return errors.New("下注必须大于当前最大注码！")
			}

			pInfo.Score -= self.Bet
			rs.MaxBet = self.Bet * base
		}

	case ACTION_OUT:
		rs.CurrentPInfo().IsOut = true

	case ACTION_VIEW:
		rs.CurrentPInfo().IsViewed = true
		// TODO[Dokiy] 2022/1/24:  发送消息给看牌玩家

	case ACTION_SHOW:
		rs.l.Lock()
		defer rs.l.Unlock()

		pInfo1, pInfo2 := rs.CurrentPInfo(), rs.getPInfoById(self.ShowId)
		if pInfo1.IsOut || pInfo2.IsOut {
			return errors.New("已经出局的玩家不能(被)开牌！")
		}

		// 下注
		{
			var base = 1
			if pInfo1.IsViewed {
				base = 2
			}
			if base*self.Bet < rs.MaxBet {
				return errors.New("下注必须大于当前最大注码！")
			}

			pInfo1.Score -= self.Bet
		}

		// 开牌, 并记录给开牌输家看牌
		{
			h1, h2 := rs.handCards[rs.current], rs.handCards[self.ShowId]
			if Compare(h1, h2) {
				pInfo1.IsOut = true
				rs.ViewLog[rs.current] = append(rs.ViewLog[rs.current], self.ShowId)
			} else {
				pInfo2.IsOut = true
				rs.ViewLog[self.ShowId] = append(rs.ViewLog[self.ShowId], rs.current)
			}
		}
	}

	return nil
}

func (self Action) isContinued() bool {
	return self.ActionMsg == ACTION_VIEW
}

func (self Action) isShow() bool {
	return self.ActionMsg == ACTION_SHOW
}

// ==========================================

func toAction(data []byte) (Action, error) {
	action := &Action{}
	return *action, errors.Wrapf(json.Unmarshal(data, action), "解析操作消息错误：")
}
