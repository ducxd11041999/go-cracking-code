package model

import (
	"testing"
)

func TestModel(t *testing.T) {
	n := 6
	var list = NewList(0)
	itor := list
	for i := 1; i < n; i++ {
		itor = itor.AppendInt(i)
	}

	list.Print()
}

func TestAppendIntSelf(t *testing.T) {
	n := 6
	var list = NewList(0)
	itor := list
	for i := 1; i < n; i++ {
		itor.AppendIntSelf(i)
	}

	list.Print()
}
