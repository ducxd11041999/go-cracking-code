package unique

import (
	"fmt"
	"testing"
)

func TestUnique(t *testing.T) {
	str := "abc"
	fmt.Println(IsUniqueElementUsingArray(str))
	str2 := "abc@@"
	fmt.Println(IsUniqueElementUsingArray(str2))
}

func TestUniqueMap(t *testing.T) {
	str := "abc"
	fmt.Println(IsUniqueElementUsingMap(str))
	str2 := "abc@@"
	fmt.Println(IsUniqueElementUsingMap(str2))
}

func TestUniqueBitsMap(t *testing.T) {
	str := "abc"
	fmt.Println(IsUniqueElementUsingBitsMap(str))
	str2 := "abca"
	fmt.Println(IsUniqueElementUsingBitsMap(str2))
}
