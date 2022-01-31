package main

import (
	"fmt"
	"github.com/dokiy/royalpoker/common"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strconv"
)

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

// ====================================================
type LoginRequest struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type HubRequest struct {
	Id int `json:"id"`
}
type Reply struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type LoginReply struct {
	Token    string `json:"token"`
	Id       int    `json:"id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	HubId    int    `json:"hub_id"`
}
type HubReply struct {
	Id int `json:"id"`
}

// ====================================================
func replyHandler(c *gin.Context, f func() (interface{}, error)) {
	data, err := f()
	var reply = Reply{Data: data}
	if err != nil {
		reply.Code = 1001
		reply.Msg = err.Error()
	}
	logrus.Info(userHubMap)
	c.JSON(200, reply)
}
func (self *handler) login(c *gin.Context) {
	replyHandler(c, func() (interface{}, error) {
		var login LoginRequest
		err1 := c.ShouldBindJSON(&login)

		if err1 == nil {
			var token string
			uid := userNameMap[login.Username]
			user, ok := userMap[uid]
			if !ok {
				return nil, errors.New("账号不存在！")
			}

			if user.Password != login.Password {
				return nil, errors.New("密码错误！")
			}

			token, err1 = common.Encode(&common.Claims{
				Uid:     user.Id,
				Name:    user.Name,
				IsAdmin: user.IsAdmin,
			})
			if err1 != nil {
				return nil, errors.Wrapf(err1, "生成token错误！")
			}

			return LoginReply{
				Token:    token,
				Id:       user.Id,
				Username: user.Name,
				IsAdmin:  user.IsAdmin,
				HubId:    userHubMap[user.Id],
			}, nil
		}

		token := c.Request.Header.Get("Authorization")
		decode, err := common.Decode(token)
		if err != nil {
			return nil, errors.Wrapf(err, "身份认证失败！")
		}

		return LoginReply{
			Token:    token,
			Id:       decode.Uid,
			Username: decode.Name,
			IsAdmin:  decode.IsAdmin,
			HubId:    userHubMap[decode.Uid],
		}, nil
	})
}
func (self *handler) createHub(c *gin.Context) {
	replyHandler(c, func() (interface{}, error) {
		uid, _ := c.Get("uid")
		user := userMap[uid.(int)]
		hubId, ok := userHubMap[user.Id]
		if ok {
			return HubReply{Id: hubId}, nil
		}

		hub := NewHub(user.Id)
		userHubMap[user.Id] = hub.Id
		return HubReply{Id: hub.Id}, nil
	})
}
func (self *handler) joinHub(c *gin.Context) {
	upGrader := websocket.Upgrader{
		Subprotocols: []string{c.Request.Header.Get("Sec-WebSocket-Protocol")},
	}
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(1000, err.Error()))
			ws.Close()
		}
	}()

	token := c.Request.Header.Get("Sec-WebSocket-Protocol")
	decode, err := common.Decode(token)
	if err != nil {
		err = errors.Wrapf(err, "授权错误！")
		return
	}

	uid := decode.Uid
	user := userMap[uid]
	hubId, err := strconv.Atoi(c.Query("hubid"))
	if err != nil {
		err = errors.New("传入参数错误！")
		return
	}
	if hubId == 0 {
		hubId = userHubMap[user.Id]
	}
	hub, ok := GetHub(hubId)
	if !ok {
		err = errors.New("房间不存在！")
		return
	}

	if player, ok := hub.Players[user.Id]; ok {
		player.SetConn(c, ws)
		hub.BroadcastHubSession(c, fmt.Sprintf("玩家【%s】正在重连游戏", user.Name))
		hub.InfoPlayerRelinkSession(c, player.GetId())
		return
	}
	player := NewPlayerWs(ws, user.Id, user.Name)
	err = hub.Register(player)
	if err != nil {
		err = errors.Wrapf(err, "加入房间失败！")
		return
	}
	hub.BroadcastHubSession(c, fmt.Sprintf("玩家【%s】加入了游戏", user.Name))
	userHubMap[user.Id] = hub.Id

	select {
	case <-player.close:
		return
	}
}
func (self *handler) outHub(c *gin.Context) {
	replyHandler(c, func() (interface{}, error) {
		uid, _ := c.Get("uid")
		user := userMap[uid.(int)]
		hub, ok := GetHub(userHubMap[user.Id])
		if !ok {
			return nil, errors.New("房间不存在！")
		}

		err := hub.Unregister(user.Id)
		if err != nil {
			return nil, errors.Wrapf(err, "退出房间失败")
		}

		delete(userHubMap, user.Id)
		if len(hub.Players) <= 0 && !hub.IsStarted {
			hub.Close(false)
		}
		return nil, nil
	})
}

func (self *handler) startHub(c *gin.Context) {
	replyHandler(c, func() (interface{}, error) {
		var req HubRequest
		uid, _ := c.Get("uid")
		user := userMap[uid.(int)]
		err := c.BindJSON(&req)
		if err != nil {
			return nil, errors.Wrapf(err, "传入参数错误！")
		}
		var hubId = req.Id
		if req.Id == 0 {
			hubId = userHubMap[user.Id]
		}
		hub, ok := GetHub(hubId)
		if !ok {
			return nil, errors.New("房间不存在！")
		}

		if hub.Owner != user.Id {
			return nil, errors.New("只有房主才能开始游戏！")
		}

		go func() {
			err := hub.Start()
			if err != nil {
				hub.BroadcastHubSession(c, err.Error())
			}
			hub.Close(true)
		}()

		return nil, nil
	})
}

func (self *handler) closeHub(c *gin.Context) {
	replyHandler(c, func() (interface{}, error) {
		uid, _ := c.Get("uid")
		user := userMap[uid.(int)]
		if !user.IsAdmin {
			return nil, errors.New("无权操作！")
		}

		hub, ok := GetHub(userHubMap[user.Id])
		if !ok {
			return nil, errors.New("房间不存在！")
		}

		for id, _ := range hub.Players {
			delete(userHubMap, id)
		}
		hub.Close(true)
		return nil, nil
	})
}
