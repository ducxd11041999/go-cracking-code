package urlify

func URLifyFixedOutput(str string, length int) string {
	ans := make([]byte, 0)

	lenStr := len(str)
	for i := 0; i < lenStr; i++ {
		if str[i] == ' ' {
			ans = append(ans, []byte("%20")...)
			continue
		}

		ans = append(ans, str[i])
	}

	return string(ans[:length])
}

func URLifyFixedInput(str string, length int) string {
	ans := make([]byte, 0)

	for i := 0; i < length; i++ {
		if str[i] == ' ' {
			ans = append(ans, []byte("%20")...)
			continue
		}

		ans = append(ans, str[i])
	}

	return string(ans)
}
