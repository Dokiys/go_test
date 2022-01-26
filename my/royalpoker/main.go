package main

import (
	"fmt"
	"github.com/dokiy/royalpoker/common"
	"github.com/gin-gonic/gin"
)

var (
	userMap    map[int]*User
	userHubMap map[int]int
)

type User struct {
	Id       int
	Name     string
	Password string
}

func init() {
	userMap = make(map[int]*User, 10)
	userMap[1] = &User{
		Id:       1,
		Name:     "zhangsan",
		Password: "123",
	}

	userHubMap = make(map[int]int)
}

func TokenHandle(c *gin.Context) {
	if c.Request.URL.Path != "/login" {
		token := c.Request.Header.Get("Authorization")
		decode, err := common.Decode(token)
		if err != nil {
			c.JSON(401, fmt.Sprintf("授权信息错误：%s", err.Error()))
		}
		c.Set("username", decode.Name)
		c.Set("isAdmin", decode.IsAdmin)
	}
	c.Next()
}
func main() {
	r, handler := gin.Default(), NewHandler()
	r.Use(TokenHandle)
	r.POST("/login", handler.login)
	r.POST("/hub/create", handler.createHub)
	r.POST("/hub/join", handler.joinHub)
	r.POST("/hub/out", handler.outHub)
	r.POST("/hub/start", handler.startHub)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
