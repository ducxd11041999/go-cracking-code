package permutation

import (
	"fmt"
	"testing"
)

// use print instead assert

func TestPermutation(t *testing.T) {
	in1 := "abcd"
	in2 := "dbca"

	fmt.Println(IsPermutationWithSort(in1, in2))
}

func TestPermutation2(t *testing.T) {
	in1 := "abcde"
	in2 := "dbca"

	fmt.Println(IsPermutationWithSort(in1, in2))
}

func TestPermutationCounter(t *testing.T) {
	in1 := "abcd"
	in2 := "dbca"

	fmt.Println(IsPermutationWithCount(in1, in2))
}

func TestPermutationCounter2(t *testing.T) {
	in1 := "abcde"
	in2 := "dbcae"

	fmt.Println(IsPermutationWithCount(in1, in2))
}
