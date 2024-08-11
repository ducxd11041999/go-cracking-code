package oneway

import (
	"fmt"
	"testing"
)

func TestOneWay1(t *testing.T) {
	str1 := "pale"
	str2 := "bale"

	fmt.Println(OneWay(str1, str2))
}

func TestOneWay2(t *testing.T) {
	str2 := "dpale"
	str1 := "dale"

	fmt.Println(OneWay(str1, str2))
}

func TestOneWayCommon(t *testing.T) {
	str2 := "dpale"
	str1 := "dale"

	fmt.Println(OneWayCommon(str1, str2))
}

func TestOneWayCommon2(t *testing.T) {
	str2 := "dpalea"
	str1 := "dale"

	fmt.Println(OneWayCommon(str1, str2))
}
