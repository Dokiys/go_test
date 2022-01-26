package main

import (
	"fmt"
	"github.com/dokiy/royalpoker/common"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

// ====================================================
type LoginRequest struct {
	Id       int    `json:"id"`
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
	Token string `json:"token"`
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
	c.JSON(200, reply)
}
func (self *handler) login(c *gin.Context) {
	replyHandler(c, func() (interface{}, error) {
		var login LoginRequest
		err := c.BindJSON(&login)
		if err != nil {
			return nil, errors.Wrapf(err, "传入参数错误！")
		}

		user, ok := userMap[login.Id]
		if !ok {
			return nil, errors.Wrapf(err, "玩家不存在！请联系张老板开账号！")
		}

		if user.Password != login.Password {
			return nil, errors.New("密码错误！")
		}

		var isAdmin bool
		if login.Id == 1 {
			isAdmin = true
		}
		token, err := common.Encode(&common.Claims{
			Uid:     user.Id,
			Name:    user.Name,
			IsAdmin: isAdmin,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "生成token错误！")
		}
		return LoginReply{Token: token}, nil
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

		// TODO[Dokiy] 2022/1/26: 这里post请求不清楚可不可以连接
		upGrader := websocket.Upgrader{}
		ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return nil, errors.Wrapf(err, "系统错误：创建websocket链接失败")
		}

		player := NewPlayerWs(ws, user.Id, user.Name)
		err = hub.Register(player)
		if err != nil {
			return nil, errors.Wrapf(err, "加入房间失败！")
		}
		hub.BroadcastHubSession(c, fmt.Sprintf("玩家【%s】加入了游戏", user.Name))
		userHubMap[user.Id] = hub.Id

		return nil, nil
	})
}

func (self *handler) outHub(c *gin.Context) {
	replyHandler(c, func() (interface{}, error) {
		// TODO[Dokiy] 2022/1/26:
	})
}

func (self *handler) startHub(c *gin.Context) {
	replyHandler(c, func() (interface{}, error) {
		// TODO[Dokiy] 2022/1/26:
	})
}
