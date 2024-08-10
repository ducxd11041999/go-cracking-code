package sumlist

import (
	model "my-code/go-cracking-code/topic/link-list/base"
	"testing"
)

func TestSumList(t *testing.T) {
	l1 := model.NewList(7)
	s1 := l1
	l1 = l1.AppendInt(1)
	l1 = l1.AppendInt(6)

	l2 := model.NewList(5)
	s2 := l2
	l2 = l2.AppendInt(9)
	l2 = l2.AppendInt(2)

	result := sum2List(s1, s2, 0)
	result.Print()
}
