package tst

import (
	"sync"
	"testing"
)

func TestPrint(t *testing.T) {
	a1 := [3]int{1, 4, 7}
	a2 := [3]int{2, 5, 8}
	a3 := [3]int{3, 6, 9}

	p1 := P{
		arr: a1,
		i:   0,
	}
	p2 := P{
		arr: a2,
		i:   0,
	}
	p3 := P{
		arr: a3,
		i:   0,
	}

	var wg sync.WaitGroup
	wg.Add(3)
	flag1, flag2, flag3 := make(chan struct{}), make(chan struct{}), make(chan struct{})

	go p1.Print(&wg, flag1, flag2)
	go p2.Print(&wg, flag2, flag3)
	go p3.Print(&wg, flag3, flag1)

	flag1 <- struct{}{}

	wg.Wait()
}
func TestPrint2(t *testing.T) {
	a1 := [3]int{1, 4, 7}
	a2 := [3]int{2, 5, 8}
	a3 := [3]int{3, 6, 9}

	p1 := P{
		arr: a1,
		i:   0,
		c:   1,
	}
	p2 := P{
		arr: a2,
		i:   0,
		c:   2,
	}
	p3 := P{
		arr: a3,
		i:   0,
		c:   0,
	}

	var wg sync.WaitGroup
	wg.Add(2)
	flag := make(chan int)

	go p1.Print2(&wg, flag)
	go p2.Print2(&wg, flag)
	go p3.Print2(&wg, flag)

	flag <- 1

	wg.Wait()
	<-flag
}

type P struct {
	arr [3]int
	i   int
	c   int
}

func (self P) Print(wg *sync.WaitGroup, start <-chan struct{}, end chan<- struct{}) {
	//for {
	//	<-start
	//	print(self.arr[self.i])
	//	self.i++
	//	end <- struct{}{}
	//	if self.i == 3 {
	//		wg.Done()
	//		return
	//	}
	//}

	defer wg.Done()
	for _, a := range self.arr {
		<-start
		print(a)
		if a != 9 {
			end <- struct{}{}
		}
	}

}

func (self P) Print2(wg *sync.WaitGroup, ch chan int) {
	for {
		if c := <-ch; c != self.c {
			ch <- c
			continue
		}
		print(self.arr[self.i])
		self.i++
		ch <- (self.c+1)%3
		if self.i == 3 && self.c != 0 {
			wg.Done()
			return
		}
	}
}
