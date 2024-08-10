package palindrome

func isPalindrome(str string) bool {
	counter := map[byte]int{}
	lenStr := len(str)
	for idx := 0; idx < lenStr; idx++ {
		counter[str[idx]]++
	}

	onceOdd := false
	for _, val := range counter {
		if val%2 == 1 {
			if onceOdd {
				return false
			}

			onceOdd = true
		}
	}

	return true
}

func isPalindromeIgnoreSpace(str string) bool {
	counter := map[byte]int{}
	lenStr := len(str)
	for idx := 0; idx < lenStr; idx++ {
		counter[str[idx]]++
	}

	onceOdd := false
	for char, val := range counter {
		if val%2 == 1 && char != ' ' {
			if onceOdd {
				return false
			}

			onceOdd = true
		}
	}

	return true
}

func isPalindromeOptimize(str string) bool {
	counter := map[byte]int{}
	lenStr := len(str)
	counterOdd := 0
	for idx := 0; idx < lenStr; idx++ {
		counter[str[idx]]++
		if counter[str[idx]]%2 == 1 {
			counterOdd++
		} else {
			counterOdd--
		}
	}

	return counterOdd <= 1
}

func isParalindromeBitMask(str string) bool {
	bitCheker := 0
	lenStr := len(str)
	for idx := 0; idx < lenStr; idx++ {
		value := str[idx] - 'a'
		if value < 0 {
			return false
		}

		mask := 1 << value
		if bitCheker&mask == 0 {
			// case value is not check -> set bit to 1 is mark as checked
			bitCheker |= mask
		} else {
			bitCheker &= ^mask
		}
	}

	return (bitCheker & (bitCheker - 1)) == 0
}
