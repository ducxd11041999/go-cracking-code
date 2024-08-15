package palindrome

import (
	"fmt"
	model "my-code/go-cracking-code/topic/link-list/base"
	"testing"
)

func TestSimple(t *testing.T) {
	l1 := model.NewList(7)
	s1 := l1
	l1 = l1.AppendInt(1)
	l1 = l1.AppendInt(6)
	l1 = l1.AppendInt(1)
	l1 = l1.AppendInt(7)
	// l1 = l1.AppendInt(9)

	fmt.Println(Palindrome(s1))
	fmt.Println(IsPalindromeUsingStack(s1))
}
