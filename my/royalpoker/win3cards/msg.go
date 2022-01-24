package win3cards

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

const (
	MSGTYPE_INFO          = "INFO"
	MSGTYPE_ROUND_SESSION = "ROUND_SESSION"
	MSGTYPE_ACTION_VIEW   = "ACTION_VIEW"
	MSGTYPE_VIEW_LOG      = "VIEW_LOG"
)

type MsgType string
type InfoMsg struct {
	Type MsgType
	Msg  string
}

type RoundSessionMsg struct {
	Type MsgType
	Data RoundSessionData
}

type RoundSessionData struct {
	PInfo  map[int]*PlayInfo // 本局玩家信息 key playerId
	PLog   []string          // 回合操作流水
	MaxBet int               // 当前轮注码(开牌值计算)
}

type ActionViewMsg struct {
	Type MsgType
	Data ActionViewData
}

type ActionViewData struct {
	HandCard HandCard
}

type ViewLogMsg struct {
	Type MsgType
	Data ViewLogData
}

type ViewLogData struct {
	HandCards map[int]HandCard
}

func GenInfoMsg(msg string) []byte {
	infoMsg := InfoMsg{
		Type: MSGTYPE_INFO,
		Msg:  msg,
	}
	bytes, err := json.Marshal(infoMsg)
	if err != nil {
		logrus.Errorf("序列化InfoMsg失败: ", err.Error())
	}

	return bytes
}

func GenRoundSessionMsg(rs *RoundSession) []byte {
	rsMsg := RoundSessionMsg{
		Type: MSGTYPE_ROUND_SESSION,
		Data: RoundSessionData{
			PInfo:  rs.PInfo,
			PLog:   rs.PLog,
			MaxBet: rs.MaxBet,
		},
	}
	bytes, err := json.Marshal(rsMsg)
	if err != nil {
		logrus.Errorf("序列化RoundSessionMsg失败: ", err.Error())
	}
	return bytes
}

func GenViewLogMsg(handCards map[int]HandCard) []byte {
	vlMsg := ViewLogMsg{
		Type: MSGTYPE_VIEW_LOG,
		Data: ViewLogData{
			HandCards: handCards,
		},
	}
	bytes, err := json.Marshal(vlMsg)
	if err != nil {
		logrus.Errorf("序列化ViewLogMsg失败: ", err.Error())
	}
	return bytes
}
func GenActionViewMsg(card HandCard) []byte {
	act := ActionViewMsg{
		Type: MSGTYPE_ACTION_VIEW,
		Data: ActionViewData{
			HandCard: card,
		},
	}
	bytes, err := json.Marshal(act)
	if err != nil {
		logrus.Errorf("序列化ActionViewMsg失败: ", err.Error())
	}

	return bytes
}
