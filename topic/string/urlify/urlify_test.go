package urlify

import (
	"fmt"
	"testing"
)

func TestUrlifyFixOut(t *testing.T) {
	url := "duc  bph    "
	fmt.Println(URLifyFixedOutput(url, 10))
}

func TestUrlifyFixIn(t *testing.T) {
	url := "duc  bph    "
	fmt.Println(URLifyFixedInput(url, 10))
}
