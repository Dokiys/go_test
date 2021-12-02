package std

import "testing"

func TestChanOverAdd(t *testing.T) {
	var ch = make(chan struct{}, 2)
	ch <- struct{}{}
	ch <- struct{}{}
	ch <- struct{}{}
}
