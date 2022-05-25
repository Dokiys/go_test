package alg

import (
	"fmt"
	"math"
	"time"
)

// 使用协程查找数组里面某个元素
func Find(slice []int, target int, workCount int) (int, bool) {
	c := make(chan int)
	stop := make(chan int)
	length := len(slice)
	index := int(math.Ceil(float64(length) / float64(workCount)))
	for i := 0; i < workCount; i++ {
		var end int
		start := i * index
		if (i+1)*index >= (length - 1) {
			end = length - 1
		} else {
			end = (i + 1) * index
		}
		seg := slice[start:end]
		go gFind(stop, c, seg, target, start)
	}

	select {
	case i := <-c:
		return i, true
	case <-time.After(3 * time.Second):
		close(stop)
		return 0, false
	}
}

func gFind(stop chan int, c chan int, slice []int, target int, index int) {
	var count int
BEGIN:
	for i, v := range slice {
		select {
		case <-stop:
			break BEGIN
		default:
			if v == target {
				close(stop)
				c <- index + i
				break BEGIN
			}
		}
		count++
	}
	fmt.Println("goroutine find: ", count, "times")
}
