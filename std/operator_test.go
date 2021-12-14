package std

import (
	"testing"
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

// TestOperatorFlag 判断int类型首位
func TestOperatorFlag(t *testing.T) {
	//i := 1
	var i int64 = -100
	t.Logf("%d", 2 + i>>64)
}