package std

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// TestMutexRW 读写锁
func TestMutexRW(t *testing.T) {
	var rw sync.RWMutex
	a := 1
	go func() {
		// 保证在写锁之后执行
		time.Sleep(1 * time.Second)

		rw.RLock()
		time.Sleep(1 * time.Second)
		//fmt.Println(a)
		print(a)
		rw.RUnlock()
	}()
	go func() {
		rw.Lock()
		// do something...
		time.Sleep(2 * time.Second)
		a++
		// do more thing...
		rw.Unlock()
	}()
	for {
	}
}

// TestMutexRWDead 不可重入
func TestMutexRWDead(t *testing.T) {
	var rw sync.RWMutex
	rw.RLock()
	rw.Lock()
}

// TestMutexRWReentrant 读写锁
func TestMutexRWReentrant(t *testing.T) {
	var rw sync.RWMutex
	a := 1
	go func() {
		// 保证在写锁之后执行
		time.Sleep(1 * time.Second)
		rw.RLock()
		fmt.Println(a)
		//print(a)
		rw.RUnlock()
	}()
	go func() {
		rw.Lock()
		// do something...
		time.Sleep(2 * time.Second)
		a++
		// do more thing...
		rw.Unlock()
	}()
	for {
	}
}

// TestMutexCond Single and Broadcast
func TestMutexCond(t *testing.T) {
	var cond = sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(10)

	flag := false

	for i := 0; i < 10; i++ {
		go func(i int) {
			cond.L.Lock()
			fmt.Printf("goroutine%d get lock\n", i)
			if !flag {
				cond.Wait()
			}
			fmt.Printf("goroutine%d: hello!\n", i)
			cond.L.Unlock()
			wg.Done()
		}(i)
	}

	time.Sleep(1 * time.Second)
	cond.L.Lock()
	flag = true
	fmt.Printf("main goroutine: i'm coming!\n")
	cond.L.Unlock()
	//cond.Broadcast()
	cond.Signal()
	wg.Wait()
}
