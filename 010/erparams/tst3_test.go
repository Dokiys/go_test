package main

import (
	"fmt"
	"sync"
	"testing"
)

type Node struct {
	Val  interface{} // 节点上保存的业务数据
	Next *Node       // 指向后一个节点
}

type List struct {
	// 单链表的定义
	head *Node
	end  *Node
	c    int
}

func NewNode(v interface{}) *Node {
	// 创建新的节点
	return &Node{Val: v, Next: nil}
}

func NewList() *List {
	// 创建新的链表
	return &List{
		head: nil,
		end:  nil,
		//c:    0,
	}
}

func (l *List) Add(es ...interface{}) {
	// 添加元素，别忘了先定位到尾部节点
	for _, e := range es {
		// TODO[Dokiy] 2022/5/10: 类型断言
		//value, ok := e.(*List)
		if l.head == nil {
			n := &Node{Val: e}
			l.head = n
			l.end = n
			continue
		}
		l.end.Next = NewNode(e)
		l.end = l.end.Next
		//l.c++
	}
}

func (l *List) Values(f func(i interface{})) {
	// 遍历单链表的节点，消费节点上的业务数据
	c := l.head
	for c != nil {
		f(c.Val)
		c = c.Next
	}
}

func (l *List) Reverse() {
	// 反转链表的节点(加分项，若时间不够写出反转思路即可)
	c := l.head
	l.end = l.head
	l.head = nil
	for c != nil {
		n := c.Next
		c.Next = nil

		r := c
		r.Next = l.head
		l.head = r

		c = n
	}

}

func f(v interface{}) {
	fmt.Print(v)
}

func TestList(t *testing.T) {
	l := NewList()

	// 添加业务数据到链表上
	l.Add(1, 2, 3)
	l.Add(4, 5, 6)
	l.Add(7, 8, 9, 10)

	// 定义 f 函数，可以把 f 是消费者，l 是生产者
	l.Values(f) // 1, 2, 3, 4, 5, 6, 7, 8, 9, 10

	// 加分项
	l.Reverse()
	l.Values(f) // 10, 9, 8, 7, 6, 5, 4, 3, 2, 1
}

func TestName(t *testing.T) {
	var wg sync.Mutex

	wg.Lock()
	wg2:=wg
	wg2.Unlock()
}