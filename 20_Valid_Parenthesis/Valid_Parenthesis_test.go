package _0_Valid_Parenthesis

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isValid(s string) bool {
	stack := list.New()
	smap := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, str := range s {
		bStr := byte(str)
		if stack.Front() == nil {
			stack.PushBack(bStr)
		} else {
			b := smap[stack.Front().Value.(byte)]
			if bStr == b {
				stack.Remove(stack.Front())
			} else {
				stack.PushFront(bStr)
			}
		}

	}
	if stack.Front() != nil {
		return false
	}
	return true
}

func TestIsValid(t *testing.T) {
	s := "{[]}"
	ok := isValid(s)
	assert.True(t, ok)
}
