package compress

import (
	"fmt"
	"testing"
)

func TestCompress(t *testing.T) {
	str := "aaaaaabcccccd"
	fmt.Println(Compress(str))
}

func TestCompress2(t *testing.T) {
	str := "abc"
	fmt.Println(Compress(str))
}

func TestCompressByte(t *testing.T) {
	str := "aaaaaabcccccd"
	fmt.Println(CompressByte(str))
}

func TestCompressByte2(t *testing.T) {
	str := "abc"
	fmt.Println(CompressByte(str))
}
