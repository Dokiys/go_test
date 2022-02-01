module royalpoker

go 1.16

replace github.com/dokiy/royalpoker/common => ./common

replace github.com/dokiy/royalpoker/win3cards => ./win3cards

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dokiy/royalpoker/common v1.0.0
	github.com/dokiy/royalpoker/win3cards v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.7
	github.com/gorilla/websocket v1.4.2
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	gopkg.in/yaml.v2 v2.2.8 // indirect
)
