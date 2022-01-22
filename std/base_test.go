package std

import (
	"sync"
	"testing"
)

func one() int {
	return 1
}
func f() (i int) {
	return one()
}

func TestBase(t *testing.T) {
	t.Log(f())
}

func TestBlock(t *testing.T) {
	var l sync.Mutex
	{
		l.Lock()
		defer l.Unlock()
		t.Log("123")
	}

	l.Lock()
	defer l.Unlock()
	t.Log("abc")
}
