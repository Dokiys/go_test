package alg

import (
	"fmt"
	"sync"
)

// 启动两个协程交替输出字母和数字
func CrossPrint() {
	var wg sync.WaitGroup
	wg.Add(2)
	AlpSign := make(chan struct{})
	NumSign := make(chan struct{})

	go printNum(&wg, AlpSign, NumSign)
	go printAlphabet(&wg, AlpSign, NumSign)
	AlpSign <- struct{}{}

	wg.Wait()
}

func printAlphabet(wg *sync.WaitGroup, AlpSign <-chan struct{}, NumSign chan<- struct{}) {
	for i := 'A'; i <= 'Z'; i += 2 {
		<-AlpSign
		fmt.Printf("%c", i)
		fmt.Printf("%c ", i+1)
		NumSign <- struct{}{}
	}
	<-AlpSign
	wg.Done()
}

func printNum(wg *sync.WaitGroup, AlpSign chan<- struct{}, NumSign <-chan struct{}) {
	for i := 1; i <= 26; i += 2 {
		<-NumSign
		fmt.Print(i)
		fmt.Print(i+1, " ")
		AlpSign <- struct{}{}
	}

	wg.Done()
}