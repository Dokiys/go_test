package win3cards

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

const (
	MSGTYPE_INFO          = "INFO"
	MSGTYPE_READY_INFO    = "READY_INFO"
	MSGTYPE_ROUND_SESSION = "ROUND_SESSION"
	MSGTYPE_ACTION_VIEW   = "ACTION_VIEW"
	MSGTYPE_VIEW_LOG      = "VIEW_LOG"
	MSGTYPE_W3C_SESSION   = "W3C_SESSION"
	MSGTYPE_W3C_RESULT    = "W3C_RESULT"
)

type MsgType string
type InfoMsg struct {
	Type MsgType
	Msg  string
}

type W3cSessionMsg struct {
	Type MsgType
	Data W3cSessionMsgData
}

type W3cSessionMsgData struct {
	Seq       []int        // 玩具顺序
	ScoreMap  map[int]int  // 玩家分数
	Round     int          // 当前回合数
	ReadyInfo map[int]bool // 准备信息
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

// =========================================================

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

//func GenW3cReadyInfo(ws *W3cSession) []byte {
//	wsMsg := W3cReadyInfo{
//		Type: MSGTYPE_READY_INFO,
//		Data: W3cReadyInfoData{
//			ReadyInfo: ws.ReadyInfo,
//		},
//	}
//	bytes, err := json.Marshal(wsMsg)
//	if err != nil {
//		logrus.Errorf("序列化RoundSessionMsg失败: ", err.Error())
//	}
//	return bytes
//}

func GenW3cSessionMsg(ws *W3cSession) []byte {
	wsMsg := W3cSessionMsg{
		Type: MSGTYPE_W3C_SESSION,
		Data: W3cSessionMsgData{
			Seq:      ws.Seq,
			ScoreMap: ws.ScoreMap,
			Round:    ws.Round,
		},
	}
	bytes, err := json.Marshal(wsMsg)
	if err != nil {
		logrus.Errorf("序列化RoundSessionMsg失败: ", err.Error())
	}
	return bytes
}

func GenW3cResultMsg(ws *W3cSession) []byte {
	wsMsg := W3cSessionMsg{
		Type: MSGTYPE_W3C_RESULT,
		Data: W3cSessionMsgData{
			Seq:      ws.Seq,
			ScoreMap: ws.ScoreMap,
			Round:    ws.Round,
		},
	}
	bytes, err := json.Marshal(wsMsg)
	if err != nil {
		logrus.Errorf("序列化RoundSessionMsg失败: ", err.Error())
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
