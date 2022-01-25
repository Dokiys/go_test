package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHub(t *testing.T) {
	ctx := context.Background()
	hub := NewHub()

	// 注册
	player := NewLocalPlayer()
	//a := player.Receive(ctx)
	//t.Log(a)

	err := hub.Register(player)
	assert.NoError(t, err)

	err = hub.Start(ctx)
	assert.NoError(t, err)
}
