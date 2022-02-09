package win3cards

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

const (
	MSGTYPE_INFO           = "INFO"
	MSGTYPE_ROUND_SESSION  = "ROUND_SESSION"
	MSGTYPE_ACTION_VIEW    = "ACTION_VIEW"
	MSGTYPE_VIEW_LOG       = "VIEW_LOG"
	MSGTYPE_SEQ            = "SEQ"
	MSGTYPE_W3C_SESSION    = "W3C_SESSION"
	MSGTYPE_W3C_RESULT     = "W3C_RESULT"
	MSGTYPE_RELINK_SESSION = "RELINK_SESSION"
)

type MsgType string
type InfoMsg struct {
	Type MsgType `json:"type"`
	Msg  string  `json:"msg"`
}

type SeqMsg struct {
	Type MsgType    `json:"type"`
	Data SeqMsgData `json:"data"`
}

type SeqMsgData struct {
	Seq []int `json:"seq"`
}

type W3cSessionMsg struct {
	Type MsgType           `json:"type"`
	Data W3cSessionMsgData `json:"data"`
}

type W3cSessionMsgData struct {
	ScoreMap  map[int]int  `json:"score_map"`  // 玩家分数
	Round     int          `json:"round"`      // 当前回合数
	ReadyInfo map[int]bool `json:"ready_info"` // 准备信息
	Seq       []int        `json:"seq"`
}

type RoundSessionMsg struct {
	Type MsgType          `json:"type"`
	Data RoundSessionData `json:"data"`
}

type RoundSessionData struct {
	PInfo         map[int]*PlayInfo `json:"pinfo"`          // 本局玩家信息 key id
	PLog          []string          `json:"plog"`           // 回合操作流水
	MaxBet        int               `json:"max_bet"`        // 当前轮注码(开牌值计算)
	CurrentPlayer int               `json:"current_player"` // id
}

type RelinkSessionMsg struct {
	Type     MsgType           `json:"type"`
	W3cData  W3cSessionMsgData `json:"w3c_data"`
	RsData   RoundSessionData  `json:"rs_data"`
	HandCard HandCard          `json:"hand_card"`
}

type ActionViewMsg struct {
	Type MsgType        `json:"type"`
	Data ActionViewData `json:"data"`
}

type ActionViewData struct {
	HandCard HandCard `json:"hand_card"`
}

type ViewLogMsg struct {
	Type MsgType     `json:"type"`
	Data ViewLogData `json:"data"`
}

type ViewLogData struct {
	HandCards map[int]HandCard `json:"hand_cards"`
	Winner    int              `json:"winner"`
}

// =========================================================

func GenInfoMsg(msg string) []byte {
	infoMsg := InfoMsg{
		Type: MSGTYPE_INFO,
		Msg:  msg,
	}
	bytes, err := json.Marshal(infoMsg)
	if err != nil {
		logrus.Errorf("序列化InfoMsg失败: %s", err.Error())
	}

	return bytes
}

func GenSeqMsg(seq []int) []byte {
	infoMsg := SeqMsg{
		Type: MSGTYPE_SEQ,
		Data: SeqMsgData{
			Seq: seq,
		},
	}
	bytes, err := json.Marshal(infoMsg)
	if err != nil {
		logrus.Errorf("序列化SeqMsg失败: %s", err.Error())
	}

	return bytes
}

func GenW3cSessionMsg(ws *W3cSession) []byte {
	wsMsg := W3cSessionMsg{
		Type: MSGTYPE_W3C_SESSION,
		Data: W3cSessionMsgData{
			ScoreMap:  ws.ScoreMap,
			Round:     ws.Round,
			ReadyInfo: ws.ReadyInfo,
			Seq:       ws.Players,
		},
	}
	bytes, err := json.Marshal(wsMsg)
	if err != nil {
		logrus.Errorf("序列化RoundSessionMsg失败: %s", err.Error())
	}
	return bytes
}

func GenW3cResultMsg(ws *W3cSession) []byte {
	wsMsg := W3cSessionMsg{
		Type: MSGTYPE_W3C_RESULT,
		Data: W3cSessionMsgData{
			ScoreMap: ws.ScoreMap,
			Round:    ws.Round,
		},
	}
	bytes, err := json.Marshal(wsMsg)
	if err != nil {
		logrus.Errorf("序列化RoundSessionMsg失败: %s", err.Error())
	}
	return bytes
}

func GenRoundSessionMsg(rs *RoundSession) []byte {
	rsMsg := RoundSessionMsg{
		Type: MSGTYPE_ROUND_SESSION,
		Data: RoundSessionData{
			PInfo:         rs.PInfo,
			PLog:          rs.PLog,
			MaxBet:        rs.MaxBet,
			CurrentPlayer: rs.Players[rs.current],
		},
	}
	bytes, err := json.Marshal(rsMsg)
	if err != nil {
		logrus.Errorf("序列化RoundSessionMsg失败: %s", err.Error())
	}
	return bytes
}

func GenRelinkSessionMsg(ws *W3cSession, id int) []byte {
	var handCard HandCard
	var currentPlayer int
	if ws.RoundSession.IsStart {
		if ws.RoundSession.PInfo[id].IsViewed {
			handCard = ws.RoundSession.handCards[id]
		}
		currentPlayer = ws.RoundSession.Players[ws.RoundSession.current]
	}

	relinkMsg := RelinkSessionMsg{
		Type: MSGTYPE_RELINK_SESSION,
		W3cData: W3cSessionMsgData{
			ScoreMap:  ws.ScoreMap,
			Round:     ws.Round,
			ReadyInfo: ws.ReadyInfo,
			Seq:       ws.Players,
		},
		RsData: RoundSessionData{
			PInfo:         ws.RoundSession.PInfo,
			PLog:          ws.RoundSession.PLog,
			MaxBet:        ws.RoundSession.MaxBet,
			CurrentPlayer: currentPlayer,
		},
		HandCard: handCard,
	}
	bytes, err := json.Marshal(relinkMsg)
	if err != nil {
		logrus.Errorf("序列化RelinkSessionMsg失败: %s", err.Error())
	}
	return bytes
}

func GenViewLogMsg(winner int, handCards map[int]HandCard) []byte {
	vlMsg := ViewLogMsg{
		Type: MSGTYPE_VIEW_LOG,
		Data: ViewLogData{
			Winner:    winner,
			HandCards: handCards,
		},
	}
	bytes, err := json.Marshal(vlMsg)
	if err != nil {
		logrus.Errorf("序列化ViewLogMsg失败: %s", err.Error())
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
		logrus.Errorf("序列化ActionViewMsg失败: %s", err.Error())
	}

	return bytes
}
