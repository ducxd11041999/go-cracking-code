package rotate

func isSubString(str1, str2 string) bool {
	n := len(str1)
	m := len(str2)

	if m == 0 {
		return false
	}

	if m > n {
		return false
	}

	for i := 0; i <= n-m; i++ {
		if str1[i:i+m] == str2 {
			return true
		}
	}

	return false
}

func isRotate(s1, s2 string) bool {
	l1 := len(s1)
	l2 := len(s2)

	if l1 != l2 {
		return false
	}

	return isSubString(s1+s1, s2)
}
