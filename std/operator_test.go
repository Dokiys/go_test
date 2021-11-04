package std

import (
	"math/rand"
	"testing"
	"time"
)

// TestOperatorBitOperation 位运算
func TestOperatorBitOperation(t *testing.T) {
	i,j := 8, 7
	t.Log(i >> 2)	// 2
	t.Log(j >> 2)	// 1
}

// TestOperatorOver 位移超出
func TestOperatorOver(t *testing.T) {
	var v int64

	//v = 1 << 63		// constant 9223372036854775808 overflows int64
	v = 1 << 62
	t.Log(v)
}

func RandNum(n int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(n)
}