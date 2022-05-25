package alg

import "testing"

func BenchmarkReverseList(b *testing.B) {
	var head *ListNode
	for i := 8; i > 0; i-- {
		head = &ListNode{i, head}
	}
	b.ResetTimer()
	for i:=0; i < b.N; i++{
		//head.String()
		head = ReverseList(head)
		//head.String()
	}
}

func TestReverseSeg(t *testing.T) {
	var head *ListNode

	for i := 8; i > 0; i-- {
		head = &ListNode{i, head}
	}
	head.String()
	head = ReverseList(head)
	head = ReverseSeg(head, 3)
	head = ReverseList(head)
	head.String()
}
