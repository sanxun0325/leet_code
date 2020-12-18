package __AddTwoNumbers

import (
	"fmt"
	"testing"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil || l2 == nil {
		return nil
	}

	//创建待返回链表
	var l3 = &ListNode{Val: -1}
	cur := l3
	//进位
	carry := 0

	for l1 != nil && l2 != nil {
		k := l1.Val + l2.Val + carry
		//计算进位
		carry = k / 10
		//计算val
		sum := k % 10

		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
		if l1 != nil && l2 == nil {
			l2 = &ListNode{Val: 0}
		}
		if l2 != nil && l1 == nil {
			l1 = &ListNode{Val: 0}
		}
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
		cur = cur.Next
	}

	return l3.Next
}

func TestRun(t *testing.T) {
	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}
	l2 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}
	l3 := addTwoNumbers(l1, l2)
	fmt.Println(l3)
}
