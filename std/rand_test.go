package std

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

const rand_alphabet = "abcdefghijklmnopqrstuvwxyz"
func TestRandAlphabetStr(t *testing.T) {
	n := 8
	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = rand_alphabet[rand.Int63() % int64(26)]
	}

	t.Log(string(b))
}