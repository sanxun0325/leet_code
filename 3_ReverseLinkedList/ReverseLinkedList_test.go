package __ReverseLinkedList

import (
	"fmt"
	"testing"
	"time"
)

type ListNode struct {
	Next *ListNode
	val  int
}

func reverseList(head *ListNode) *ListNode {
	fmt.Println(1.0e9)
	time.Sleep(1.0e9)
	var prev *ListNode
	cur := head
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}

func TestReverseList(t *testing.T) {
	head := &ListNode{val: 1, Next: &ListNode{val: 2, Next: &ListNode{val: 3, Next: &ListNode{val: 4, Next: &ListNode{val: 5}}}}}
	fmt.Println(reverseList(head))
}
