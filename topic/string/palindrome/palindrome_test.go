package palindrome

import (
	"fmt"
	"testing"
)

func TestPalindrome(t *testing.T) {
	in1 := "taco cat"
	fmt.Println(isPalindrome(in1))
}

func TestPalindromeIgnoreSpace(t *testing.T) {
	in1 := "taco cat"
	fmt.Println(isPalindromeIgnoreSpace(in1))
}

func TestPalindromeOptimize(t *testing.T) {
	in1 := "tacocat"
	fmt.Println(isPalindromeOptimize(in1))
	in2 := "abccbb"
	fmt.Println(isPalindromeOptimize(in2))
}

func TestPalindromeBitMask(t *testing.T) {
	in1 := "tacocat"
	fmt.Println(isParalindromeBitMask(in1))
	in2 := "abccbb"
	fmt.Println(isParalindromeBitMask(in2))
}
