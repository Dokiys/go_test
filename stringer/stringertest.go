//go:generate stringer -type=Num
package stringer

import (
	"fmt"
)

type Num int
const (
	One Num = iota
	Two Num = iota
	Three Num = iota
	Four Num = iota
	Five Num = iota
)

func TestA() {
	fmt.Println(One)
	fmt.Println(Two)
	fmt.Println(Three)
	fmt.Println(Four)
	fmt.Println(Five)
}
