package alg

import "fmt"

type ListNode struct {
	i    int
	Next *ListNode
}

func (list *ListNode) String() {
	for list.Next != nil {
		fmt.Printf("%v -> ", list.i)
		list = list.Next
	}
	fmt.Printf("%v\n", list.i)
}

// 根据指定常数反转链表，比如1->2->3->4->5->6
// 指定2，输出：2->1->4->3->6->5
// 指定3，输出：3->2->1->6->5->4
func ReverseSeg(head *ListNode, c int) (result *ListNode) {
	curr := head
	for i := 1; i < c && curr != nil; i++ {
		curr = curr.Next
	}
	if curr == nil {
		return head
	}

	tmp := curr.Next
	curr.Next = nil
	result = ReverseList(head)
	head.Next = ReverseSeg(tmp, c)
	return
}

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode

	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
