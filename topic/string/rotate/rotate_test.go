package rotate

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	in1 := "waterbotle"
	in2 := "erbotlewat"

	fmt.Println(isRotate(in1, in2))
}
