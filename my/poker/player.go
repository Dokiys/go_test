package poker

import "github.com/gorilla/websocket"

type Player struct {
	Id  int
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}
