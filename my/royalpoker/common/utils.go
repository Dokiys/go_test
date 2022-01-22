package common

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
func RandAlphabetStr(n int) string {
	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = alphabet[rand.Int63() % int64(26)]
	}
	return string(b)
}

func RandNum(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}