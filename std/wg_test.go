package std

import (
	"sync"
	"testing"
)

// TestWg 测试wg在Wait()方法返回之前，调用Add()会报错
func TestWg(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan struct{})
	wg.Add(1)
	go func() {
		<- ch
		wg.Wait()	// 1
		t.Logf("123")
	}()
	go func() {
		<- ch
		wg.Done()	// 2
		wg.Add(1) // 3
		t.Logf("321")
	}()
	close(ch)
	for{}
}