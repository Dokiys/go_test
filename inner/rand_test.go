package inner

import (
	"math/rand"
	"testing"
	"time"
)

// TestRand 生成随机整数
func TestRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	v := rand.Int31n(3)
	j := rand.Int31n(30)
	t.Log(v)
	t.Log(j)
}