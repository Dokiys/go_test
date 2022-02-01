package main

import (
	"embed"
	"flag"
	"fmt"
	"github.com/dokiy/royalpoker/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"net/http"
)

var (
	userMap     map[int]*User
	userNameMap map[string]int
	userHubMap  map[int]int
)

type User struct {
	Id       int	`yaml:"id"`
	Name     string	`yaml:"name"`
	Password string	`yaml:"password"`
	IsAdmin  bool	`yaml:"is_admin"`
}

func init() {
	userMap = make(map[int]*User, 10)
	userNameMap = make(map[string]int, 10)
	userHubMap = make(map[int]int, 3)
	userMap[999] = &User{
		Id:       999,
		Name:     "zhangsan",
		Password: "123",
		IsAdmin:  true,
	}
	userNameMap["zhangsan"] = 999

	var users []*User
	file, err := ioutil.ReadFile("./user.conf")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, &users)
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		userMap[user.Id] = user
		userNameMap[user.Name] = user.Id
	}
}

var addr = flag.String("addr", "localhost:8080", "http service address")

//go:embed templates
var FS embed.FS
func main() {
	flag.Parse()

	r, handler := gin.Default(), NewHandler()
	tmpl := template.Must(template.New("").ParseFS(FS, "templates/*.html"))
	r.SetHTMLTemplate(tmpl)
	//fe, _ := fs.Sub(FS, "static")
	//r.StaticFS("templates/static", http.FS(fe))

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", home)
	r.POST("/login", handler.login)
	r.GET("/hub/join", handler.joinHub)

	r.POST("/hub/create", TokenHandle, handler.createHub)
	r.POST("/hub/out", TokenHandle, handler.outHub)
	r.POST("/hub/start", TokenHandle, handler.startHub)
	r.POST("/hub/closeHub", TokenHandle, handler.closeHub)

	err := r.Run(*addr)
	if err != nil {
		panic(err)
	}
}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"addr": *addr})
}

func TokenHandle(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	decode, err := common.Decode(token)
	if err != nil {
		c.JSON(401, fmt.Sprintf("授权信息错误：%s", err.Error()))
		c.Abort()
		return
	}
	if _, ok := userMap[decode.Uid]; !ok {
		c.JSON(401, fmt.Sprintf("授权信息错误：未找到该用户"))
		c.Abort()
		return
	}
	c.Set("uid", decode.Uid)
	c.Set("username", decode.Name)
	c.Set("isAdmin", decode.IsAdmin)
	c.Next()
}