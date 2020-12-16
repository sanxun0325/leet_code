package Add_Two_Numbers

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

	var len1, len2, count int

	if l1 == nil || l2 == nil {
		return nil
	}

	point1 := l1
	point2 := l2
	for nil != point1.Next {
		point1 = point1.Next
		len1++
	}

	for nil != point2.Next {
		point2 = point2.Next
		len2++
	}

	if len1 > len2 {
		for i := 0; i < len1-len2; i++ {
			point2.Next = &ListNode{
				Val: 0,
			}
			point2 = point2.Next
		}
		count = len1
	} else if len2 > len1 {
		for i := 0; i < len2-len1; i++ {
			point1.Next = &ListNode{
				Val: 0,
			}
			if point1.Next != nil {
				point1 = point1.Next
			}
		}
		count = len2
	} else {
		count = len1
	}
	var l3 = &ListNode{Val: -1}
	cur := l3
	carry := 0
	for i := 0; i < count+1; i++ {
		k := l1.Val + l2.Val + carry
		carry = k / 10
		sum := k % 10
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
		if i == count && carry>0{
			cur.Next = &ListNode{Val: carry}
			cur = cur.Next
		}
	}

	return l3.Next
}

func TestRun(t *testing.T) {
	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}
	l2 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}
	l3 := addTwoNumbers(l1, l2)
	fmt.Println(l3)
}
