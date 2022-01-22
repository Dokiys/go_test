package poker

import (
	"context"
	"github.com/pkg/errors"
	"sync"
)

const baseRound = 3

type Hub struct {
	Owner   int
	Players map[int]*Player
	//Broadcast chan []byte

	hubSession *HubSession
	register   chan *Player
	unregister chan *Player
	start      chan struct{}
}
type HubSession struct {
	Players     map[int]*Player
	count       int          // 玩家人数
	Round       int          // 当前回合数
	ScoreMap    map[int]int  // 玩家分数 key playerId
	Seq         []int        // 玩家顺序
	ReadyInfo   map[int]bool // 准备信息， 全部准备表示已经开局
	PlaySession *RoundSession

	l sync.Mutex
}

func (h *Hub) isStarted() bool {
	_, r := <-h.start
	return !r
}

func NewHub(id int) *Hub {
	return &Hub{
		Owner:   id,
		Players: make(map[int]*Player),
		//Broadcast:  make(chan []byte),
		register:   make(chan *Player),
		unregister: make(chan *Player),
		start:      make(chan struct{}),
	}
}

func (h *Hub) Run() error {
	ctx := context.Background()
	for {
		select {
		case player := <-h.register:
			if h.isStarted() {
				return errors.New("已开局，无法加入游戏！")
			}
			h.Players[player.Id] = player
		case player := <-h.unregister:
			if h.isStarted() {
				return errors.New("已开局，无法退出游戏！")
			}
			if _, ok := h.Players[player.Id]; ok {
				delete(h.Players, player.Id)
				close(player.send)
			}
		case _, ok := <-h.start:
			if !ok {
				return errors.New("获取开始状态信息失败！")
			}
			err := h.InitHubSession()
			if err != nil {
				return errors.Wrapf(err, "初始化开局信息失败！")
			}
			err = h.Start(ctx)
			if err != nil {
				return errors.Wrapf(err, "开局失败！")
			}

			// 发送结果给所有玩家
			// TODO[Dokiy] 2022/1/21:

		}
	}
}

func (h *Hub) InitHubSession() error {
	hubs := &HubSession{}
	// Init
	length := len(h.Players)
	hubs.Players = h.Players
	hubs.count = length
	hubs.ReadyInfo = make(map[int]bool, length)
	hubs.Seq = make([]int, length)

	// 开局
	i := 0
	for id, _ := range hubs.Players {
		hubs.Seq[i] = id
		hubs.ScoreMap[id] = 0
		hubs.ReadyInfo[id] = true

		i++
	}
	hubs.PlaySession = NewRoundSession(h.Players)
	return nil
}

func (h *Hub) Start(ctx context.Context) error {
	close(h.start)
	return errors.Wrapf(h.hubSession.Run(ctx), "开局失败：")
}

func (hubs *HubSession) Run(ctx context.Context) error {
	for r := 0; r < hubs.count*baseRound; r++ {
		hubs.Round = r

		l := len(hubs.Seq)
		seq := make([]int, l)
		for i, id := range hubs.Seq {
			seq[(r+i)%l] = id
		}
		winner, err := hubs.PlaySession.Play(ctx, seq)
		if err != nil {
			return err
		}

		hubs.settle(winner)
	}
	return nil
}

func (hubs *HubSession) settle(winner *Player) {
	hubs.l.Lock()
	{
		var bet int
		for id, info := range hubs.PlaySession.PInfo {
			bet += info.Score
			hubs.ScoreMap[id] -= info.Score
		}
		hubs.ScoreMap[winner.Id] += bet
	}
	hubs.l.Unlock()
}
