package inner

import "testing"

// TestBitOperation 位运算
func TestOperatorBitOperation(t *testing.T) {
	i,j := 8, 7
	t.Log(i >> 2)	// 2
	t.Log(j >> 2)	// 1
}