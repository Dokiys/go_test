package alg

import (
	"context"
	"fmt"
	"time"
)

// 使用协程查找数组里面某个元素
func Search(slice []int, target int) bool {
	done := make(chan bool)
	length := len(slice)
	index := length/2 + 1
	go gSearch(done, slice[:index], target)
	go gSearch(done, slice[index:], target)

	select {
	case <-done:
		return true
	case <-time.After(3 * time.Second):
		close(done)
		return false
	}
}

func gSearch(done chan bool, slice []int, target int) {
	var count int
BEGIN:
	for _, v := range slice {
		select {
		case <-done:
			break BEGIN
		default:
			if v == target {
				close(done)
				break BEGIN
			}
		}
		count++
	}
	fmt.Println("goroutine find: ", count, "times")
}

func Search2(slice []int, target int) (bool, error) {
	done := make(chan bool)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	length := len(slice)
	index := (length >> 1) + 1
	go gSearch2(ctx, done, slice[:index], target)
	go gSearch2(ctx, done, slice[index:], target)

	select {
	case <-ctx.Done():
		return false, ctx.Err()
	case <-done:
		cancel()
		return true, nil
	}
}

func gSearch2(ctx context.Context, done chan<- bool, slice []int, target int) {
	var count int
BEGIN:
	for _, v := range slice {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break BEGIN
		default:
			if v == target {
				done <- true
				break BEGIN
			}
		}
		//time.Sleep(1 * time.Second)
		count++
	}
	fmt.Println("goroutine find: ", count, "times")
}
