module royalpoker

go 1.16

require (
	github.com/dokiy/royalpoker/common v1.0.0
	github.com/dokiy/royalpoker/win3cards v1.0.0
	github.com/gorilla/websocket v1.4.2
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
)

replace (
	github.com/dokiy/royalpoker/common v1.0.0 => ./common
	github.com/dokiy/royalpoker/win3cards v1.0.0 => ./win3cards

)